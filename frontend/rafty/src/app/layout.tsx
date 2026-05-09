import type { Metadata } from "next";
import "./globals.css";

import { QueryProvider } from "@/providers/query-provider";

export const metadata: Metadata = {
  title: "Rafty",
  description: "Flow through the internet.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="bg-[#071018] text-white antialiased">
        <QueryProvider>{children}</QueryProvider>
      </body>
    </html>
  );
}