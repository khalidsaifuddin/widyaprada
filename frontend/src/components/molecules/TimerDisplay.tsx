"use client";

interface TimerDisplayProps {
  secondsRemaining: number;
  onTimeUp?: () => void;
  warningThreshold?: number;
}

export default function TimerDisplay({
  secondsRemaining,
  warningThreshold = 300,
}: TimerDisplayProps) {
  const mins = Math.floor(secondsRemaining / 60);
  const secs = secondsRemaining % 60;
  const isWarning = secondsRemaining <= warningThreshold && secondsRemaining > 0;

  return (
    <div
      className={`inline-flex items-center gap-2 px-4 py-2 rounded-lg font-mono text-lg font-semibold ${
        isWarning ? "bg-amber-100 text-amber-800" : "bg-gray-100 text-gray-800"
      }`}
    >
      <span aria-label="Waktu tersisa">
        {String(mins).padStart(2, "0")}:{String(secs).padStart(2, "0")}
      </span>
    </div>
  );
}
