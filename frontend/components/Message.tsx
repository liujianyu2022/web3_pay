"use client"

import React from "react"
import { MessageType, MessageTypeEnum } from "../interfaces/messageTypes"
import { createRoot } from "react-dom/client"

const colorMap = {
    [MessageTypeEnum.success]: "bg-green-600",
    [MessageTypeEnum.info]: "bg-gray-600",
    [MessageTypeEnum.warning]: "bg-orange-600",
    [MessageTypeEnum.error]: "bg-red-500"
}

const Message = (props: MessageType) => {
    const { content, type, duration = 50, onClose } = props

    const [visible, setVisible] = React.useState<boolean>(true)

    React.useEffect(() => {
        const timer = setTimeout(() => {
            setVisible(false);
            onClose && onClose();
        }, duration * 1000)

        return () => clearTimeout(timer)
    }, [duration, onClose])

    return (
        <>
            {
                !visible ? null :
                    <div className={
                        `${colorMap[type]} flex justify-center p-2 border rounded-md min-h-[40px] text-white 
                        transform transition-all duration-1000 ease-in-out ${visible ? "opacity-100 translate-y-0" : "opacity-0 -translate-y-5"}`
                    }>
                        <div className="w-auto">{content}</div>
                        <div className="cursor-pointer ml-3" onClick={()=>setVisible(false)}>×</div>
                    </div>
            }
        </>
    )
}

const message = (content: string, type: MessageTypeEnum, duration?: number, onClose?: () => void) => {

    const container = document.createElement('div')

    Object.assign(container.style, {
        position: 'fixed',          // 脱离文档流
        top: "20px",                // 顶部距离
        left: "50%",                // 将容器的左边对齐到视口的中间
        transform: "translateX(-50%)", // 将容器向左移动其自身宽度的一半，实现居中
        width: 'auto',              // 宽度自适应
        zIndex: '9999',
    })

    document.body.appendChild(container)

    const root = createRoot(container)

    const onCloseHandler = () => {
        root.unmount()                          
        document.body.removeChild(container)
        onClose && onClose()
    }

    root.render(
        <Message content={content} type={type} duration={duration} onClose={onCloseHandler} />
    )
}

export default message