"use client"

import Table from "../../../components/Table"
import { Column } from "../../../interfaces/commonTypes"
import { MerchantTableType } from "../../../interfaces/merchantTypes"

const columns: Column<MerchantTableType>[] = [
    { key: "id", title: "ID", align: "center" },
    { key: "name", title: "Name", align: "center" },
    { key: "email", title: "Email", align: "center" },
    { key: "phone", title: "Phone", align: "center" },
    { key: "action", title: "Action", align: "center" },
]

const data: Array<MerchantTableType> = [
    { id: "1", name: "1", phone: "1", email: "1", action: "1" },
    { id: "2", name: "2", phone: "2", email: "2", action: "2" },
]

export default function MerchantPage() {



    return (
        <div>
            <Table columns={columns} data={data} />
        </div>
    )
}