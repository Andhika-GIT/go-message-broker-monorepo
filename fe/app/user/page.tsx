import { DataTable } from "@/components/data-table";

import data from "../dashboard/data.json";

export default function Page() {
  return (
    <>
      <DataTable data={data} />
    </>
  );
}
