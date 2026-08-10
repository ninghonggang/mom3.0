import Link from 'next/link';

export function Header() {
  return (
    <header className="flex h-14 items-center justify-between border-b bg-white px-6">
      {/* Breadcrumb / Title Area */}
      <div className="flex items-center gap-3">
        <div className="flex items-center gap-2">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="#6b7280" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
            <path d="m3 9 9-7 9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z" />
            <polyline points="9 22 9 12 15 12 15 22" />
          </svg>
          <span className="text-sm text-gray-400">/</span>
          <span className="text-sm font-medium text-gray-700">制造运营管理系统</span>
        </div>
      </div>

      {/* Right side - User Info */}
      <div className="flex items-center gap-4">
        {/* Quick search */}
        <div className="relative hidden md:block">
          <input
            type="text"
            placeholder="快速搜索..."
            className="w-64 rounded-lg border border-gray-200 bg-gray-50 py-1.5 pl-8 pr-3 text-sm text-gray-600 placeholder-gray-400 focus:border-blue-400 focus:outline-none focus:ring-1 focus:ring-blue-400"
          />
          <svg
            width="14"
            height="14"
            viewBox="0 0 24 24"
            fill="none"
            stroke="#9ca3af"
            strokeWidth="2"
            className="absolute left-2.5 top-1/2 -translate-y-1/2"
          >
            <circle cx="11" cy="11" r="8" />
            <path d="m21 21-4.3-4.3" />
          </svg>
        </div>

        {/* Notifications */}
        <button className="relative rounded-lg p-1.5 text-gray-400 hover:bg-gray-100 hover:text-gray-600">
          <svg width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
            <path d="M18 8A6 6 0 0 0 6 8c0 7-3 9-3 9h18s-3-2-3-9" />
            <path d="M13.73 21a2 2 0 0 1-3.46 0" />
          </svg>
          <span className="absolute -right-0.5 -top-0.5 flex h-4 w-4 items-center justify-center rounded-full bg-red-500 text-[10px] font-medium text-white">
            3
          </span>
        </button>

        {/* User */}
        <div className="flex items-center gap-2 pl-3 border-l border-gray-200">
          <div className="flex h-8 w-8 items-center justify-center rounded-full bg-blue-600 text-sm font-medium text-white">
            张
          </div>
          <div className="hidden sm:block">
            <div className="text-sm font-medium text-gray-700">张操作员</div>
            <div className="text-xs text-gray-400">生产主管</div>
          </div>
        </div>
      </div>
    </header>
  );
}
