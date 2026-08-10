'use client';

import { useState } from 'react';
import Link from 'next/link';
import { usePathname } from 'next/navigation';
import { cn } from '@/lib/utils';

interface NavItem {
  label: string;
  href: string;
  icon: string;
  children?: { label: string; href: string }[];
}

const navItems: NavItem[] = [
  {
    label: '仪表盘',
    href: '/',
    icon: 'dashboard',
  },
  {
    label: '制造执行 MES',
    href: '/mes',
    icon: 'mes',
  },
  {
    label: '质量管理 QMS',
    href: '/qms',
    icon: 'qms',
    children: [
      { label: '检验单管理', href: '/qms' },
      { label: '不合格品 NCR', href: '/qms/ncr' },
    ],
  },
  {
    label: '设备管理 EAM',
    href: '/eam',
    icon: 'eam',
    children: [
      { label: '设备台账', href: '/eam' },
      { label: '维修工单', href: '/eam/repairs' },
    ],
  },
  {
    label: '仓储管理 WMS',
    href: '/wms',
    icon: 'wms',
    children: [
      { label: '仓管总览', href: '/wms' },
      { label: '库存余额', href: '/wms/inventory' },
      { label: '收货单', href: '/wms/receive' },
      { label: '发货单', href: '/wms/delivery' },
    ],
  },
  {
    label: '高级排程 APS',
    href: '/aps',
    icon: 'aps',
    children: [
      { label: 'MPS 计划', href: '/aps' },
      { label: '排程甘特图', href: '/aps/schedule' },
    ],
  },
  {
    label: '质量追溯',
    href: '/trace',
    icon: 'trace',
  },
  {
    label: '安灯系统',
    href: '/andon',
    icon: 'andon',
    children: [
      { label: '活跃呼叫', href: '/andon' },
      { label: '告警记录', href: '/andon/alerts' },
    ],
  },
];

const icons: Record<string, React.ReactNode> = {
  dashboard: (
    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <rect x="3" y="3" width="7" height="7" rx="1" />
      <rect x="14" y="3" width="7" height="7" rx="1" />
      <rect x="3" y="14" width="7" height="7" rx="1" />
      <rect x="14" y="14" width="7" height="7" rx="1" />
    </svg>
  ),
  mes: (
    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <path d="M14.7 6.3a1 1 0 0 0 0 1.4l1.6 1.6a1 1 0 0 0 1.4 0l3.77-3.77a6 6 0 0 1-7.94 7.94l-6.91 6.91a2.12 2.12 0 0 1-3-3l6.91-6.91a6 6 0 0 1 7.94-7.94l-3.76 3.76z" />
    </svg>
  ),
  qms: (
    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <path d="M9 12l2 2 4-4" />
      <path d="M12 2a10 10 0 1 0 0 20 10 10 0 0 0 0-20z" />
    </svg>
  ),
  eam: (
    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <path d="M14 12a2 2 0 1 1-4 0 2 2 0 0 1 4 0z" />
      <path d="M6.343 17.657L2 22v-7l4.343-4.343" />
      <path d="M17.657 6.343L22 2v7l-4.343 4.343" />
      <path d="M2 2l5 5" />
    </svg>
  ),
  wms: (
    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <rect x="2" y="7" width="20" height="14" rx="2" />
      <path d="M16 7V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v2" />
      <path d="M12 12v6" />
      <path d="M9 15h6" />
    </svg>
  ),
  aps: (
    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <circle cx="12" cy="12" r="10" />
      <polyline points="12 6 12 12 16 14" />
    </svg>
  ),
  trace: (
    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <circle cx="11" cy="11" r="8" />
      <path d="m21 21-4.3-4.3" />
    </svg>
  ),
  andon: (
    <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
      <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9" />
      <path d="M13.73 21a2 2 0 0 1-3.46 0" />
    </svg>
  ),
};

export function Sidebar() {
  const pathname = usePathname();
  const [collapsed, setCollapsed] = useState(false);
  const [expandedGroups, setExpandedGroups] = useState<Set<string>>(
    new Set(['质量管理 QMS', '设备管理 EAM', '仓储管理 WMS', '高级排程 APS', '安灯系统'])
  );

  const toggleGroup = (label: string) => {
    setExpandedGroups((prev) => {
      const next = new Set(prev);
      if (next.has(label)) next.delete(label);
      else next.add(label);
      return next;
    });
  };

  const isActive = (href: string) => {
    if (href === '/') return pathname === '/';
    return pathname.startsWith(href);
  };

  return (
    <aside
      className={cn(
        'flex flex-col bg-gray-900 text-gray-300 transition-all duration-300',
        collapsed ? 'w-16' : 'w-56'
      )}
    >
      {/* Logo */}
      <div className="flex h-14 items-center justify-between px-4 border-b border-gray-700">
        {!collapsed && (
          <Link href="/" className="flex items-center gap-2">
            <span className="text-lg font-bold text-white">MOM 3.0</span>
          </Link>
        )}
        <button
          onClick={() => setCollapsed(!collapsed)}
          className="rounded p-1.5 hover:bg-gray-700 transition-colors text-gray-400 hover:text-white"
        >
          <svg
            width="18"
            height="18"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            strokeWidth="2"
            className={cn('transition-transform', collapsed && 'rotate-180')}
          >
            <polyline points="15 18 9 12 15 6" />
          </svg>
        </button>
      </div>

      {/* Nav Items */}
      <nav className="flex-1 overflow-y-auto py-2 px-2">
        {navItems.map((item) => (
          <div key={item.label}>
            {item.children ? (
              <div>
                <button
                  onClick={() => toggleGroup(item.label)}
                  className={cn(
                    'flex w-full items-center gap-3 rounded-lg px-3 py-2 text-sm transition-colors hover:bg-gray-800',
                    isActive(item.href) && !collapsed && 'bg-gray-800 text-white'
                  )}
                >
                  <span className="flex-shrink-0">{icons[item.icon]}</span>
                  {!collapsed && (
                    <>
                      <span className="flex-1 text-left truncate">{item.label}</span>
                      <svg
                        width="14"
                        height="14"
                        viewBox="0 0 24 24"
                        fill="none"
                        stroke="currentColor"
                        strokeWidth="2"
                        className={cn(
                          'transition-transform',
                          expandedGroups.has(item.label) && 'rotate-90'
                        )}
                      >
                        <polyline points="9 18 15 12 9 6" />
                      </svg>
                    </>
                  )}
                </button>
                {!collapsed && expandedGroups.has(item.label) && (
                  <div className="ml-9 mt-1 space-y-1">
                    {item.children.map((child) => (
                      <Link
                        key={child.href}
                        href={child.href}
                        className={cn(
                          'block rounded-lg px-3 py-1.5 text-xs transition-colors hover:bg-gray-800 hover:text-white',
                          isActive(child.href)
                            ? 'bg-gray-800 text-white'
                            : 'text-gray-400'
                        )}
                      >
                        {child.label}
                      </Link>
                    ))}
                  </div>
                )}
              </div>
            ) : (
              <Link
                href={item.href}
                className={cn(
                  'flex items-center gap-3 rounded-lg px-3 py-2 text-sm transition-colors hover:bg-gray-800 hover:text-white',
                  isActive(item.href) && 'bg-gray-800 text-white',
                  collapsed && 'justify-center px-2'
                )}
              >
                <span className="flex-shrink-0">{icons[item.icon]}</span>
                {!collapsed && <span>{item.label}</span>}
              </Link>
            )}
          </div>
        ))}
      </nav>

      {/* Footer */}
      {!collapsed && (
        <div className="border-t border-gray-700 p-3">
          <div className="text-xs text-gray-500">MOM Platform v3.0</div>
          <div className="text-xs text-gray-500">离散制造运营管理</div>
        </div>
      )}
    </aside>
  );
}
