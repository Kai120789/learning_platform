import { LoginForm } from "@/features/LoginForm";

export default function LoginPage() {
    return (
        <div className="flex min-h-svh flex-col items-center justify-center p-4 md:p-8">
            <div className="w-full max-w-sm md:max-w-4xl">
                <LoginForm />
            </div>
        </div>
    )
}
