import { SectionCards } from "@/components/section-cards";

import { getDasboardData } from "../action/dashboard";

export default async function Page() {

  const data = await getDasboardData()

  return (
    <>
      <SectionCards total_orders={data.total_orders} total_users={data.total_users}/>
      {/* <div className="px-4 lg:px-6">
        <ChartAreaInteractive />
      </div> */}
    </>
  );
}
