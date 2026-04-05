import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "Rafty",
  description: "A modern social network MVP built with Next.js and Go",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="pt-BR">
      <body>{children}</body>
    </html>
  );
}