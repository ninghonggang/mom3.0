/**
 * 深浅蛇形(snake_case) <-> 驼峰(camelCase) 互转工具。
 * gRPC 服务端(protoc 生成)使用 snake_case 字段名，REST 层使用 camelCase。
 */

function toCamel(s: string): string {
  return s.replace(/_([a-z0-9])/g, (_, c: string) => c.toUpperCase());
}

function toSnake(s: string): string {
  return s.replace(/[A-Z]/g, (c) => '_' + c.toLowerCase());
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
function isPlainObject(v: any): boolean {
  // Accept both plain objects and validated DTO class instances (NestJS @Body).
  // Exclude arrays, dates, and other special objects.
  return (
    v !== null &&
    typeof v === 'object' &&
    !Array.isArray(v) &&
    !(v instanceof Date)
  );
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function deepSnakeToCamel<T = any>(obj: any): T {
  if (Array.isArray(obj)) return obj.map((v) => deepSnakeToCamel(v)) as unknown as T;
  if (isPlainObject(obj)) {
    const out: Record<string, unknown> = {};
    for (const k of Object.keys(obj)) {
      out[toCamel(k)] = deepSnakeToCamel(obj[k]);
    }
    return out as T;
  }
  return obj as T;
}

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function deepCamelToSnake<T = any>(obj: any): T {
  if (Array.isArray(obj)) return obj.map((v) => deepCamelToSnake(v)) as unknown as T;
  if (isPlainObject(obj)) {
    const out: Record<string, unknown> = {};
    for (const k of Object.keys(obj)) {
      out[toSnake(k)] = deepCamelToSnake(obj[k]);
    }
    return out as T;
  }
  return obj as T;
}

/**
 * 将 ISO-8601 时间字符串转换为 protobuf Timestamp 结构（{seconds, nanos}）。
 * @grpc/proto-loader 期望 google.protobuf.Timestamp 为 {seconds, nanos} 而非字符串。
 * 传入空值时返回 undefined（proto 字段留空）。
 */
export function toProtoTimestamp(
  iso: string | undefined | null,
): { seconds: number; nanos: number } | undefined {
  if (!iso) return undefined;
  const ms = Date.parse(iso);
  if (Number.isNaN(ms)) return undefined;
  return {
    seconds: Math.floor(ms / 1000),
    nanos: (ms % 1000) * 1e6,
  };
}

/** 列表分页风格：MDM/MES/APS 嵌套 pagination，EAM/WMS/QMS/Trace 用 page */
export type PaginationStyle = 'pagination' | 'page';

/**
 * 由 controller 的 query（page/pageSize + 过滤条件，camelCase）构造 gRPC 列表请求。
 */
// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function buildListRequest(
  query: Record<string, any> | undefined,
  style: PaginationStyle,
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
): Record<string, any> {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const q: Record<string, any> = { ...(query ?? {}) };
  const page = Number(q.page) || 1;
  const pageSize = Number(q.pageSize) || 20;
  delete q.page;
  delete q.pageSize;
  const pg = { page, page_size: pageSize };
  return style === 'pagination' ? { pagination: pg, ...q } : { page: pg, ...q };
}

/**
 * 把外部传入的枚举值规范成 proto 枚举名。
 *
 * proto 枚举一律带域前缀（如 `INSPECTION_TYPE_IQC`），而 REST 客户端习惯
 * 传裸值（`IQC` / `iqc`）。此函数补齐前缀并大写，已带前缀的原样返回。
 * 传空时返回 `${prefix}_UNSPECIFIED`，保证 proto-loader 不会因未知枚举报错。
 */
export function toProtoEnum(prefix: string, value?: string | null): string {
  const p = prefix.toUpperCase();
  if (!value) return `${p}_UNSPECIFIED`;
  const v = String(value).trim().toUpperCase().replace(/[-\s]+/g, '_');
  return v.startsWith(`${p}_`) ? v : `${p}_${v}`;
}

/** 把任意值安全转成 proto int64（number），非法值归零。 */
export function toInt(value: unknown): number {
  const n = Number(value);
  return Number.isFinite(n) ? Math.trunc(n) : 0;
}

/** 把任意值转成 proto 的 decimal-as-string 表示，非法值归 "0"。 */
export function toDecimalString(value: unknown): string {
  const n = Number(value);
  return Number.isFinite(n) ? String(n) : '0';
}
