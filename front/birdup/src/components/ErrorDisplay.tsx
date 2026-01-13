
export interface ErrorDisplayProps {
    text: string
}

export default function ErrorDisplay({ text }: ErrorDisplayProps) {
    return (
        <div className="rounded-md border-red-500 bg-red-300">
            {text}
        </div>
    );
}