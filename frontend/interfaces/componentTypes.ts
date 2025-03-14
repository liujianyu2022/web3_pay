// 定义 Popover 的 props 类型
export interface PopoverProps {
  trigger: React.ReactNode; // 触发器元素
  content: React.ReactNode; // Popover 内容
  placement?: 'top' | 'bottom' | 'left' | 'right'; // 弹出位置
}