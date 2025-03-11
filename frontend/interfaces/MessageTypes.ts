export enum MessageTypeEnum {
    success = "success",
    info = "info" ,
    warning = "warning",
    error = "error"
}

export interface MessageType {
    content: string
    type:  MessageTypeEnum 
    duration?: number
    onClose?: () => void
}