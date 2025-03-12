import { redirect } from "next/navigation"

// 首页（对应 `/` 路由）
export default function Home() {
  redirect("/login")
}
