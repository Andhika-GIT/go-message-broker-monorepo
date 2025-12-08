"use client";

import { IconBell, IconBellFilled } from "@tabler/icons-react";
import { useEffect, useRef, useState } from "react";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";

export const Notification = () => {
  const [messages, setMessages] = useState<string[]>([]);
  const [isConnected, setIsConnected] = useState(false);
  const wsRef = useRef<WebSocket | null>(null);

  useEffect(() => {
    const wsUrl = "ws://localhost:9080";
    const ws = new WebSocket(wsUrl);
    wsRef.current = ws;

    ws.onopen = () => {
      console.log("WebSocket Connected!");
      setIsConnected(true);
    };

    ws.onmessage = (event) => {
      console.log("Message received:", event.data);
      try {
        const data = JSON.parse(event.data);
        setMessages((prev) => [...prev, `📨 ${JSON.stringify(data)}`]);
      } catch {
        setMessages((prev) => [...prev, `📨 ${event.data}`]);
      }
    };

    ws.onerror = (error) => {
      console.error("WebSocket Error:", error);
    };

    ws.onclose = () => {
      console.log("WebSocket Disconnected");
      setIsConnected(false);
    };

    // Cleanup
    return () => {
      if (wsRef.current) {
        wsRef.current.close();
      }
    };
  }, []);

  if (messages.length < 1) return (
    <div><IconBell size={28} /></div>
  )
  return (
    <Popover>
      <PopoverTrigger asChild>
        <div>
          <IconBellFilled size={28} />
        </div>
      </PopoverTrigger>
      <PopoverContent className="w-80">
        <div className="grid gap-4">
          {messages.length > 0 &&
            messages.map((message) => (
              <div className="grid gap-2">
                <div className="p-2">
                  <p className="text-sm">{message}</p>
                </div>
              </div>
            ))}
        </div>
      </PopoverContent>
    </Popover>
  );
};
