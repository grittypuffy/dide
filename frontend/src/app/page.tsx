import App from "@/components/App";
import {HeroUIProvider} from "@heroui/react";

export default function Home() {
    return (
        <HeroUIProvider>
          <App />
        </HeroUIProvider>
      );
}
