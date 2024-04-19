import { cn } from "@/lib/utils";
import type { LucideIcon } from "lucide-react";
import { SparkleIcon } from "./svg/template-page";

type Props = {
  className?: string;
  IconLeft?: LucideIcon;
  label: string;
  IconRight?: LucideIcon;
  onClick?: React.MouseEventHandler<HTMLButtonElement>;
  shiny?: boolean;
};

export const PrimaryButton: React.FC<Props> = ({
  className,
  IconLeft,
  label,
  IconRight,
  onClick,
  shiny = false,
}) => {
  return (
    <div className="relative group/button">
      <div
        aria-hidden
        className="absolute -inset-0.5 bg-white rounded-lg blur-2xl group-hover/button:opacity-30 transition duration-300  opacity-0 "
      />
      <button
        onClick={onClick}
        type="button"
        className={cn(
          "relative flex items-center px-4 gap-2 text-sm font-semibold text-black group-hover:bg-white/90 duration-1000 rounded-lg h-10",
          {
            "bg-white": !shiny,
            "bg-gradient-to-r from-white/80 to-white": shiny,
          },
          className,
        )}
      >
        {IconLeft ? <IconLeft className="w-4 h-4" /> : null}
        {label}
        {IconRight ? <IconRight className="w-4 h-4" /> : null}

        {shiny && (
          <div
            aria-hidden
            className="pointer-events-none absolute inset-0 opacity-0 group-hover/button:[animation-delay:.2s] group-hover/button:animate-button-shine rounded-[inherit] bg-[length:200%_100%] bg-[linear-gradient(110deg,transparent,35%,rgba(255,255,255,.7),75%,transparent)]"
          />
        )}
      </button>
    </div>
  );
};

export const SecondaryButton: React.FC<Props> = ({ className, IconLeft, label, IconRight }) => {
  return (
    <button
      type="button"
      className={cn(
        "items-center gap-2 px-4 duration-500 text-white/70 hover:text-white h-10 flex",
        className,
      )}
    >
      {IconLeft ? <IconLeft className="w-4 h-4" /> : null}
      {label}
      {IconRight ? <IconRight className="w-4 h-4" /> : null}
    </button>
  );
};

export const RainbowDarkButton: React.FC<Props> = ({ className, label, IconRight }) => {
  return (
    <div
      className={cn(
        "p-[.75px] hero-hiring-gradient rounded-full w-fit mx-auto relative z-50",
        className,
      )}
    >
      <button
        type="button"
        className="items-center gap-4 px-3 py-1.5 bg-black text-white rounded-full flex flex-block text-sm"
      >
        <SparkleIcon className="text-white" />
        {label}
        {IconRight ? <IconRight className="w-4 h-4" /> : null}
      </button>
    </div>
  );
};
