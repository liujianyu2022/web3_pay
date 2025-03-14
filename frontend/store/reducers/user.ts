import { UserAction, UserActionType, UserStateType } from "../../interfaces/reducerTypes";

const initialState: UserStateType = {
    accessToken: ""
}

const userReducer = (state: UserStateType = initialState, action: UserActionType) => {
    const { type, payload } = action
    switch (type) {
        case UserAction.ACCESS_TOKEN_INITIAL:
            return state
        case UserAction.ACCESS_TOKEN_CHANGE:
            return { ...state, accessToken: payload.accessToken }
        default:
            return state
    }
}

export default userReducer