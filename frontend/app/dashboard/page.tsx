"use client"

import { useSelector } from "react-redux"
import { StoreType } from "../../store/store"
import { UserStateType } from "../../interfaces/reducerTypes"
import React from "react"

const displayLable = ["Total Merchants", "Total Order", "Success in 7 Days", "Total Amount"]
const initialDisplayValue = [0, 0, 0, 0]

export default function DashboradPage() {
    const [displayValue, setDisplayValue] = React.useState<Array<number>>(initialDisplayValue)

    const { accessToken } = useSelector<StoreType, UserStateType>(store => store.userReducer)

    return (
        <div className="w-full h-full">
            <div className="h-20 grid grid-cols-1 md:grid-cols-4 gap-4">
                {
                    displayLable.map((label, index) => (
                        <div className="bg-white rounded-md flex flex-col justify-center items-center">
                            <p className="h-1/2 font-bold text-xl">{label}</p>
                            <p className="text-xl">{displayValue[index]}</p>
                        </div>
                    ))
                }
            </div>
        </div>
    )
}