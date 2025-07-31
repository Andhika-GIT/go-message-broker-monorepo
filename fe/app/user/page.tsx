import { DataTable } from "@/components/data-table";

import user from "./user.json";
import { columns } from "./column";

export default function Page() {
  return (
    <>
      <DataTable columns={columns} data={user} />
    </>
  );
}
