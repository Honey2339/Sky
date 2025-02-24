import Features from "@/components/Features";
import { Button } from "@/components/ui/button";
import SplineComponent from "@/components/ui/Spline";
import { cn } from "@/lib/utils";
import { Geist } from "next/font/google";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

export default function Home() {
  return (
    <section className="">
      <section className="p-10">
        <div className="flex relative">
          <div className="max-w-md mt-6 relative z-10">
            <h1 className="text-white font-serif text-7xl font-medium">Sky</h1>
            <p className="text-zinc-300 font-serif text-xl font-medium mt-8">
              High-performance reverse proxy built with Go, designed for
              developers who demand speed, flexibility, and security.
            </p>
            <Button variant="secondary" size="sm" className="mt-5 font-medium">
              {">_"}Copy Cmd
            </Button>
          </div>

          <div className="absolute -top-16 -right-60 w-[800px] h-[800px]">
            <SplineComponent />
          </div>
        </div>
      </section>
      <section
        className={cn(
          geistSans.className,
          "border-t border-zinc-700 mt-3 p-10"
        )}
      >
        <div className="flex flex-col items-center justify-center">
          <h1 className="text-white font-medium text-xl">Why Reverse Proxy?</h1>
          <p className="text-zinc-300/90 font-medium text-sm text-center mt-5">
            They enhance security by concealing the identities of origin
            servers. This protects them from direct attacks on the public
            internet. Reverse proxies also optimize performance through
            efficient content management and caching, leading to faster loading
            times and an improved user experience.
          </p>
        </div>
      </section>
      <Features />
    </section>
  );
}
