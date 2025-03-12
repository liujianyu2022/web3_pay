import React from "react";

export interface LayoutType {
    children: React.ReactNode 
}

export interface ResponseType {
    code: number
    message: string
    data: any
}