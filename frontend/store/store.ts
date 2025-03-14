import { configureStore } from '@reduxjs/toolkit';
import { PersistConfig, persistReducer, persistStore } from 'redux-persist';
import storage from 'redux-persist/lib/storage'; // 默认使用 localStorage
import reducer from './reducers/reducer';

export type StoreType = ReturnType<typeof reducer>

// 定义持久化配置
const persistConfig: PersistConfig<StoreType> = {
    key: "root",
    storage,
    whitelist: ["userReducer"]
}

// 创建持久化的 reducer
// @ts-ignore
const persistedReducer = persistReducer(persistConfig, reducer)

// 创建 store
const store = configureStore({
    reducer: persistedReducer,
    // devTools: process.env.NODE_ENV !== 'production',

    // middleware: (getDefaultMiddleware) =>
    //     getDefaultMiddleware({
    //         serializableCheck: {
    //             ignoredActions: ['persist/PERSIST', 'persist/REHYDRATE'],        // 忽略 redux-persist 的特定动作
    //         },
    //     }),
})

// @ts-ignore
const persistor = persistStore(store)

export { store, persistor }