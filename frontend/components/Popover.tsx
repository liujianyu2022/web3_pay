import React, { useState, useRef, useEffect } from 'react';
import { PopoverProps } from '../interfaces/componentTypes';


const Popover: React.FC<PopoverProps> = ({ trigger, content, placement = 'bottom' }) => {
  const [isVisible, setIsVisible] = useState(false); // 控制 Popover 显示状态
  const triggerRef = useRef<HTMLDivElement>(null); // 触发器的引用
  const popoverRef = useRef<HTMLDivElement>(null); // Popover 的引用

  // 处理触发器点击事件
  const handleTriggerClick = () => {
    setIsVisible(!isVisible);
  };

  // 处理点击外部关闭 Popover
  const handleClickOutside = (event: MouseEvent) => {
    if (
      popoverRef.current &&
      !popoverRef.current.contains(event.target as Node) &&
      triggerRef.current &&
      !triggerRef.current.contains(event.target as Node)
    ) {
      setIsVisible(false);
    }
  };

  // 监听点击外部事件
  useEffect(() => {
    document.addEventListener('mousedown', handleClickOutside);
    return () => {
      document.removeEventListener('mousedown', handleClickOutside);
    };
  }, []);

  // 计算 Popover 的位置
  const getPosition = () => {
    if (triggerRef.current) {
      const triggerRect = triggerRef.current.getBoundingClientRect();
      switch (placement) {
        case 'top':
          return { top: triggerRect.top - 10, left: triggerRect.left + triggerRect.width / 2 };
        case 'bottom':
          return { top: triggerRect.bottom + 10, left: triggerRect.left + triggerRect.width / 2 };
        case 'left':
          return { top: triggerRect.top + triggerRect.height / 2, left: triggerRect.left - 10 };
        case 'right':
          return { top: triggerRect.top + triggerRect.height / 2, left: triggerRect.right + 10 };
        default:
          return { top: triggerRect.bottom + 10, left: triggerRect.left + triggerRect.width / 2 };
      }
    }
    return { top: 0, left: 0 };
  };

  return (
    <div className="relative">
      {/* 触发器 */}
      <div ref={triggerRef} onClick={handleTriggerClick}>{trigger}</div>

      {/* Popover 内容 */}
      {isVisible && (
        <div
          ref={popoverRef}
          className="absolute z-50 p-2 bg-white border border-gray-200 rounded shadow-lg"
          style={{
            top: `${getPosition().top}px`,
            left: `${getPosition().left}px`,
            transform: 'translateX(-50%)',
          }}
        >
          {content}
        </div>
      )}
    </div>
  );
};

export default Popover;