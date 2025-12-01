"use client";

import { IconBell } from "@tabler/icons-react";
import { useEffect, useRef, useState } from "react";

export const Notification = () => {
  const [messages, setMessages] = useState<string[]>([]);
  const [isConnected, setIsConnected] = useState(false);
  const wsRef = useRef<WebSocket | null>(null);

  useEffect(() => {
    // URL WebSocket kamu
    const wsUrl = "ws://localhost:9080"; // atau 'ws://localhost:9080/ws' tergantung konfig
    const ws = new WebSocket(wsUrl);
    wsRef.current = ws;

    ws.onopen = () => {
      console.log("WebSocket Connected!");
      setIsConnected(true);
      setMessages((prev) => [...prev, "Connected to WebSocket server"]);
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
      setMessages((prev) => [...prev, "❌ Connection error"]);
    };

    ws.onclose = () => {
      console.log("WebSocket Disconnected");
      setIsConnected(false);
      setMessages((prev) => [...prev, "Disconnected from server"]);
    };

    // Cleanup
    return () => {
      if (wsRef.current) {
        wsRef.current.close();
      }
    };
  }, []);
  return <div><IconBell size={28} /></div>;
};

