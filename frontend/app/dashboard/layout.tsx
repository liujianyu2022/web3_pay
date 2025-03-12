"use client"

import React from "react";
import { LayoutType } from "../../interfaces/commonTypes";

import styles from "./dashboard.module.css"
import { useRouter } from "next/navigation";

const menus = [
    { name: "dashboard", path: "/" },
    { name: "orders", path: "/dashboard/order" },
    { name: "merchants", path: "/dashboard/merchant" },
    { name: "merchantAPIs", path: "/dashboard/merchantAPI" },
    { name: "wallet", path: "/dashboard/wallet" },
    { name: "system", path: "/dashboard/system" },
]

export default function Dashboard(props: LayoutType) {
    const { children } = props
    const router = useRouter()

    const [active, setActive] = React.useState<string>("/")

    const click = (path: string) => {
        setActive(path)
        router.push(path)
    }

    return (
        <div className={`${styles.dashborad_layout} flex w-full h-full p-2 text-base`} >
            <div className="w-1/12 h-full">
                {
                    menus.map(menu => (
                        <div
                            className={`h-12 text-center cursor-pointer hover:bg-blue-300 ${menu.path === active ? 'bg-blue-300' : ''}`}
                            key={menu.name}
                            onClick={() => click(menu.path)}
                        >
                            {menu.name}
                        </div>
                    ))
                }
            </div>
            <div>

            </div>
            {children}
        </div>
    )
}