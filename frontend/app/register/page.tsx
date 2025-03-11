"use client"

import { useState } from "react";
import { AiFillMail } from "react-icons/ai";
import { FaLock, FaUser } from "react-icons/fa"
import message from "../../componets/Message";
import { MessageTypeEnum } from "../../interfaces/MessageTypes";
import { register } from "../../api/user";

export default function RegisterPage() {
    const [email, setEmail] = useState<string>('')
    const [nickname, setNickname] = useState<string>('')
    const [password, setPassword] = useState<string>('')
    const [rePassword, setRePassword] = useState<string>('')

    const registers = async (event: React.FormEvent) => {
        event.preventDefault()

        try {
            const result = await register({email, nickname, password, re_password: rePassword})
            console.log("result = ", result)
        } catch (error: any) {
            message(error.message, MessageTypeEnum.error)
        }
    }

    return (
        <div className="w-full h-full flex items-center justify-center">
            <div className="w-1/3 h-3/5 p-5 flex items-center justify-center bg-white rounded-md">
                <form onSubmit={registers} className="w-full flex flex-col justify-center items-center">
                    <div className="mb-8 relative w-4/5">
                        < AiFillMail className="absolute left-3 top-1/2 transform -translate-y-1/2" />
                        <input
                            type="text"
                            className="w-full pl-10 p-3 rounded-md border border-gray-300 focus:outline-none focus:ring-2"
                            id="email"
                            name="email"
                            value={email}
                            onChange={event => setEmail(event.target.value)}
                            placeholder="Enter your email"
                            required
                        />
                    </div>

                    <div className="mb-8 relative w-4/5">
                        <FaUser className="absolute left-3 top-1/2 transform -translate-y-1/2" />
                        <input
                            type="text"
                            className="w-full pl-10 p-3 rounded-md border border-gray-300 focus:outline-none focus:ring-2"
                            id="nickname"
                            name="nickname"
                            value={nickname}
                            onChange={event => setNickname(event.target.value)}
                            placeholder="Enter your nickname"
                            required
                        />
                    </div>

                    <div className="mb-8 relative w-4/5">
                        <FaLock className="absolute left-3 top-1/2 transform -translate-y-1/2" />
                        <input
                            type="password"
                            className="w-full pl-10 p-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2"
                            id="password"
                            name="password"
                            value={password}
                            onChange={event => setPassword(event.target.value)}
                            placeholder="Enter your password"
                            required
                        />
                    </div>

                    <div className="mb-8 relative w-4/5">
                        <FaLock className="absolute left-3 top-1/2 transform -translate-y-1/2" />
                        <input
                            type="password"
                            className="w-full pl-10 p-3 border border-gray-300 rounded-md focus:outline-none focus:ring-2"
                            id="rePassword"
                            name="rePassword"
                            value={rePassword}
                            onChange={event => setRePassword(event.target.value)}
                            placeholder="Reenter your password"
                            required
                        />
                    </div>

                    <div className="w-4/5 flex justify-between">
                        <button
                            type="submit"
                            className="w-2/5 pt-2 pb-2 border rounded-md bg-blue-500 text-white hover:bg-blue-600"
                        >
                            Register
                        </button>

                        <button
                            type="button"
                            className="w-2/5 pt-2 pb-2 border rounded-md border-gray-300 hover:border-gray-400"
                            onClick={() => {setEmail(""); setNickname(""); setPassword(""); setRePassword("") }}
                        >
                            Reset
                        </button>
                    </div>
                </form>
            </div>
        </div>
    )
}