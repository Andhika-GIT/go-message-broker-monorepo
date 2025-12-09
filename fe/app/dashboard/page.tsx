import { ChartAreaInteractive } from "@/components/chart-area-interactive";
import { DataTable } from "@/components/data-table";
import { SectionCards } from "@/components/section-cards";

import data from "./data.json";
import { columns } from './column'
import { getDasboardData } from "../action/dashboard";

export default async function Page() {

  const data = await getDasboardData()

  console.log(data)

  return (
    <>
      <SectionCards />
      {/* <div className="px-4 lg:px-6">
        <ChartAreaInteractive />
      </div> */}
    </>
  );
}
