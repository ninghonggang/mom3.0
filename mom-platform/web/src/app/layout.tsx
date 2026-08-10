import type { Metadata } from 'next';
import './globals.css';
import { Sidebar } from '@/components/layout/sidebar';
import { Header } from '@/components/layout/header';

export const metadata: Metadata = {
  title: 'MOM 3.0 - 制造运营管理系统',
  description: '离散制造多品种小批量 MOM/MES 平台',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="zh-CN">
      <body className="min-h-screen bg-gray-50">
        <div className="flex h-screen overflow-hidden">
          <Sidebar />
          <div className="flex flex-1 flex-col overflow-hidden">
            <Header />
            <main className="flex-1 overflow-auto p-6">{children}</main>
          </div>
        </div>
      </body>
    </html>
  );
}
