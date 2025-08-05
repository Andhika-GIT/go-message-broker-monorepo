import { DataTable } from "@/components/data-table";

import user from "./user.json";
import { columns } from "./column";
import { UploadSection } from "@/components/molecules";

export default function Page() {
  return (
    <>
      <UploadSection />
      <DataTable columns={columns} data={user} />
    </>
  );
}
