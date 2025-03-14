"use client"

import React from "react"
import { redirect } from "next/navigation"
import { Provider, useSelector } from "react-redux";
import { persistor, store, StoreType } from "../store/store";
import { PersistGate } from "redux-persist/integration/react";
import { UserStateType } from "../interfaces/reducerTypes";

// 首页（对应 `/` 路由）
export default function Home() {
  const { accessToken } = useSelector<StoreType, UserStateType>(store => store.userReducer)

  React.useEffect(() => {
    if(!accessToken) redirect("/login")
  }, [accessToken])

}
