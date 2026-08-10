export function cn(...classes: (string | boolean | undefined | null)[]): string {
  return classes.filter(Boolean).join(' ');
}

export function formatDate(date: string | undefined | null): string {
  if (!date) return '-';
  return new Date(date).toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
  });
}

export function formatDateTime(date: string | undefined | null): string {
  if (!date) return '-';
  return new Date(date).toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
  });
}

export function formatNumber(n: number, decimals = 0): string {
  return n.toLocaleString('zh-CN', {
    minimumFractionDigits: decimals,
    maximumFractionDigits: decimals,
  });
}

export function formatPercent(n: number, decimals = 1): string {
  return `${(n * 100).toFixed(decimals)}%`;
}

export const ORDER_STATUS_MAP: Record<string, { label: string; color: string }> = {
  DRAFT: { label: '草稿', color: 'bg-gray-100 text-gray-600' },
  RELEASED: { label: '已下达', color: 'bg-blue-100 text-blue-700' },
  IN_PROGRESS: { label: '生产中', color: 'bg-yellow-100 text-yellow-700' },
  COMPLETED: { label: '已完成', color: 'bg-green-100 text-green-700' },
  CLOSED: { label: '已关闭', color: 'bg-gray-200 text-gray-500' },
  ON_HOLD: { label: '暂停', color: 'bg-orange-100 text-orange-700' },
};

export const EQUIPMENT_STATUS_MAP: Record<string, { label: string; color: string }> = {
  RUNNING: { label: '运行中', color: 'bg-green-100 text-green-700' },
  IDLE: { label: '待机', color: 'bg-gray-100 text-gray-600' },
  MAINTENANCE: { label: '保养中', color: 'bg-blue-100 text-blue-700' },
  REPAIR: { label: '维修中', color: 'bg-red-100 text-red-700' },
  STOPPED: { label: '停机', color: 'bg-red-200 text-red-800' },
};

export const INSPECTION_STATUS_MAP: Record<string, { label: string; color: string }> = {
  PENDING: { label: '待检验', color: 'bg-gray-100 text-gray-600' },
  IN_PROGRESS: { label: '检验中', color: 'bg-blue-100 text-blue-700' },
  COMPLETED: { label: '已完成', color: 'bg-green-100 text-green-700' },
  CLOSED: { label: '已关闭', color: 'bg-gray-200 text-gray-500' },
};

export const NCR_SEVERITY_MAP: Record<string, { label: string; color: string }> = {
  CRITICAL: { label: '严重', color: 'bg-red-100 text-red-700' },
  MAJOR: { label: '主要', color: 'bg-orange-100 text-orange-700' },
  MINOR: { label: '次要', color: 'bg-yellow-100 text-yellow-700' },
};

export const NCR_STATUS_MAP: Record<string, { label: string; color: string }> = {
  OPEN: { label: '开启', color: 'bg-red-100 text-red-700' },
  DISPOSITION: { label: '处置中', color: 'bg-yellow-100 text-yellow-700' },
  IN_REWORK: { label: '返工中', color: 'bg-blue-100 text-blue-700' },
  VERIFIED: { label: '已验证', color: 'bg-green-100 text-green-700' },
  CLOSED: { label: '已关闭', color: 'bg-gray-200 text-gray-500' },
};

export const REPAIR_STATUS_MAP: Record<string, { label: string; color: string }> = {
  REPORTED: { label: '已报修', color: 'bg-red-100 text-red-700' },
  ASSIGNED: { label: '已分配', color: 'bg-blue-100 text-blue-700' },
  IN_PROGRESS: { label: '维修中', color: 'bg-yellow-100 text-yellow-700' },
  COMPLETED: { label: '已修复', color: 'bg-green-100 text-green-700' },
  CONFIRMED: { label: '已确认', color: 'bg-gray-200 text-gray-500' },
};

export const ANDON_STATUS_MAP: Record<string, { label: string; color: string }> = {
  ACTIVE: { label: '活跃', color: 'bg-red-500 text-white' },
  RESPONDED: { label: '已响应', color: 'bg-yellow-100 text-yellow-700' },
  RESOLVED: { label: '已解决', color: 'bg-green-100 text-green-700' },
  TIMEOUT: { label: '超时', color: 'bg-orange-100 text-orange-700' },
};

export const INVENTORY_STATUS_MAP: Record<string, { label: string; color: string }> = {
  NORMAL: { label: '正常', color: 'bg-green-100 text-green-700' },
  LOCKED: { label: '锁定', color: 'bg-yellow-100 text-yellow-700' },
  ON_HOLD: { label: '冻结', color: 'bg-orange-100 text-orange-700' },
  SCRAP: { label: '报废', color: 'bg-red-100 text-red-700' },
};

export function getStatusLabel(map: Record<string, { label: string; color: string }>, status: string): string {
  return map[status]?.label || status;
}

export function getStatusColor(map: Record<string, { label: string; color: string }>, status: string): string {
  return map[status]?.color || 'bg-gray-100 text-gray-600';
}
