"use client";

import JurnalDetail from "@/components/organisms/JurnalDetail";
import { useParams } from "next/navigation";

export default function JurnalDetailPage() {
  const params = useParams();
  const id = (params?.id as string) ?? "";
  return <JurnalDetail id={id} />;
}
