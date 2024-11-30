import React, { use, useEffect } from "react";
import { useSelector, useDispatch } from "react-redux";
import { LeaveDTO, setLeave } from "../store/leave";
import axios from "axios";
function LeaveTable() {
  const { leave } = useSelector((state: any) => state.leave);
  const { user } = useSelector((state: any) => state.user);

  const dispatch = useDispatch();
  useEffect(() => {
    const fetchLeave = async () => {
      const { data } = await axios.get("http://localhost:8000/leaves/me", {
        withCredentials: true,
      });
      dispatch(setLeave(data));
    };
    if (!!!user) {
      fetchLeave();
    }
  }, []);

  return (
    <table className="table bg-base-100 overflow-hidden shadow-md text-zinc-300 text-center">
      {/* head */}
      <thead>
        <tr className="text-lg text-white bg-base-300">
          <th></th>
          <th>Jenis</th>
          <th>Jumlah Hari Cuti</th>
          <th>Mulai Cuti</th>
          <th>Akhir Cuti</th>
          <th>Deskripsi</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody className="text-md font-normal">
        {/* row 1 */}
        {leave &&
          leave?.leave_responses?.map((item: LeaveDTO, index: number) => {
            let typeLabel = "";
            if (item.type == "sakit") {
              typeLabel = "Sakit";
            } else if (item.type == "absen") {
              typeLabel = "Absen";
            } else if (item.type == "liburan") {
              typeLabel = "Liburan";
            }

            let statusColor = "";
            let statusLabel = "";
            let icon = "";

            if (item.status === "requested") {
              statusColor = "rounded-full text-white bg-blue-600 py-1 px-3";
              statusLabel = "Requested";
              icon = "⏳";
            } else if (item.status === "approved") {
              statusColor = "rounded-full text-white bg-green-600 py-1 px-3";
              statusLabel = "Approved";
              icon = "✔️";
            } else if (item.status === "declined") {
              statusColor = "rounded-full text-white bg-red-600 py-1 px-3";
              statusLabel = "Declined";
              icon = "❌";
            }

            let date = new Date(item.created_at);
            // let newDate =
            //   date.getDate() +
            //   "/" +
            //   (date.getMonth() + 1) +
            //   "/" +
            //   date.getFullYear();

            let startDate = new Date(item.time_start);
            let newStartDate =
              startDate.getDate() +
              "/" +
              (startDate.getMonth() + 1) +
              "/" +
              startDate.getFullYear();

            let endDate = new Date(item.time_end);
            let newEndDate =
              endDate.getDate() +
              "/" +
              (endDate.getMonth() + 1) +
              "/" +
              endDate.getFullYear();

            return (
              <tr key={index} className="">
                <td>{index + 1}</td>
                <td>
                  <p>{typeLabel}</p>
                </td>
                <td>{item.leave_day}</td>
                <td>{newStartDate}</td>
                <td>{newEndDate}</td>
                <td>{item.detail}</td>
                <td>
                  <span className={`${statusColor} font-medium`}>
                    {icon} {statusLabel}
                  </span>
                </td>
              </tr>
            );
          })}
      </tbody>
    </table>
  );
}

export default LeaveTable;
