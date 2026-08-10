'use client';

import { cn } from '@/lib/utils';
import { ButtonHTMLAttributes, forwardRef } from 'react';

export interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'default' | 'destructive' | 'outline' | 'ghost' | 'link';
  size?: 'sm' | 'md' | 'lg';
}

const variantClasses: Record<string, string> = {
  default:
    'bg-blue-600 text-white hover:bg-blue-700 shadow-sm',
  destructive:
    'bg-red-600 text-white hover:bg-red-700 shadow-sm',
  outline:
    'border border-gray-300 bg-white text-gray-700 hover:bg-gray-50',
  ghost:
    'text-gray-700 hover:bg-gray-100',
  link:
    'text-blue-600 underline-offset-4 hover:underline',
};

const sizeClasses: Record<string, string> = {
  sm: 'h-8 px-3 text-xs rounded',
  md: 'h-9 px-4 text-sm rounded-md',
  lg: 'h-10 px-6 text-sm rounded-md',
};

const Button = forwardRef<HTMLButtonElement, ButtonProps>(
  ({ className, variant = 'default', size = 'md', children, disabled, ...props }, ref) => {
    return (
      <button
        ref={ref}
        className={cn(
          'inline-flex items-center justify-center gap-2 font-medium transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-blue-500 disabled:pointer-events-none disabled:opacity-50',
          variantClasses[variant],
          sizeClasses[size],
          className
        )}
        disabled={disabled}
        {...props}
      >
        {children}
      </button>
    );
  }
);

Button.displayName = 'Button';
export { Button };
