export default function LoadingSpinner({ message = "Loading" }) {
  return (
    <div className="flex flex-col items-center justify-center">
      <div className="animate-spin rounded-full h-16 w-16 border-t-2 border-blue-400" />
      <div>{message}</div>
    </div>
  );
}

