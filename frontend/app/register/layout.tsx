import { LayoutType } from "../../interfaces/commonTypes";

import styles from "./register.module.css"

export default function RegisterLayout(props:LayoutType) {
    const { children } = props

    return (
        <div className={`${styles.register_container} w-full h-full`}>
            {children}
        </div>
    )
}