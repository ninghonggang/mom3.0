import { BadRequestException, Injectable, Logger } from '@nestjs/common';
import { GrpcService } from './grpc.service';

const MDM_PKG = 'mom.mdm';
const MDM_URL = process.env.MDM_SERVICE_URL || 'localhost:50051';
const WMS_PKG = 'mom.wms';
const WMS_URL = process.env.WMS_SERVICE_URL || 'localhost:50055';

/** A material as resolved from the MDM service. */
export interface ResolvedMaterial {
  id: number;
  materialCode: string;
  materialName: string;
  spec: string;
}

/** A storage location as resolved from the WMS service. */
export interface ResolvedLocation {
  id: number;
  locationCode: string;
  warehouseId: number;
}

/**
 * Cross-domain code → id resolver.
 *
 * REST clients speak in human-readable business codes ("RAW-002", "A-01-01"),
 * while the gRPC contracts are strictly id-based. This service performs the
 * lookup once per request path and caches the result for the process lifetime,
 * so hot write paths do not pay a round-trip per call.
 */
@Injectable()
export class ResolverService {
  private readonly logger = new Logger(ResolverService.name);
  private readonly materialCache = new Map<string, ResolvedMaterial>();
  private readonly locationCache = new Map<string, ResolvedLocation>();
  private readonly warehouseCache = new Map<string, number>();

  constructor(private readonly grpc: GrpcService) {}

  /** Resolve a material code to its MDM id. Throws 400 when unknown. */
  async material(code: string): Promise<ResolvedMaterial> {
    if (!code) throw new BadRequestException('materialCode is required');
    const cached = this.materialCache.get(code);
    if (cached) return cached;

    const res: any = await this.grpc.call(
      MDM_PKG,
      'MaterialService',
      'ListMaterials',
      { keyword: code, pagination: { page: 1, page_size: 200 } },
      MDM_URL,
    );
    const items: any[] = res?.items ?? [];
    const hit = items.find((m) => m.materialCode === code);
    if (!hit) {
      throw new BadRequestException(`material not found: ${code}`);
    }
    const resolved: ResolvedMaterial = {
      id: Number(hit.base?.id ?? hit.id ?? 0),
      materialCode: hit.materialCode ?? code,
      materialName: hit.materialName ?? '',
      spec: hit.spec ?? '',
    };
    this.materialCache.set(code, resolved);
    return resolved;
  }

  /** Resolve a material code to just its numeric id. */
  async materialId(code: string): Promise<number> {
    return (await this.material(code)).id;
  }

  /**
   * Resolve a warehouse code to its WMS id.
   * Accepts a numeric string as a pass-through so callers may send either.
   */
  async warehouseId(codeOrId: string): Promise<number> {
    if (!codeOrId) return 0;
    if (/^\d+$/.test(codeOrId)) return Number(codeOrId);
    const cached = this.warehouseCache.get(codeOrId);
    if (cached) return cached;

    const res: any = await this.grpc.call(
      WMS_PKG,
      'WmsService',
      'ListWarehouses',
      { page: { page: 1, page_size: 200 } },
      WMS_URL,
    );
    const hit = (res?.items ?? []).find((w: any) => w.warehouseCode === codeOrId);
    if (!hit) throw new BadRequestException(`warehouse not found: ${codeOrId}`);
    const id = Number(hit.id ?? 0);
    this.warehouseCache.set(codeOrId, id);
    return id;
  }

  /**
   * Resolve a location code to its WMS id.
   * Accepts a numeric string as a pass-through. Returns 0 when no code given,
   * letting callers treat the location as optional.
   *
   * Pass `required: true` for movements that must land in a real bin
   * (putaway / picking) — silently writing `location_id = 0` there would
   * corrupt the inventory ledger, so an unknown code raises 400 instead.
   */
  async locationId(
    codeOrId?: string,
    warehouseId?: number,
    opts: { required?: boolean } = {},
  ): Promise<number> {
    if (!codeOrId) {
      if (opts.required) throw new BadRequestException('locationCode is required');
      return 0;
    }
    if (/^\d+$/.test(codeOrId)) return Number(codeOrId);
    const cacheKey = `${warehouseId ?? 0}:${codeOrId}`;
    const cached = this.locationCache.get(cacheKey);
    if (cached) return cached.id;

    const req: any = { page: { page: 1, page_size: 500 } };
    if (warehouseId) req.warehouse_id = warehouseId;
    const res: any = await this.grpc.call(
      WMS_PKG,
      'WmsService',
      'ListLocations',
      req,
      WMS_URL,
    );
    const hit = (res?.items ?? []).find((l: any) => l.locationCode === codeOrId);
    if (!hit) {
      if (opts.required) {
        throw new BadRequestException(
          `location not found: ${codeOrId} — create it first via POST /api/wms/locations`,
        );
      }
      this.logger.warn(`location not found: ${codeOrId}, falling back to 0`);
      return 0;
    }
    const resolved: ResolvedLocation = {
      id: Number(hit.id ?? 0),
      locationCode: hit.locationCode,
      warehouseId: Number(hit.warehouseId ?? 0),
    };
    this.locationCache.set(cacheKey, resolved);
    return resolved.id;
  }

  /** Drop cached locations — call after creating one so lookups see it. */
  invalidateLocations(): void {
    this.locationCache.clear();
  }

  /** Drop cached warehouses — call after creating one so lookups see it. */
  invalidateWarehouses(): void {
    this.warehouseCache.clear();
  }
}
