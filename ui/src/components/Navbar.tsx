import Link from "next/link";
import React from "react";
import Logo from "./logo.png";
import { Button } from "./ui/button";
import Image from "next/image";
import { Geist } from "next/font/google";
import { cn } from "@/lib/utils";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const Navbar = () => {
  return (
    <header
      className={cn(
        geistSans.className,
        `sticky top-0 z-50 w-full border-b border-zinc-700 bg-zinc-900 backdrop-blur`
      )}
    >
      <div className="container flex h-12 max-w-screen-2xl items-center">
        <Link href="/" className="flex items-center space-x-2">
          {/* <Image src={Logo} height={100} width={50} alt="" /> */}
          {/* <span className="font-semibold">Sky</span> */}
        </Link>
        <div className="flex flex-1 items-center justify-end space-x-4">
          <nav className="flex items-center space-x-6 ">
            <Link
              href="/docs"
              className="text-sm text-white font-medium transition duration-200 hover:underline"
            >
              Copy Cmd
            </Link>
            <Button variant="secondary" size="sm" className="" asChild>
              <Link href="/dashboard" className="font-medium">
                Contact
              </Link>
            </Button>
          </nav>
        </div>
      </div>
    </header>
  );
};

export default Navbar;
