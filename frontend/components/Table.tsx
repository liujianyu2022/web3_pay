import React from "react";
import { TableProps, TableSortConfig } from "../interfaces/commonTypes";


const Table = <T extends object>(props: TableProps<T>) => {
    const { columns, data, onSort } = props

    const [sortConfig, setSortConfig] = React.useState<TableSortConfig<T>>({ key: null, direction: "ascend" })

    const handleSort = (key: keyof T) => {
        if (!onSort) return;

        let direction: "descend" | "ascend" = "ascend";
        if (sortConfig.key === key && sortConfig.direction === "ascend") direction = "descend"

        setSortConfig({ key, direction })

        onSort({ key, direction })
    }

    const renderHeader = () => (
        <thead>
            <tr>
                {columns.map((column) => (
                    <th
                        key={column.key.toString()}
                        onClick={() => column.sortable && handleSort(column.key)}
                        className={`px-6 py-3 bg-gray-50 text-xs font-medium text-gray-500 tracking-wider 
                            ${column.align === "left" ? "text-left" : column.align === "center" ? "text-center" : column.align === "right" ? "text-right" : "text-left"}`
                        }
                    >
                        {column.title}
                        {sortConfig.key === column.key && (
                            <span className="ml-2">{sortConfig.direction === "ascend" ? '↑' : '↓'}</span>
                        )}
                    </th>
                ))}
            </tr>
        </thead>
    )


    const renderBody = () => (
        <tbody className="bg-white divide-y divide-gray-200">
            {data.map((row, rowIndex) => (
                <tr key={rowIndex}>
                    {columns.map((column) => (
                        <td
                            key={column.key.toString()}
                            className={`px-6 py-4 whitespace-nowrap text-sm text-gray-500
                                ${column.align === "left" ? "text-left" : column.align === "center" ? "text-center" : column.align === "right" ? "text-right" : "text-left"}`
                            }
                        >
                            {row[column.key] as React.ReactNode}
                        </td>
                    ))}
                </tr>
            ))}
        </tbody>
    )

    return (
        <div className="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
            <table className="min-w-full divide-y divide-gray-200">
                {renderHeader()}
                {renderBody()}
            </table>
        </div>
    )
}

export default Table