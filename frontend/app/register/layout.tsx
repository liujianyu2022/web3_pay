import { LayoutProps } from "../../.next/types/app/layout"

import styles from "./register.module.css"

export default function RegisterLayout(props: LayoutProps) {
    const { children } = props

    return (
        <div className={`${styles.register_container} w-full h-full`}>
            {children}
        </div>
    )
}