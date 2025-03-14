export enum UserAction {
    ACCESS_TOKEN_INITIAL = "user-access-token-initial",
    ACCESS_TOKEN_CHANGE = "user-access-token-change"
}
export interface UserStateType {
    email?: string
    username?: string
    nickname?: string
    accessToken?: string
}
export interface UserActionType {
    type: UserAction,
    payload: {
        accessToken?: string
    }
}