"use client"

import React from "react"
import { MessageType, MessageTypeEnum } from "../interfaces/MessageTypes"
import { createRoot } from "react-dom/client"

const colorMap = {
    [MessageTypeEnum.success]: "bg-green-100",
    [MessageTypeEnum.info]: "bg-gray-100",
    [MessageTypeEnum.warning]: "bg-yellow-100",
    [MessageTypeEnum.error]: "bg-red-600"
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
                    <div className={`${colorMap[type]} p-2 border min-w-[120px] min-h-[40px] text-black w-10 h-8 grid place-items-center`}>
                        <span>{content}</span>
                        <span>×</span>
                    </div>
            }
        </>
    )
}

const message = (content: string, type: MessageTypeEnum, duration?: number, onClose?: () => void) => {
    console.log("调用了 = ")

    const container = document.createElement('div')

    Object.assign(container.style, {
        position: 'fixed',          // 脱离文档流
        top: "20px",                // 顶部距离
        left: "50%",                // 将容器的左边对齐到视口的中间
        transform: "translateX(-50%)", // 将容器向左移动其自身宽度的一半，实现居中
        width: 'auto',              // 宽度自适应
        zIndex: '9999',
        pointerEvents: 'none'       // 不影响页面交互
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