import { cn } from "@/lib/utils";
import { Geist } from "next/font/google";
import React, { ElementType, ReactNode } from "react";
import { FlickeringGrid } from "@/components/magicui/flickering-grid";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const FeatureCard = ({
  title,
  description,
  icon,
}: {
  title: string;
  description: string;
  icon: ReactNode;
}) => {
  return (
    <div className="border border-zinc-700 rounded-md p-4 transition duration-300 hover:bg-zinc-800 z-10">
      <div className="flex items-start gap-3">
        {icon && <div className="text-zinc-400">{icon}</div>}
        <div>
          <h3 className="text-white font-medium mb-2">{title}</h3>
          <p className="text-zinc-400 text-sm">{description}</p>
        </div>
      </div>
    </div>
  );
};

const Features = () => {
  const featuresList = [
    {
      title: "Secure Connections",
      description: "End-to-end encrypted proxy connections with zero logging.",
      icon: <span>🔒</span>,
    },
    {
      title: "Global Network",
      description:
        "Access points in over 50 countries for reliable connections.",
      icon: <span>🌐</span>,
    },
    {
      title: "Custom Rules",
      description:
        "Create domain-specific routing rules to optimize your browsing.",
      icon: <span>⚙️</span>,
    },
    {
      title: "Low Latency",
      description:
        "Optimized for speed with minimal impact on connection quality.",
      icon: <span>⚡</span>,
    },
  ];

  return (
    <>
      <section
        className={cn(
          geistSans.className,
          "border-t border-b border-zinc-700 mt-3 p-2"
        )}
      >
        <div className="flex flex-col items-center justify-center relative">
          <FlickeringGrid
            className="relative w-full inset-0 z-0 [mask-image:linear-gradient(to_bottom,white,transparent)]"
            squareSize={2}
            gridGap={4}
            color="#BAC4C8"
            maxOpacity={0.5}
            flickerChance={0.1}
            height={100}
          />
          <h1 className="text-white font-medium text-xl absolute z-10">
            Features
          </h1>
        </div>
      </section>
      <section
        className={cn(geistSans.className, "border-b border-zinc-700 mt-3 p-5")}
      >
        <div className="flex flex-col items-center justify-center">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4 w-full max-w-4xl">
            {featuresList.map((feature, index) => (
              <FeatureCard
                key={index}
                title={feature.title}
                description={feature.description}
                icon={feature.icon}
              />
            ))}
          </div>
        </div>
      </section>
    </>
  );
};

export default Features;
