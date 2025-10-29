import { DataTable } from "@/components/data-table";

import { columns } from "./column";
import { UploadSection } from "@/components/molecules";
import { getAllOrders } from "../action/order";

export default async function Page() {
  return (
    <>
      {/* <UploadSection uploadFn={UploadUserExcel} /> */}
      <DataTable columns={columns} fetchFunction={getAllOrders}/>
    </>
  );
}
