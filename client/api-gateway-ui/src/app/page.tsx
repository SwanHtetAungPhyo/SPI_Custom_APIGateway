'use client'

import React, { useState, useEffect } from 'react'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { Accordion, AccordionContent, AccordionItem, AccordionTrigger } from "@/components/ui/accordion"
import { Badge } from "@/components/ui/badge"
import { Button } from "@/components/ui/button"
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from "@/components/ui/tooltip"
import { Globe, Server, Users, DollarSign, PieChart, TrendingUp, ChevronRight, Clock, RefreshCcw, AlertCircle, Info, Loader2 } from 'lucide-react'

type ApiGatewayConfig = {
    Name: string;
    Version: string;
    Description: string;
    DefaultRoute: string;
    GateWayInfo: string;
    LoadBalancing: string;
    MainApp: string;
    Services: Array<{
        Name: string;
        URL: string;
        Leader: string;
        Instance: number[];
        Routes: Array<{
            Description: string;
            Path: string[];
            Method: string[];
            Timeout: string;
            Retries: number;
            GeneratedRoute: string[];
        }>;
    }>;
};

const ServiceIcon = ({ name }: { name: string }) => {
    switch (name.toLowerCase()) {
        case 'user':
            return <Users className="h-5 w-5" />
        case 'expense':
            return <DollarSign className="h-5 w-5" />
        case 'budgets':
            return <PieChart className="h-5 w-5" />
        case 'tracking':
            return <TrendingUp className="h-5 w-5" />
        case 'investment':
            return <TrendingUp className="h-5 w-5" />
        default:
            return <Server className="h-5 w-5" />
    }
}

export default function ApiGatewayUI() {
    const [apiGatewayConfig, setApiGatewayConfig] = useState<ApiGatewayConfig | null>(null)
    const [activeTab, setActiveTab] = useState<string>('')
    const [loading, setLoading] = useState(true)
    const [error, setError] = useState<string | null>(null)

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch('http://localhost:8081/gate/services')
                if (!response.ok) {
                    throw new Error('Failed to fetch API Gateway configuration')
                }
                const data: ApiGatewayConfig = await response.json()
                setApiGatewayConfig(data)
                setActiveTab(data.Services[0].Name)
                setLoading(false)
            } catch (err :unknown) {
                // @ts-expect-error Ignoring error because err.data may not exist
                setError(err.data)
                setLoading(false)
            }
        }

        fetchData()
    }, [])

    if (loading) {
        return (
            <div className="flex items-center justify-center min-h-screen bg-gray-50">
                <Loader2 className="h-12 w-12 animate-spin text-blue-500" />
            </div>
        )
    }

    if (error) {
        return (
            <div className="flex items-center justify-center min-h-screen bg-gray-50">
                <Card className="w-full max-w-md">
                    <CardHeader>
                        <CardTitle className="text-2xl text-red-500 flex items-center gap-2">
                            <AlertCircle className="h-6 w-6" />
                            Error
                        </CardTitle>
                    </CardHeader>
                    <CardContent>
                        <p className="text-gray-600">{error}</p>
                        <Button className="mt-4" onClick={() => window.location.reload()}>
                            Retry
                        </Button>
                    </CardContent>
                </Card>
            </div>
        )
    }

    if (!apiGatewayConfig) return null

    return (
        <TooltipProvider>
            <div className="container mx-auto p-4 min-h-screen bg-gray-50">
                <Card className="mb-8 shadow-sm border-t-4 border-blue-500">
                    <CardHeader>
                        <div className="flex items-center justify-between">
                            <CardTitle className="text-3xl text-gray-900 flex items-center gap-2">
                                <Server className="h-8 w-8 text-blue-500" />
                                {apiGatewayConfig.Name}
                            </CardTitle>
                            <Badge variant="outline" className="text-lg">
                                {apiGatewayConfig.Version}
                            </Badge>
                        </div>
                        <CardDescription className="text-lg mt-2">{apiGatewayConfig.Description}</CardDescription>
                    </CardHeader>
                    <CardContent>
                        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                            <div className="space-y-3">
                                <p className="flex items-center gap-2 text-gray-700">
                                    <Globe className="h-5 w-5 text-blue-500" />
                                    <span className="font-medium">Default Route:</span> {apiGatewayConfig.DefaultRoute}
                                </p>
                                <p className="flex items-center gap-2 text-gray-700">
                                    <Server className="h-5 w-5 text-blue-500" />
                                    <span className="font-medium">Gateway Info:</span> {apiGatewayConfig.GateWayInfo}
                                </p>
                            </div>
                            <div className="space-y-3">
                                <p className="flex items-center gap-2 text-gray-700">
                                    <RefreshCcw className="h-5 w-5 text-blue-500" />
                                    <span className="font-medium">Load Balancing:</span> {apiGatewayConfig.LoadBalancing}
                                </p>
                                <p className="flex items-center gap-2 text-gray-700">
                                    <Users className="h-5 w-5 text-blue-500" />
                                    <span className="font-medium">Main App:</span> {apiGatewayConfig.MainApp}
                                </p>
                            </div>
                        </div>
                    </CardContent>
                </Card>

                <Tabs value={activeTab} onValueChange={setActiveTab}>
                    <TabsList className="mb-6 flex flex-wrap justify-start gap-2 bg-transparent">
                        {apiGatewayConfig.Services.map((service) => (
                            <TabsTrigger
                                key={service.Name}
                                value={service.Name}
                                className="data-[state=active]:bg-blue-100 data-[state=active]:text-blue-700 border border-gray-200"
                            >
                                <ServiceIcon name={service.Name} />
                                <span className="ml-2">{service.Name}</span>
                            </TabsTrigger>
                        ))}
                    </TabsList>

                    {apiGatewayConfig.Services.map((service) => (
                        <TabsContent key={service.Name} value={service.Name}>
                            <Card>
                                <CardHeader>
                                    <CardTitle className="text-2xl text-gray-900 flex items-center gap-2">
                                        <ServiceIcon name={service.Name} />
                                        {service.Name} Service
                                    </CardTitle>
                                    <CardDescription>
                                        <p className="text-sm"><span className="font-medium">URL:</span> {service.URL}</p>
                                        <p className="text-sm"><span className="font-medium">Leader:</span> {service.Leader}</p>
                                        <p className="text-sm"><span className="font-medium">Instances:</span> {service.Instance.join(', ')}</p>
                                    </CardDescription>
                                </CardHeader>
                                <CardContent>
                                    <Accordion type="single" collapsible className="space-y-2">
                                        {service.Routes.map((route, idx) => (
                                            <AccordionItem key={idx} value={`route-${idx}`}>
                                                <AccordionTrigger className="hover:bg-gray-100 rounded-lg px-4">
                          <span className="flex items-center gap-2">
                            <Globe className="h-5 w-5 text-blue-500" />
                              {route.Description}
                          </span>
                                                </AccordionTrigger>
                                                <AccordionContent className="px-4 py-2 bg-gray-50 rounded-lg mt-2">
                                                    <div className="space-y-2">
                                                        <p className="flex items-center gap-2">
                                                            <ChevronRight className="h-4 w-4 text-blue-500" />
                                                            <span className="font-medium">Path:</span> {route.Path.join(', ')}
                                                        </p>
                                                        <p className="flex items-center gap-2">
                                                            <ChevronRight className="h-4 w-4 text-blue-500" />
                                                            <span className="font-medium">Method:</span> {route.Method.join(', ')}
                                                        </p>
                                                        <p className="flex items-center gap-2">
                                                            <Clock className="h-4 w-4 text-blue-500" />
                                                            <span className="font-medium">Timeout:</span> {route.Timeout}
                                                        </p>
                                                        <p className="flex items-center gap-2">
                                                            <RefreshCcw className="h-4 w-4 text-blue-500" />
                                                            <span className="font-medium">Retries:</span> {route.Retries}
                                                        </p>
                                                        <div>
                                                            <p className="font-medium mb-1">Generated Routes:</p>
                                                            <ul className="list-disc list-inside space-y-1 text-sm text-gray-600">
                                                                {route.GeneratedRoute.map((genRoute, i) => (
                                                                    <li key={i}>{genRoute}</li>
                                                                ))}
                                                            </ul>
                                                        </div>
                                                    </div>
                                                </AccordionContent>
                                            </AccordionItem>
                                        ))}
                                    </Accordion>
                                </CardContent>
                            </Card>
                        </TabsContent>
                    ))}
                </Tabs>

                <div className="mt-8 text-center text-sm text-gray-500">
                    <Tooltip>
                        <TooltipTrigger asChild>
                            <Button variant="link" className="text-blue-500 hover:text-blue-700">
                                <Info className="h-5 w-5 mr-2" />
                                About API Gateway Configuration
                            </Button>
                        </TooltipTrigger>
                        <TooltipContent>
                            <p>This UI displays the configuration of your API Gateway.</p>
                            <p>Click on the service tabs to view details for each service.</p>
                        </TooltipContent>
                    </Tooltip>
                </div>
            </div>
        </TooltipProvider>
    )
}