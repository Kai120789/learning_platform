import { FaRegCheckCircle } from "react-icons/fa";
import { cn } from "@/shared/lib/utils";
import type { ReactNode } from "react";

export type Step = {
    id: number
    icon: ReactNode
};

type StepperProps = {
    steps: Step[];
    currentStep: number;
    onStepClick?: (step: number) => void;
};

export function Stepper({
    steps,
    currentStep,
    onStepClick,
}: StepperProps) {
    return (
        <div className="flex w-full items-start">
            {steps.map((step, index) => {
                const completed = step.id < currentStep;
                const active = step.id === currentStep;

                return (
                    <div
                        key={step.id}
                        className={cn(
                            "flex flex-col",
                            index !== steps.length - 1 && "flex-1"
                        )}
                    >
                        <div className="flex w-full items-center">
                            <button
                                onClick={() => onStepClick?.(step.id)}
                                className={cn(
                                    "z-10 flex h-10 w-10 items-center justify-center rounded-lg border-2 text-sm font-semibold transition-all duration-300",
                                    completed &&
                                    "border-primary bg-primary text-primary-foreground",
                                    active &&
                                    "border-primary bg-background text-primary",
                                    !completed &&
                                    !active &&
                                    "border-border bg-background text-muted-foreground",
                                    step.id > currentStep &&
                                    "cursor-not-allowed"
                                )}
                                disabled={step.id > currentStep}
                            >
                                {completed ? <FaRegCheckCircle className="h-5 w-5" /> : step.icon}
                            </button>

                            {index !== steps.length - 1 && (
                                <div
                                    className={cn(
                                        "h-0.5 flex-1 transition-colors duration-300",
                                        completed ? "bg-primary" : "bg-border"
                                    )}
                                />
                            )}
                        </div>
                    </div>
                );
            })}
        </div>
    );
}