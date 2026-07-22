import { Button } from "@workspace/ui/components/button"

export default function Page() {
  return (
    <main className="flex min-h-screen items-center justify-center">
      <section className="max-w-2xl text-center space-y-6">
        <h1 className="text-5xl font-bold tracking-tight">
          Kairo
        </h1>

        <p className="text-lg text-muted-foreground">
          Software should adapt to your business, not the other way around.
        </p>

        <Button>Get Started</Button>
      </section>
    </main>
  )
}