import { DataTable } from "@/components/data-table";

import user from "./user.json";
import { columns } from "./column";
import { UploadSection } from "@/components/molecules";
import { UploadUserExcel } from "../action/user";

export default function Page() {
  return (
    <>
      <UploadSection uploadFn={UploadUserExcel} />
      <DataTable columns={columns} data={user} />
    </>
  );
}
