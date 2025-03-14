import React from "react";

export interface LayoutType {
    children: React.ReactNode
}

export interface ResponseType {
    code: number
    message: string
    data: any
}



export interface Column<T> {
    key: keyof T
    title: string
    sortable?: boolean
    align?: "center" | "left" | "right"
}
export interface TableSortConfig<T> {
    key: keyof T | null
    direction: "descend" | "ascend"
}

export interface TableProps<T> {
    columns: Array<Column<T>>
    data: Array<T>
    onSort?: (parameters: TableSortConfig<T>) => void
}

