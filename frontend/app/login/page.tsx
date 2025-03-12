"use client"

import { useRouter } from "next/navigation";
import React, { useState } from "react"
import { FaUser, FaLock } from 'react-icons/fa';
import { login } from "../../api/user";
import message from "../../components/Message";
import { MessageTypeEnum } from "../../interfaces/messageTypes";

export default function LoginPage() {

    const router = useRouter()

    const [email, setEmail] = useState<string>('');
    const [password, setPassword] = useState<string>('');

    const logins = async (event: React.FormEvent) => {
        event.preventDefault()
        try {
            const { data: { code, message: msg, data }, status } = await login({ email, password })
            if (status === 200 && code === 0) {
                message(msg, MessageTypeEnum.success)
            }else{
                throw Error(msg)
            }
        } catch (error: any) {
            message(error.message, MessageTypeEnum.error)
        }
    }

    const register = () => {
        router.push("/register", {})
    }

    const iconStyle = "absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-500"
    // p-3: 对于输入框非常有用，可以确保文本不会紧贴边框
    // focus:outline-none： 移除输入框在获取焦点时默认显示的外部轮廓。这通常用于自定义焦点样式，以避免浏览器默认的样式影响页面设计
    // focus:ring-2: 为输入框的聚焦状态添加一个 2px 宽的环形边框。ring 在聚焦时环绕在元素的外部，类似于 box-shadow
    const inputStyle = "w-full pl-10 p-3  border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"

    return (
        <div className="w-full h-full flex items-center justify-center">
            <div className="flex flex-col justify-center items-center pt-6 pr-3 pb-6 pl-3 w-1/3 h-3/5 bg-white rounded-lg">
                <div className="text-lg font-bold mb-20">Welcome Back</div>
                <form onSubmit={logins} className="w-full flex flex-col items-center mb-4">
                    <div className="mb-4 relative w-4/5">
                        <FaUser className={iconStyle} />
                        <input
                            className={inputStyle}
                            type="text"
                            id="email"
                            name="email"
                            value={email}
                            onChange={event => setEmail(event.target.value)}
                            placeholder="Enter your email"
                            required
                        />
                    </div>
                    <div className="mb-4 relative w-4/5">
                        <FaLock className={iconStyle} />
                        <input
                            className={inputStyle}
                            type="password"
                            id="password"
                            name="password"
                            value={password}
                            onChange={event => setPassword(event.target.value)}
                            placeholder="Enter your email"
                            required
                        />
                    </div>
                    <div className="w-4/5 flex justify-between">
                        <button
                            type="submit"
                            className="pt-2 pb-2 w-2/5 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus: ring-2 focus:ring-opacity-50"
                        >
                            Login
                        </button>
                        <button
                            type="button"
                            onClick={() => { setEmail(""); setPassword("") }}
                            className="w-2/5 pt-2 pb-2 border border-gray-300 rounded-md hover:border-gray-400"
                        >
                            Reset
                        </button>
                    </div>
                </form>

                <div>
                    <span>Don't have a account yet? </span>
                    <span className="cursor-pointer text-blue-500" onClick={register}>Register Now!</span>
                </div>
            </div>
        </div>

    )
}