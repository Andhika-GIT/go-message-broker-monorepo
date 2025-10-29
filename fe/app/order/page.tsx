import { DataTable } from "@/components/data-table";

import { columns } from "./column";
import { UploadSection } from "@/components/molecules";
import { getAllOrders, UploadOrderExcel } from "../action/order";

export default async function Page() {
  return (
    <>
      <UploadSection uploadFn={UploadOrderExcel} />
      <DataTable columns={columns} fetchFunction={getAllOrders}/>
    </>
  );
}
