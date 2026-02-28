"use client";

interface QuestionNumberGridProps {
  total: number;
  currentIndex: number;
  answeredIndices: Set<number>;
  onSelect: (index: number) => void;
}

export default function QuestionNumberGrid({
  total,
  currentIndex,
  answeredIndices,
  onSelect,
}: QuestionNumberGridProps) {
  return (
    <div className="flex flex-wrap gap-2">
      {Array.from({ length: total }, (_, i) => {
        const num = i + 1;
        const isCurrent = i === currentIndex;
        const isAnswered = answeredIndices.has(i);
        return (
          <button
            key={num}
            type="button"
            onClick={() => onSelect(i)}
            className={`w-10 h-10 rounded-lg font-medium text-sm transition-colors ${
              isCurrent
                ? "bg-blue-600 text-white ring-2 ring-blue-300"
                : isAnswered
                  ? "bg-green-100 text-green-800 hover:bg-green-200"
                  : "bg-gray-100 text-gray-700 hover:bg-gray-200"
            }`}
          >
            {num}
          </button>
        );
      })}
    </div>
  );
}
