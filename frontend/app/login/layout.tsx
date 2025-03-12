import { LayoutType } from "../../interfaces/commonTypes";

import styles from "./login.module.css"


export default function LoginLayout(props: LayoutType) {
    const { children } = props

    return (
        // 不要在子布局 (LoginLayout) 中直接使用 <html> 和 <body>，这些标签应仅存在于根布局中
        <div className={`${styles.login_container} w-full h-full`}>
            {children}
        </div>
    )
}