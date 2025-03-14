"use client"

import React from "react";
import { LayoutType } from "../../interfaces/commonTypes";
import styles from "./dashboard.module.css"
import { useRouter, usePathname } from "next/navigation";
import Popover from "../../components/Popover";

const menus = [
    { name: "dashboard", path: "/dashboard" },
    { name: "orders", path: "/dashboard/order" },
    { name: "merchants", path: "/dashboard/merchant" },
    { name: "merchantAPIs", path: "/dashboard/merchantAPI" },
    { name: "wallet", path: "/dashboard/wallet" },
    { name: "system", path: "/dashboard/system" },
]

export default function Dashboard(props: LayoutType) {
    const { children } = props
    const router = useRouter()
    const pathname = usePathname()
    
    const [active, setActive] = React.useState<string>("/dashboard")

    const click = (path: string) => {
        setActive(path)
        router.push(path) 
    }

    const trigger = (
        <div>

        </div>
    )

    const content = (
        <div>

        </div>
    )

    return (
        <div className={`${styles.dashborad_layout} flex w-full h-full text-base`} >
            <div className="w-1/6 h-full hidden md:block">
                {
                    menus.map(menu => (
                        <div
                            className={`h-12 leading-12 text-center cursor-pointer hover:bg-blue-200 ${menu.path === active ? 'bg-blue-200 border-r-2 border-r-blue-500' : ''}`}
                            key={menu.name}
                            onClick={() => click(menu.path)}
                        >
                            {menu.name}
                        </div>
                    ))
                }
            </div>

            <div className="w-full md:w-5/6 h-full ">
                <div className="w-full h-12 flex items-center p-4">
                    <div className="text-xl">Home / {pathname.split("/").pop()}</div>
                    <div></div>
                </div>
                <div className="w-full p-4 bg-gray-100 h-[calc(100%-3rem)]">{children}</div>
            </div>
        </div>
    )
}