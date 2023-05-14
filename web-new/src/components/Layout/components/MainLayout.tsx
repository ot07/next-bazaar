import { Header } from "./Header";
import Head from "next/head";
import { type ReactNode } from "react";
import { AppShell } from "@mantine/core";

interface MainLayoutProps {
  children: ReactNode;
}

export function MainLayout({ children }: MainLayoutProps) {
  return (
    <>
      <Head>
        <title>Next Bazaar</title>
        <meta name="description" content="Generated by create-t3-app" />
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <AppShell header={<Header />}>
        <main>{children}</main>
      </AppShell>
    </>
  );
}
