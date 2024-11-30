import { createSlice } from "@reduxjs/toolkit";



export interface UserInfo {
    username: string
    password: string
    firstName: string
    lastName: string
}

const initialState = {
    user: null as UserInfo | null,
}
const userSlice = createSlice({
    name: "user",
    initialState,
    reducers: {
        login: (state, action) => {
            state.user = action.payload;
        },
        logout: (state) => {
            state.user = null;
        },
    },
});

export const { login, logout } = userSlice.actions;
export default userSlice.reducer;   