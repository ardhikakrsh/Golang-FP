import Image from "next/image";
import { Inter } from "next/font/google";
import Nav from "../components/Nav";
import { use, useEffect } from "react";
import LeaveCount from "../components/LeaveCount";
import LeaveTable from "../components/LeaveTable";
import AddLeaveBtn from "../components/AddLeaveBtn";
import { useSelector } from "react-redux";
import { useRouter } from "next/router";

export default function Home() {
  const { user } = useSelector((state: any) => state.user);
  const router = useRouter();
  useEffect(() => {
    const { user } = JSON.parse(localStorage.getItem("user") || "{}");
    if (!user) {
      router.push("/login");
    }
  }, []);
  return (
    <div className="p-2 flex flex-col gap-10 items-center">
      <Nav />
      <div className="flex flex-col gap-10 w-8/12">
        <div className="flex flex-col gap-4">
          <h1 className="text-4xl font-bold text-center">
            Welcome <span className="text-blue-700">{user?.first_name}</span> 👋🏻
          </h1>
          <h1 className="text-xl font-normal text-center">
            Riwayat Total Cuti
          </h1>
        </div>
        <LeaveCount />
        <AddLeaveBtn />
        <LeaveTable />
      </div>
    </div>
  );
}
