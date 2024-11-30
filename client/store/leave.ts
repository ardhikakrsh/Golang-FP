import { createSlice } from "@reduxjs/toolkit";

export interface LeaveDTO {
  id: number;
  type: string;
  status: "requested" | "approved" | "declined";
  time_start: string;
  time_end: string;
  detail: string;
  created_at: string;
  updated_at: string;
  leave_day: number;
}
export interface LeaveReponseWithCountDTO {
  leave: LeaveDTO[];
  count_sick: number;
  count_business: number;
  count_vacation: number;
}

const leaveSlice = createSlice({
  name: "leave",
  initialState: {
    leave: [] as LeaveDTO[] | null,
  },
  reducers: {
    setLeave: (state, action) => {
      state.leave = action.payload;
    },
    addLeave: (state, action) => {
      return;
    },
  },
});

export const { setLeave, addLeave } = leaveSlice.actions;
export default leaveSlice.reducer;
