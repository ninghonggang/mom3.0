--
-- PostgreSQL database dump
--

\restrict NjEIdbnFSsNlsfQSLqiIFhLECW00clQSU4mt2jafdZCf6rRWgmLIH6QHDhBj390

-- Dumped from database version 15.14
-- Dumped by pg_dump version 15.14

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: ai_chat_conversation; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ai_chat_conversation (
    id bigint NOT NULL,
    tenant_id bigint NOT NULL,
    user_id bigint NOT NULL,
    session_id uuid,
    title character varying(200),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.ai_chat_conversation OWNER TO postgres;

--
-- Name: ai_chat_conversation_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ai_chat_conversation_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.ai_chat_conversation_id_seq OWNER TO postgres;

--
-- Name: ai_chat_conversation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ai_chat_conversation_id_seq OWNED BY public.ai_chat_conversation.id;


--
-- Name: ai_chat_message; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ai_chat_message (
    id bigint NOT NULL,
    tenant_id bigint NOT NULL,
    user_id bigint NOT NULL,
    conversation_id bigint NOT NULL,
    role character varying(20),
    content text,
    intent_json text,
    operation_type character varying(20),
    status character varying(20) DEFAULT 'pending'::character varying,
    tool_result text,
    created_at timestamp with time zone
);


ALTER TABLE public.ai_chat_message OWNER TO postgres;

--
-- Name: ai_chat_message_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ai_chat_message_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.ai_chat_message_id_seq OWNER TO postgres;

--
-- Name: ai_chat_message_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ai_chat_message_id_seq OWNED BY public.ai_chat_message.id;


--
-- Name: ai_config; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ai_config (
    id bigint NOT NULL,
    tenant_id bigint NOT NULL,
    config_name character varying(100),
    provider character varying(50),
    endpoint character varying(500),
    api_version character varying(100),
    model_name character varying(100),
    api_key character varying(500),
    temperature numeric(3,2),
    max_tokens bigint,
    timeout bigint,
    enable boolean DEFAULT true,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.ai_config OWNER TO postgres;

--
-- Name: ai_config_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ai_config_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.ai_config_id_seq OWNER TO postgres;

--
-- Name: ai_config_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ai_config_id_seq OWNED BY public.ai_config.id;


--
-- Name: alert_escalation_rule; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.alert_escalation_rule (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    rule_code character varying(50) NOT NULL,
    rule_name character varying(200) NOT NULL,
    alert_type character varying(30),
    severity_level character varying(10),
    escalation_levels jsonb,
    is_enabled bigint DEFAULT 1,
    tenant_id bigint NOT NULL,
    updated_at timestamp with time zone
);


ALTER TABLE public.alert_escalation_rule OWNER TO postgres;

--
-- Name: alert_escalation_rule_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.alert_escalation_rule_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.alert_escalation_rule_id_seq OWNER TO postgres;

--
-- Name: alert_escalation_rule_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.alert_escalation_rule_id_seq OWNED BY public.alert_escalation_rule.id;


--
-- Name: alert_notification_log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.alert_notification_log (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    alert_id bigint NOT NULL,
    alert_no character varying(50),
    channel character varying(20) NOT NULL,
    receiver_type character varying(20),
    receiver_id bigint,
    receiver_name character varying(50),
    receiver_value character varying(200),
    notification_status character varying(20),
    sent_time timestamp with time zone,
    read_time timestamp with time zone,
    error_code character varying(50),
    error_msg character varying(500),
    retry_count bigint DEFAULT 0,
    tenant_id bigint NOT NULL
);


ALTER TABLE public.alert_notification_log OWNER TO postgres;

--
-- Name: alert_notification_log_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.alert_notification_log_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.alert_notification_log_id_seq OWNER TO postgres;

--
-- Name: alert_notification_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.alert_notification_log_id_seq OWNED BY public.alert_notification_log.id;


--
-- Name: alert_record; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.alert_record (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    alert_no character varying(50) NOT NULL,
    rule_id bigint,
    rule_code character varying(50),
    rule_name character varying(200),
    alert_type character varying(30) NOT NULL,
    severity_level character varying(10) NOT NULL,
    title character varying(200) NOT NULL,
    content text NOT NULL,
    trigger_time timestamp with time zone NOT NULL,
    source_module character varying(30),
    source_id bigint,
    source_no character varying(50),
    source_data jsonb,
    status character varying(20) DEFAULT 'TRIGGERED'::character varying,
    urgency_level character varying(10) DEFAULT 'NORMAL'::character varying,
    acknowledged_by bigint,
    acknowledged_by_name character varying(50),
    acknowledged_time timestamp with time zone,
    acknowledged_remark character varying(200),
    resolved_by bigint,
    resolved_by_name character varying(50),
    resolved_time timestamp with time zone,
    resolution_remark character varying(500),
    closed_time timestamp with time zone,
    escalation_count bigint DEFAULT 0,
    last_escalation_time timestamp with time zone,
    tenant_id bigint NOT NULL,
    updated_at timestamp with time zone
);


ALTER TABLE public.alert_record OWNER TO postgres;

--
-- Name: alert_record_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.alert_record_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.alert_record_id_seq OWNER TO postgres;

--
-- Name: alert_record_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.alert_record_id_seq OWNED BY public.alert_record.id;


--
-- Name: alert_rule; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.alert_rule (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    rule_code character varying(50) NOT NULL,
    rule_name character varying(200) NOT NULL,
    alert_type character varying(30) NOT NULL,
    biz_module character varying(30),
    condition_expression character varying(500) NOT NULL,
    condition_params jsonb,
    severity_level character varying(10) DEFAULT 'MEDIUM'::character varying,
    notification_channels jsonb,
    notify_templates jsonb,
    escalation_rule_id bigint,
    is_enabled bigint DEFAULT 1,
    check_interval bigint DEFAULT 60,
    max_trigger_count bigint DEFAULT 0,
    trigger_count bigint DEFAULT 0,
    last_trigger_time timestamp with time zone,
    tenant_id bigint NOT NULL,
    created_by character varying(50),
    updated_by character varying(50),
    updated_at timestamp with time zone
);


ALTER TABLE public.alert_rule OWNER TO postgres;

--
-- Name: alert_rule_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.alert_rule_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.alert_rule_id_seq OWNER TO postgres;

--
-- Name: alert_rule_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.alert_rule_id_seq OWNED BY public.alert_rule.id;


--
-- Name: andon_call; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.andon_call (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    call_no character varying(50) NOT NULL,
    workshop_id bigint,
    workshop_name character varying(100),
    production_line_id bigint,
    production_line_name character varying(100),
    workstation_id bigint,
    workstation_name character varying(100),
    andon_type character varying(30),
    andon_type_name character varying(50),
    call_level bigint DEFAULT 1,
    priority bigint DEFAULT 5,
    description character varying(500),
    photos jsonb,
    audio_url character varying(500),
    call_by character varying(50),
    call_time timestamp with time zone,
    response_by character varying(50),
    response_time timestamp with time zone,
    handle_by character varying(50),
    handle_time timestamp with time zone,
    handle_result character varying(20),
    handle_remark character varying(500),
    related_order_id bigint,
    related_ncr_id bigint,
    related_repair_id bigint,
    status character varying(20) DEFAULT 'CALLING'::character varying,
    is_escalated bigint DEFAULT 0,
    escalated_at timestamp with time zone,
    escalation_count bigint DEFAULT 0,
    response_duration bigint,
    handle_duration bigint
);


ALTER TABLE public.andon_call OWNER TO postgres;

--
-- Name: andon_call_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.andon_call_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.andon_call_id_seq OWNER TO postgres;

--
-- Name: andon_call_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.andon_call_id_seq OWNED BY public.andon_call.id;


--
-- Name: andon_calls; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.andon_calls (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    call_no character varying(50) NOT NULL,
    line_id bigint,
    line_name character varying(100),
    station_id bigint,
    station_name character varying(100),
    call_type character varying(20),
    call_level bigint DEFAULT 1,
    call_desc character varying(500),
    call_user_id bigint,
    call_user_name character varying(50),
    call_time timestamp(6) with time zone,
    response_user_id bigint,
    response_time timestamp(6) with time zone,
    resolve_time timestamp(6) with time zone,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.andon_calls OWNER TO postgres;

--
-- Name: andon_calls_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.andon_calls_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.andon_calls_id_seq OWNER TO postgres;

--
-- Name: andon_calls_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.andon_calls_id_seq OWNED BY public.andon_calls.id;


--
-- Name: andon_escalation_log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.andon_escalation_log (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    call_id bigint NOT NULL,
    from_level bigint NOT NULL,
    to_level bigint NOT NULL,
    escalation_type character varying(20),
    trigger_user character varying(50),
    trigger_reason character varying(200),
    notify_result jsonb
);


ALTER TABLE public.andon_escalation_log OWNER TO postgres;

--
-- Name: andon_escalation_log_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.andon_escalation_log_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.andon_escalation_log_id_seq OWNER TO postgres;

--
-- Name: andon_escalation_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.andon_escalation_log_id_seq OWNED BY public.andon_escalation_log.id;


--
-- Name: andon_escalation_rule; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.andon_escalation_rule (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    rule_code character varying(50) NOT NULL,
    rule_name character varying(100) NOT NULL,
    andon_type character varying(30),
    workshop_id bigint,
    priority_range character varying(20),
    is_default bigint DEFAULT 0,
    level1_timeout bigint NOT NULL,
    level1_notify_type character varying(20),
    level1_notify_json jsonb,
    level2_timeout bigint,
    level2_notify_type character varying(20),
    level2_notify_json jsonb,
    level3_timeout bigint,
    level3_notify_type character varying(20),
    level3_notify_json jsonb,
    level4_timeout bigint,
    level4_notify_type character varying(20),
    level4_notify_json jsonb,
    escalation_mode character varying(20) DEFAULT 'TIMEOUT'::character varying,
    max_escalation_level bigint DEFAULT 4,
    audio_enabled bigint DEFAULT 0,
    audio_message_template character varying(200),
    audio_repeat_times bigint DEFAULT 3,
    is_enabled bigint DEFAULT 1,
    sort_order bigint DEFAULT 0,
    remark character varying(500),
    tenant_id bigint NOT NULL,
    created_by character varying(50)
);


ALTER TABLE public.andon_escalation_rule OWNER TO postgres;

--
-- Name: andon_escalation_rule_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.andon_escalation_rule_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.andon_escalation_rule_id_seq OWNER TO postgres;

--
-- Name: andon_escalation_rule_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.andon_escalation_rule_id_seq OWNED BY public.andon_escalation_rule.id;


--
-- Name: andon_notification_log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.andon_notification_log (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    call_id bigint NOT NULL,
    channel character varying(20) NOT NULL,
    receiver_type character varying(20),
    receiver_id character varying(50),
    receiver_name character varying(100),
    title character varying(200),
    content text,
    priority bigint,
    send_time timestamp with time zone,
    send_result character varying(20),
    error_msg character varying(500),
    tenant_id bigint NOT NULL
);


ALTER TABLE public.andon_notification_log OWNER TO postgres;

--
-- Name: andon_notification_log_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.andon_notification_log_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.andon_notification_log_id_seq OWNER TO postgres;

--
-- Name: andon_notification_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.andon_notification_log_id_seq OWNED BY public.andon_notification_log.id;


--
-- Name: andon_report; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.andon_report (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    report_date timestamp with time zone NOT NULL,
    workshop_id bigint,
    workshop_name character varying(100),
    line_id bigint,
    line_name character varying(100),
    station_id bigint,
    station_name character varying(100),
    total_call_count bigint,
    material_call_count bigint,
    quality_call_count bigint,
    equipment_call_count bigint,
    other_call_count bigint,
    avg_response_time numeric,
    avg_resolve_time numeric,
    unresolved_count bigint,
    remark character varying(500)
);


ALTER TABLE public.andon_report OWNER TO postgres;

--
-- Name: andon_report_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.andon_report_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.andon_report_id_seq OWNER TO postgres;

--
-- Name: andon_report_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.andon_report_id_seq OWNED BY public.andon_report.id;


--
-- Name: aps_capacity_analysis; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aps_capacity_analysis (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    workshop_id bigint,
    workshop_name character varying(100),
    line_id bigint,
    line_name character varying(100),
    work_date character varying(10),
    shift_id bigint,
    shift_name character varying(50),
    plan_capacity numeric(18,2),
    actual_capacity numeric(18,2),
    utilization numeric(5,2),
    output_qty numeric(18,2),
    target_qty numeric(18,2),
    reject_qty numeric(18,2),
    reject_rate numeric(5,2),
    uptime numeric(5,2),
    downtime numeric(5,2),
    status bigint DEFAULT 1
);


ALTER TABLE public.aps_capacity_analysis OWNER TO postgres;

--
-- Name: aps_capacity_analysis_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.aps_capacity_analysis_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.aps_capacity_analysis_id_seq OWNER TO postgres;

--
-- Name: aps_capacity_analysis_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.aps_capacity_analysis_id_seq OWNED BY public.aps_capacity_analysis.id;


--
-- Name: aps_changeover_matrix; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aps_changeover_matrix (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    from_product_id bigint,
    from_product_code character varying(50),
    from_product_name character varying(100),
    to_product_id bigint,
    to_product_code character varying(50),
    to_product_name character varying(100),
    changeover_time numeric(10,2),
    setup_time numeric(10,2),
    clean_time numeric(10,2),
    is_optimized bigint DEFAULT 0,
    priority bigint DEFAULT 1,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.aps_changeover_matrix OWNER TO postgres;

--
-- Name: aps_changeover_matrix_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.aps_changeover_matrix_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.aps_changeover_matrix_id_seq OWNER TO postgres;

--
-- Name: aps_changeover_matrix_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.aps_changeover_matrix_id_seq OWNED BY public.aps_changeover_matrix.id;


--
-- Name: aps_delivery_rate; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aps_delivery_rate (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    order_no character varying(50),
    customer_id bigint,
    customer_name character varying(100),
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    plan_delivery_date character varying(10),
    actual_delivery_date timestamp with time zone,
    plan_qty numeric(18,4),
    delivery_qty numeric(18,4),
    on_time_qty numeric(18,4),
    on_time_rate numeric(5,2),
    fulfillment_rate numeric(5,2),
    status bigint DEFAULT 1
);


ALTER TABLE public.aps_delivery_rate OWNER TO postgres;

--
-- Name: aps_delivery_rate_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.aps_delivery_rate_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.aps_delivery_rate_id_seq OWNER TO postgres;

--
-- Name: aps_delivery_rate_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.aps_delivery_rate_id_seq OWNED BY public.aps_delivery_rate.id;


--
-- Name: aps_jit_demand; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aps_jit_demand (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    demand_no character varying(50),
    demand_type character varying(20),
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    customer_id bigint,
    customer_name character varying(100),
    demand_qty numeric(18,4),
    demand_time timestamp with time zone,
    priority bigint DEFAULT 1,
    frequency character varying(20),
    lead_time bigint,
    kanban_qty numeric(18,4),
    status bigint DEFAULT 1
);


ALTER TABLE public.aps_jit_demand OWNER TO postgres;

--
-- Name: aps_jit_demand_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.aps_jit_demand_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.aps_jit_demand_id_seq OWNER TO postgres;

--
-- Name: aps_jit_demand_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.aps_jit_demand_id_seq OWNED BY public.aps_jit_demand.id;


--
-- Name: aps_mps; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aps_mps (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    mps_no character varying(50) NOT NULL,
    plan_month character varying(10),
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    quantity numeric(18,4),
    status bigint DEFAULT 1
);


ALTER TABLE public.aps_mps OWNER TO postgres;

--
-- Name: aps_mps_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.aps_mps_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.aps_mps_id_seq OWNER TO postgres;

--
-- Name: aps_mps_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.aps_mps_id_seq OWNED BY public.aps_mps.id;


--
-- Name: aps_mrp; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aps_mrp (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    mrp_no character varying(50) NOT NULL,
    mrp_type character varying(20),
    plan_date timestamp(6) without time zone,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.aps_mrp OWNER TO postgres;

--
-- Name: aps_mrp_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.aps_mrp_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.aps_mrp_id_seq OWNER TO postgres;

--
-- Name: aps_mrp_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.aps_mrp_id_seq OWNED BY public.aps_mrp.id;


--
-- Name: aps_mrp_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aps_mrp_item (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    mrp_id bigint,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    quantity numeric(18,4),
    stock_qty numeric(18,4),
    allocated_qty numeric(18,4),
    net_qty numeric(18,4),
    source_type character varying(20),
    source_no character varying(50)
);


ALTER TABLE public.aps_mrp_item OWNER TO postgres;

--
-- Name: aps_mrp_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.aps_mrp_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.aps_mrp_item_id_seq OWNER TO postgres;

--
-- Name: aps_mrp_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.aps_mrp_item_id_seq OWNED BY public.aps_mrp_item.id;


--
-- Name: aps_resource; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aps_resource (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    resource_code character varying(50) NOT NULL,
    resource_name character varying(100) NOT NULL,
    resource_type character varying(20),
    workshop_id bigint,
    capacity numeric(18,2),
    unit character varying(20),
    efficiency numeric(10,2) DEFAULT 100,
    status bigint DEFAULT 1
);


ALTER TABLE public.aps_resource OWNER TO postgres;

--
-- Name: aps_resource_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.aps_resource_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.aps_resource_id_seq OWNER TO postgres;

--
-- Name: aps_resource_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.aps_resource_id_seq OWNED BY public.aps_resource.id;


--
-- Name: aps_rolling_schedule; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aps_rolling_schedule (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    schedule_no character varying(50),
    schedule_type character varying(20),
    start_date character varying(10),
    end_date character varying(10),
    workshop_id bigint,
    workshop_name character varying(100),
    line_id bigint,
    line_name character varying(100),
    plan_qty numeric(18,4),
    completed_qty numeric(18,4),
    start_time timestamp with time zone,
    end_time timestamp with time zone,
    horizon_days bigint,
    status bigint DEFAULT 1,
    execute_time timestamp with time zone
);


ALTER TABLE public.aps_rolling_schedule OWNER TO postgres;

--
-- Name: aps_rolling_schedule_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.aps_rolling_schedule_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.aps_rolling_schedule_id_seq OWNER TO postgres;

--
-- Name: aps_rolling_schedule_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.aps_rolling_schedule_id_seq OWNED BY public.aps_rolling_schedule.id;


--
-- Name: aps_schedule_plan; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aps_schedule_plan (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    plan_no character varying(50) NOT NULL,
    plan_type character varying(20),
    start_date timestamp(6) without time zone,
    end_date timestamp(6) without time zone,
    algorithm character varying(20),
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.aps_schedule_plan OWNER TO postgres;

--
-- Name: aps_schedule_plan_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.aps_schedule_plan_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.aps_schedule_plan_id_seq OWNER TO postgres;

--
-- Name: aps_schedule_plan_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.aps_schedule_plan_id_seq OWNED BY public.aps_schedule_plan.id;


--
-- Name: aps_schedule_result; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aps_schedule_result (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    plan_id bigint,
    order_id bigint,
    order_no character varying(50),
    sequence bigint,
    line_id bigint,
    line_name character varying(100),
    station_id bigint,
    station_name character varying(100),
    plan_start_time timestamp(6) without time zone,
    plan_end_time timestamp(6) without time zone,
    actual_start_time timestamp(6) without time zone,
    actual_end_time timestamp(6) without time zone,
    status bigint DEFAULT 1
);


ALTER TABLE public.aps_schedule_result OWNER TO postgres;

--
-- Name: aps_schedule_result_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.aps_schedule_result_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.aps_schedule_result_id_seq OWNER TO postgres;

--
-- Name: aps_schedule_result_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.aps_schedule_result_id_seq OWNED BY public.aps_schedule_result.id;


--
-- Name: aps_work_center; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aps_work_center (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    work_center_code character varying(50) NOT NULL,
    work_center_name character varying(100) NOT NULL,
    workshop_id bigint,
    capacity numeric(18,2),
    status character varying(20) DEFAULT 'ACTIVE'::character varying,
    work_center_type character varying(20) DEFAULT 'LINE'::character varying,
    capacity_unit character varying(20) DEFAULT 'HOUR'::character varying,
    standard_capacity numeric(18,3) DEFAULT 1,
    max_capacity numeric(18,3) DEFAULT 1,
    efficiency_factor numeric(5,2) DEFAULT 100,
    utilization_target numeric(5,2) DEFAULT 85,
    setup_time bigint DEFAULT 0,
    description character varying(500)
);


ALTER TABLE public.aps_work_center OWNER TO postgres;

--
-- Name: aps_work_center_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.aps_work_center_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.aps_work_center_id_seq OWNER TO postgres;

--
-- Name: aps_work_center_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.aps_work_center_id_seq OWNED BY public.aps_work_center.id;


--
-- Name: aps_working_calendar; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.aps_working_calendar (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    workshop_id bigint NOT NULL,
    calendar_name character varying(100),
    work_days jsonb,
    shifts jsonb,
    holiday_dates jsonb,
    special_work_dates jsonb,
    effective_from timestamp with time zone,
    effective_to timestamp with time zone,
    status bigint DEFAULT 1
);


ALTER TABLE public.aps_working_calendar OWNER TO postgres;

--
-- Name: aps_working_calendar_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.aps_working_calendar_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.aps_working_calendar_id_seq OWNER TO postgres;

--
-- Name: aps_working_calendar_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.aps_working_calendar_id_seq OWNED BY public.aps_working_calendar.id;


--
-- Name: bom_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bom_items (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    bom_id bigint NOT NULL,
    material_id bigint NOT NULL,
    material_code character varying(50),
    material_name character varying(100),
    quantity numeric(18,4),
    unit character varying(20),
    scrap_rate numeric DEFAULT 0,
    level bigint DEFAULT 1,
    parent_id bigint,
    sort bigint DEFAULT 0
);


ALTER TABLE public.bom_items OWNER TO postgres;

--
-- Name: bom_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bom_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bom_items_id_seq OWNER TO postgres;

--
-- Name: bom_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bom_items_id_seq OWNED BY public.bom_items.id;


--
-- Name: boms; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.boms (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    bom_code character varying(50) NOT NULL,
    product_id bigint NOT NULL,
    product_code character varying(50),
    product_name character varying(100),
    version character varying(20),
    bom_type character varying(20),
    status bigint DEFAULT 1,
    eff_date timestamp(6) with time zone,
    exp_date timestamp(6) with time zone
);


ALTER TABLE public.boms OWNER TO postgres;

--
-- Name: boms_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.boms_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.boms_id_seq OWNER TO postgres;

--
-- Name: boms_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.boms_id_seq OWNED BY public.boms.id;


--
-- Name: bpm_approval_record; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bpm_approval_record (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    task_id bigint NOT NULL,
    instance_id bigint NOT NULL,
    approver_id bigint NOT NULL,
    approver_name character varying(50) NOT NULL,
    approver_dept character varying(100),
    action character varying(20) NOT NULL,
    comment text,
    action_time timestamp with time zone NOT NULL,
    assignee_id bigint,
    assignee_name character varying(50),
    is_rollback bigint DEFAULT 0,
    rollback_target character varying(50),
    tenant_id bigint NOT NULL
);


ALTER TABLE public.bpm_approval_record OWNER TO postgres;

--
-- Name: bpm_approval_record_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bpm_approval_record_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bpm_approval_record_id_seq OWNER TO postgres;

--
-- Name: bpm_approval_record_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bpm_approval_record_id_seq OWNED BY public.bpm_approval_record.id;


--
-- Name: bpm_biz_mapping; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bpm_biz_mapping (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    biz_type character varying(50) NOT NULL,
    model_id bigint NOT NULL,
    trigger_event character varying(50) NOT NULL,
    is_active bigint DEFAULT 1,
    tenant_id bigint NOT NULL
);


ALTER TABLE public.bpm_biz_mapping OWNER TO postgres;

--
-- Name: bpm_biz_mapping_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bpm_biz_mapping_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bpm_biz_mapping_id_seq OWNER TO postgres;

--
-- Name: bpm_biz_mapping_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bpm_biz_mapping_id_seq OWNED BY public.bpm_biz_mapping.id;


--
-- Name: bpm_delegate_record; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bpm_delegate_record (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    delegate_id bigint NOT NULL,
    delegate_name character varying(50),
    delegatee_id bigint NOT NULL,
    delegatee_name character varying(50),
    start_date date,
    end_date date,
    biz_types jsonb,
    is_active bigint DEFAULT 1,
    tenant_id bigint NOT NULL
);


ALTER TABLE public.bpm_delegate_record OWNER TO postgres;

--
-- Name: bpm_delegate_record_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bpm_delegate_record_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bpm_delegate_record_id_seq OWNER TO postgres;

--
-- Name: bpm_delegate_record_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bpm_delegate_record_id_seq OWNED BY public.bpm_delegate_record.id;


--
-- Name: bpm_form_definition; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bpm_form_definition (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    form_code character varying(50) NOT NULL,
    form_name character varying(200) NOT NULL,
    form_type character varying(20) NOT NULL,
    version character varying(20) DEFAULT '1.0'::character varying,
    category character varying(50),
    is_published bigint DEFAULT 0,
    published_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    created_by character varying(50),
    updated_at timestamp with time zone
);


ALTER TABLE public.bpm_form_definition OWNER TO postgres;

--
-- Name: bpm_form_definition_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bpm_form_definition_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bpm_form_definition_id_seq OWNER TO postgres;

--
-- Name: bpm_form_definition_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bpm_form_definition_id_seq OWNED BY public.bpm_form_definition.id;


--
-- Name: bpm_form_field; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bpm_form_field (
    id bigint NOT NULL,
    form_id bigint NOT NULL,
    field_code character varying(50) NOT NULL,
    field_name character varying(100) NOT NULL,
    field_type character varying(30) NOT NULL,
    field_config jsonb NOT NULL,
    default_value text,
    validation_rules jsonb,
    is_required bigint DEFAULT 0,
    is_readonly bigint DEFAULT 0,
    is_hidden bigint DEFAULT 0,
    sort_order bigint DEFAULT 0,
    tenant_id bigint NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.bpm_form_field OWNER TO postgres;

--
-- Name: bpm_form_field_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bpm_form_field_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bpm_form_field_id_seq OWNER TO postgres;

--
-- Name: bpm_form_field_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bpm_form_field_id_seq OWNED BY public.bpm_form_field.id;


--
-- Name: bpm_form_version; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bpm_form_version (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    form_id bigint NOT NULL,
    version character varying(20) NOT NULL,
    change_description text,
    fields_snapshot jsonb,
    published_by bigint,
    published_at timestamp with time zone,
    is_current bigint DEFAULT 0,
    tenant_id bigint NOT NULL
);


ALTER TABLE public.bpm_form_version OWNER TO postgres;

--
-- Name: bpm_form_version_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bpm_form_version_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bpm_form_version_id_seq OWNER TO postgres;

--
-- Name: bpm_form_version_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bpm_form_version_id_seq OWNED BY public.bpm_form_version.id;


--
-- Name: bpm_model_version; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bpm_model_version (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    model_id bigint NOT NULL,
    version character varying(20) NOT NULL,
    change_description text,
    published_by bigint,
    published_at timestamp with time zone,
    is_current bigint DEFAULT 0,
    tenant_id bigint NOT NULL
);


ALTER TABLE public.bpm_model_version OWNER TO postgres;

--
-- Name: bpm_model_version_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bpm_model_version_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bpm_model_version_id_seq OWNER TO postgres;

--
-- Name: bpm_model_version_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bpm_model_version_id_seq OWNED BY public.bpm_model_version.id;


--
-- Name: bpm_node_definition; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bpm_node_definition (
    id bigint NOT NULL,
    model_id bigint NOT NULL,
    node_id character varying(50) NOT NULL,
    node_name character varying(200) NOT NULL,
    node_type character varying(30) NOT NULL,
    position_x bigint DEFAULT 0,
    position_y bigint DEFAULT 0,
    width bigint DEFAULT 120,
    height bigint DEFAULT 80,
    node_config jsonb NOT NULL,
    sort_order bigint DEFAULT 0,
    tenant_id bigint NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.bpm_node_definition OWNER TO postgres;

--
-- Name: bpm_node_definition_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bpm_node_definition_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bpm_node_definition_id_seq OWNER TO postgres;

--
-- Name: bpm_node_definition_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bpm_node_definition_id_seq OWNED BY public.bpm_node_definition.id;


--
-- Name: bpm_node_timeout; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bpm_node_timeout (
    id bigint NOT NULL,
    model_id bigint NOT NULL,
    node_id character varying(50) NOT NULL,
    timeout_minutes bigint DEFAULT 0,
    timeout_action character varying(20),
    timeout_notice_roles jsonb,
    escalation_node_id character varying(50),
    tenant_id bigint NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.bpm_node_timeout OWNER TO postgres;

--
-- Name: bpm_node_timeout_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bpm_node_timeout_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bpm_node_timeout_id_seq OWNER TO postgres;

--
-- Name: bpm_node_timeout_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bpm_node_timeout_id_seq OWNED BY public.bpm_node_timeout.id;


--
-- Name: bpm_order_day_workorder_map; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bpm_order_day_workorder_map (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    day_plan_id bigint NOT NULL,
    day_plan_item_id bigint NOT NULL,
    work_order_id bigint NOT NULL,
    work_order_no character varying(50) NOT NULL,
    tenant_id bigint NOT NULL
);


ALTER TABLE public.bpm_order_day_workorder_map OWNER TO postgres;

--
-- Name: bpm_order_day_workorder_map_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bpm_order_day_workorder_map_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bpm_order_day_workorder_map_id_seq OWNER TO postgres;

--
-- Name: bpm_order_day_workorder_map_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bpm_order_day_workorder_map_id_seq OWNED BY public.bpm_order_day_workorder_map.id;


--
-- Name: bpm_process_instance; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bpm_process_instance (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    instance_no character varying(50) NOT NULL,
    model_id bigint NOT NULL,
    model_code character varying(50),
    model_name character varying(200),
    version character varying(20),
    biz_type character varying(50),
    biz_id bigint,
    biz_no character varying(100),
    title character varying(500) NOT NULL,
    initiator_id bigint NOT NULL,
    initiator_name character varying(50),
    current_node_id character varying(50),
    current_node_name character varying(200),
    status character varying(20) NOT NULL,
    priority bigint DEFAULT 5,
    due_date timestamp with time zone,
    completed_at timestamp with time zone,
    cancel_reason text,
    form_data jsonb,
    business_data jsonb,
    tenant_id bigint NOT NULL,
    updated_at timestamp with time zone
);


ALTER TABLE public.bpm_process_instance OWNER TO postgres;

--
-- Name: bpm_process_instance_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bpm_process_instance_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bpm_process_instance_id_seq OWNER TO postgres;

--
-- Name: bpm_process_instance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bpm_process_instance_id_seq OWNED BY public.bpm_process_instance.id;


--
-- Name: bpm_process_model; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bpm_process_model (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    model_code character varying(50) NOT NULL,
    model_name character varying(200) NOT NULL,
    model_type character varying(20) NOT NULL,
    version character varying(20) DEFAULT '1.0'::character varying,
    category character varying(50),
    description text,
    form_type character varying(20) NOT NULL,
    form_definition_id bigint,
    form_url character varying(500),
    is_published bigint DEFAULT 0,
    is_active bigint DEFAULT 1,
    published_at timestamp with time zone,
    published_by bigint,
    config jsonb,
    tenant_id bigint NOT NULL,
    created_by character varying(50),
    updated_at timestamp with time zone
);


ALTER TABLE public.bpm_process_model OWNER TO postgres;

--
-- Name: bpm_process_model_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bpm_process_model_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bpm_process_model_id_seq OWNER TO postgres;

--
-- Name: bpm_process_model_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bpm_process_model_id_seq OWNED BY public.bpm_process_model.id;


--
-- Name: bpm_reminder_record; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bpm_reminder_record (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    instance_id bigint NOT NULL,
    task_id bigint,
    remind_type character varying(20) NOT NULL,
    remind_from character varying(50),
    remind_to character varying(50),
    remind_time timestamp with time zone NOT NULL,
    remind_count bigint DEFAULT 1,
    tenant_id bigint NOT NULL
);


ALTER TABLE public.bpm_reminder_record OWNER TO postgres;

--
-- Name: bpm_reminder_record_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bpm_reminder_record_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bpm_reminder_record_id_seq OWNER TO postgres;

--
-- Name: bpm_reminder_record_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bpm_reminder_record_id_seq OWNED BY public.bpm_reminder_record.id;


--
-- Name: bpm_sequence_flow; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bpm_sequence_flow (
    id bigint NOT NULL,
    model_id bigint NOT NULL,
    flow_id character varying(50) NOT NULL,
    source_node_id character varying(50) NOT NULL,
    target_node_id character varying(50) NOT NULL,
    flow_name character varying(200),
    condition_type character varying(20) DEFAULT 'NONE'::character varying,
    condition_expression text,
    is_default bigint DEFAULT 0,
    flow_config jsonb,
    sort_order bigint DEFAULT 0,
    tenant_id bigint NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.bpm_sequence_flow OWNER TO postgres;

--
-- Name: bpm_sequence_flow_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bpm_sequence_flow_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bpm_sequence_flow_id_seq OWNER TO postgres;

--
-- Name: bpm_sequence_flow_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bpm_sequence_flow_id_seq OWNED BY public.bpm_sequence_flow.id;


--
-- Name: bpm_task_assignment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bpm_task_assignment (
    id bigint NOT NULL,
    model_id bigint NOT NULL,
    node_id character varying(50) NOT NULL,
    rule_type character varying(20) NOT NULL,
    rule_value text NOT NULL,
    priority bigint DEFAULT 5,
    allow_assign bigint DEFAULT 1,
    allow_transfer bigint DEFAULT 1,
    tenant_id bigint NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.bpm_task_assignment OWNER TO postgres;

--
-- Name: bpm_task_assignment_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bpm_task_assignment_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bpm_task_assignment_id_seq OWNER TO postgres;

--
-- Name: bpm_task_assignment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bpm_task_assignment_id_seq OWNED BY public.bpm_task_assignment.id;


--
-- Name: bpm_task_instance; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.bpm_task_instance (
    id bigint NOT NULL,
    task_no character varying(50) NOT NULL,
    instance_id bigint NOT NULL,
    node_id character varying(50) NOT NULL,
    node_name character varying(200) NOT NULL,
    node_type character varying(30) NOT NULL,
    task_type character varying(20) DEFAULT 'APPROVAL'::character varying,
    assign_type character varying(20) NOT NULL,
    assign_value character varying(200),
    assignee_id bigint,
    assignee_name character varying(50),
    assignee_list jsonb,
    action_result character varying(20),
    action_comment text,
    action_time timestamp with time zone,
    sign_type character varying(20),
    required_approvers bigint DEFAULT 1,
    current_approvers bigint DEFAULT 0,
    status character varying(20) NOT NULL,
    is_current bigint DEFAULT 0,
    created_at timestamp with time zone,
    completed_at timestamp with time zone
);


ALTER TABLE public.bpm_task_instance OWNER TO postgres;

--
-- Name: bpm_task_instance_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.bpm_task_instance_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.bpm_task_instance_id_seq OWNER TO postgres;

--
-- Name: bpm_task_instance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.bpm_task_instance_id_seq OWNED BY public.bpm_task_instance.id;


--
-- Name: container_master; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.container_master (
    id bigint NOT NULL,
    tenant_id bigint NOT NULL,
    container_code character varying(50) NOT NULL,
    container_name character varying(100) NOT NULL,
    container_type character varying(30) NOT NULL,
    standard_qty bigint DEFAULT 0,
    applicable_products jsonb,
    status character varying(20) DEFAULT 'IN_STOCK'::character varying,
    location_type character varying(20),
    current_location character varying(100),
    customer_id bigint,
    barcode character varying(100),
    total_trips bigint DEFAULT 0,
    last_clean_date timestamp with time zone,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.container_master OWNER TO postgres;

--
-- Name: container_master_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.container_master_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.container_master_id_seq OWNER TO postgres;

--
-- Name: container_master_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.container_master_id_seq OWNED BY public.container_master.id;


--
-- Name: container_movement; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.container_movement (
    id bigint NOT NULL,
    tenant_id bigint NOT NULL,
    container_id bigint NOT NULL,
    container_code character varying(50),
    movement_type character varying(20) NOT NULL,
    from_location character varying(100),
    to_location character varying(100),
    qty bigint DEFAULT 1,
    related_order_no character varying(50),
    operator_id bigint,
    operator_name character varying(50),
    movement_time timestamp with time zone NOT NULL,
    remark character varying(200),
    created_at timestamp with time zone
);


ALTER TABLE public.container_movement OWNER TO postgres;

--
-- Name: container_movement_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.container_movement_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.container_movement_id_seq OWNER TO postgres;

--
-- Name: container_movement_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.container_movement_id_seq OWNED BY public.container_movement.id;


--
-- Name: data_collections; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.data_collections (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    equipment_id bigint,
    equipment_code character varying(50),
    station_id bigint,
    data_type character varying(20),
    data_key character varying(50),
    data_value character varying(200),
    unit character varying(20),
    collect_time timestamp(6) with time zone
);


ALTER TABLE public.data_collections OWNER TO postgres;

--
-- Name: data_collections_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.data_collections_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.data_collections_id_seq OWNER TO postgres;

--
-- Name: data_collections_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.data_collections_id_seq OWNED BY public.data_collections.id;


--
-- Name: dc_collect_record; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dc_collect_record (
    id bigint NOT NULL,
    tenant_id bigint NOT NULL,
    point_id bigint NOT NULL,
    device_id bigint NOT NULL,
    value_raw character varying(200),
    value_numeric numeric(18,6),
    value_text character varying(500),
    quality text,
    collect_time timestamp with time zone NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.dc_collect_record OWNER TO postgres;

--
-- Name: dc_collect_record_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.dc_collect_record_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.dc_collect_record_id_seq OWNER TO postgres;

--
-- Name: dc_collect_record_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.dc_collect_record_id_seq OWNED BY public.dc_collect_record.id;


--
-- Name: dc_data_point; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dc_data_point (
    id bigint NOT NULL,
    tenant_id bigint NOT NULL,
    point_code character varying(50) NOT NULL,
    point_name character varying(100) NOT NULL,
    device_id bigint NOT NULL,
    data_type character varying(20),
    protocol character varying(20) NOT NULL,
    address character varying(200),
    unit character varying(20),
    scan_rate bigint DEFAULT 1000,
    deadband numeric(10,4) DEFAULT 0,
    store_policy character varying(20) DEFAULT 'ALL'::character varying,
    alarm_enabled bigint DEFAULT 0,
    alarm_high numeric(18,4),
    alarm_low numeric(18,4),
    map_to_field character varying(100),
    status character varying(20) DEFAULT 'ACTIVE'::character varying,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.dc_data_point OWNER TO postgres;

--
-- Name: dc_data_point_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.dc_data_point_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.dc_data_point_id_seq OWNER TO postgres;

--
-- Name: dc_data_point_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.dc_data_point_id_seq OWNED BY public.dc_data_point.id;


--
-- Name: dc_scan_log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dc_scan_log (
    id bigint NOT NULL,
    tenant_id bigint NOT NULL,
    workshop_id bigint,
    scan_type character varying(20) NOT NULL,
    scan_code character varying(200) NOT NULL,
    scan_device character varying(50),
    workstation_id bigint,
    scan_user_id bigint,
    scan_time timestamp with time zone,
    parse_result jsonb,
    business_type character varying(50),
    related_id bigint,
    status character varying(20) DEFAULT 'SUCCESS'::character varying,
    fail_reason character varying(200),
    created_at timestamp with time zone
);


ALTER TABLE public.dc_scan_log OWNER TO postgres;

--
-- Name: dc_scan_log_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.dc_scan_log_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.dc_scan_log_id_seq OWNER TO postgres;

--
-- Name: dc_scan_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.dc_scan_log_id_seq OWNED BY public.dc_scan_log.id;


--
-- Name: defect_codes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.defect_codes (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    defect_code character varying(20) NOT NULL,
    defect_name character varying(100) NOT NULL,
    defect_type character varying(20),
    severity bigint DEFAULT 1,
    status bigint DEFAULT 1
);


ALTER TABLE public.defect_codes OWNER TO postgres;

--
-- Name: defect_codes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.defect_codes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.defect_codes_id_seq OWNER TO postgres;

--
-- Name: defect_codes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.defect_codes_id_seq OWNED BY public.defect_codes.id;


--
-- Name: defect_records; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.defect_records (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    record_no character varying(50) NOT NULL,
    order_id bigint,
    order_no character varying(50),
    process_id bigint,
    process_name character varying(100),
    defect_code_id bigint,
    defect_code character varying(20),
    defect_name character varying(100),
    quantity numeric(18,4),
    handle_method bigint DEFAULT 1,
    handle_user_id bigint,
    handle_date timestamp(6) with time zone,
    status bigint DEFAULT 1
);


ALTER TABLE public.defect_records OWNER TO postgres;

--
-- Name: defect_records_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.defect_records_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.defect_records_id_seq OWNER TO postgres;

--
-- Name: defect_records_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.defect_records_id_seq OWNED BY public.defect_records.id;


--
-- Name: delivery_order_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.delivery_order_items (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    delivery_id bigint,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    quantity numeric(18,4),
    shipped_qty numeric(18,4) DEFAULT 0,
    unit character varying(20),
    batch_no character varying(50)
);


ALTER TABLE public.delivery_order_items OWNER TO postgres;

--
-- Name: delivery_order_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.delivery_order_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.delivery_order_items_id_seq OWNER TO postgres;

--
-- Name: delivery_order_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.delivery_order_items_id_seq OWNED BY public.delivery_order_items.id;


--
-- Name: delivery_orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.delivery_orders (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    delivery_no character varying(50) NOT NULL,
    customer_id bigint,
    customer_name character varying(100),
    warehouse_id bigint,
    delivery_date timestamp(6) with time zone,
    delivery_user_id bigint,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.delivery_orders OWNER TO postgres;

--
-- Name: delivery_orders_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.delivery_orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.delivery_orders_id_seq OWNER TO postgres;

--
-- Name: delivery_orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.delivery_orders_id_seq OWNED BY public.delivery_orders.id;


--
-- Name: delivery_report; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.delivery_report (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    report_month timestamp with time zone NOT NULL,
    customer_id bigint,
    customer_name character varying(100),
    order_count bigint,
    total_order_qty numeric,
    delivered_qty numeric,
    on_time_deliver_qty numeric,
    delivery_rate numeric,
    on_time_rate numeric,
    late_deliver_qty numeric,
    remark character varying(500)
);


ALTER TABLE public.delivery_report OWNER TO postgres;

--
-- Name: delivery_report_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.delivery_report_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.delivery_report_id_seq OWNER TO postgres;

--
-- Name: delivery_report_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.delivery_report_id_seq OWNED BY public.delivery_report.id;


--
-- Name: dispatches; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dispatches (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    order_id bigint NOT NULL,
    order_no character varying(50),
    process_id bigint,
    process_name character varying(100),
    station_id bigint,
    station_name character varying(100),
    assign_user_id bigint,
    assign_user_name character varying(50),
    quantity numeric(18,4),
    status bigint DEFAULT 1
);


ALTER TABLE public.dispatches OWNER TO postgres;

--
-- Name: dispatches_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.dispatches_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.dispatches_id_seq OWNER TO postgres;

--
-- Name: dispatches_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.dispatches_id_seq OWNED BY public.dispatches.id;


--
-- Name: eam_equipment_downtime; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.eam_equipment_downtime (
    id bigint NOT NULL,
    tenant_id bigint DEFAULT 1 NOT NULL,
    equipment_id bigint NOT NULL,
    equipment_code character varying(50),
    equipment_name character varying(100),
    downtime_type character varying(20),
    downtime_reason character varying(200),
    start_time timestamp with time zone,
    end_time timestamp with time zone,
    duration bigint,
    lost_production numeric,
    work_order_id bigint,
    work_order_code character varying(50),
    shift_id bigint,
    operator_id bigint,
    operator_name character varying(50),
    maintainer_id bigint,
    maintainer_name character varying(50),
    status character varying(20),
    remark character varying(500),
    created_by character varying(50),
    updated_by character varying(50),
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.eam_equipment_downtime OWNER TO postgres;

--
-- Name: eam_equipment_downtime_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.eam_equipment_downtime_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.eam_equipment_downtime_id_seq OWNER TO postgres;

--
-- Name: eam_equipment_downtime_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.eam_equipment_downtime_id_seq OWNED BY public.eam_equipment_downtime.id;


--
-- Name: eam_repair_flow; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.eam_repair_flow (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    flow_code character varying(50),
    flow_name character varying(100),
    flow_steps text,
    status character varying(20) DEFAULT 'ACTIVE'::character varying,
    created_by bigint,
    updated_by bigint
);


ALTER TABLE public.eam_repair_flow OWNER TO postgres;

--
-- Name: eam_repair_flow_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.eam_repair_flow_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.eam_repair_flow_id_seq OWNER TO postgres;

--
-- Name: eam_repair_flow_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.eam_repair_flow_id_seq OWNED BY public.eam_repair_flow.id;


--
-- Name: eam_repair_job; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.eam_repair_job (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    job_code character varying(50),
    equipment_id bigint,
    equipment_code character varying(50),
    fault_type character varying(30),
    fault_reason character varying(30),
    fault_desc character varying(500),
    level character varying(20),
    status character varying(20) DEFAULT 'PENDING'::character varying,
    reporter_id bigint,
    reporter_name character varying(100),
    assignee_id bigint,
    assignee_name character varying(100),
    plan_start_time timestamp with time zone,
    plan_end_time timestamp with time zone,
    actual_start_time timestamp with time zone,
    actual_end_time timestamp with time zone,
    result character varying(500),
    evaluation character varying(20),
    created_by bigint,
    updated_by bigint
);


ALTER TABLE public.eam_repair_job OWNER TO postgres;

--
-- Name: eam_repair_job_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.eam_repair_job_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.eam_repair_job_id_seq OWNER TO postgres;

--
-- Name: eam_repair_job_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.eam_repair_job_id_seq OWNED BY public.eam_repair_job.id;


--
-- Name: eam_repair_std; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.eam_repair_std (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    std_code character varying(50),
    std_name character varying(100),
    fault_type character varying(30),
    repair_steps text,
    tools_required character varying(500),
    materials_required character varying(500),
    standard_hours numeric,
    status character varying(20) DEFAULT 'ACTIVE'::character varying,
    created_by bigint,
    updated_by bigint
);


ALTER TABLE public.eam_repair_std OWNER TO postgres;

--
-- Name: eam_repair_std_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.eam_repair_std_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.eam_repair_std_id_seq OWNER TO postgres;

--
-- Name: eam_repair_std_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.eam_repair_std_id_seq OWNED BY public.eam_repair_std.id;


--
-- Name: ene_energy_record; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ene_energy_record (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    energy_type character varying(20),
    workshop_id bigint,
    workshop_name character varying(100),
    equipment_id bigint,
    equipment_name character varying(100),
    meter_no character varying(50),
    quantity numeric(18,2),
    unit character varying(20),
    unit_price numeric(18,4),
    amount numeric(18,2),
    record_date timestamp(6) without time zone NOT NULL,
    remark character varying(500)
);


ALTER TABLE public.ene_energy_record OWNER TO postgres;

--
-- Name: ene_energy_record_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ene_energy_record_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.ene_energy_record_id_seq OWNER TO postgres;

--
-- Name: ene_energy_record_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ene_energy_record_id_seq OWNED BY public.ene_energy_record.id;


--
-- Name: energy_records; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.energy_records (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    energy_type character varying(20),
    workshop_id bigint,
    workshop_name character varying(100),
    equipment_id bigint,
    equipment_name character varying(100),
    meter_no character varying(50),
    quantity numeric(18,2),
    unit character varying(20),
    unit_price numeric,
    amount numeric,
    record_date timestamp(6) with time zone,
    remark character varying(500)
);


ALTER TABLE public.energy_records OWNER TO postgres;

--
-- Name: energy_records_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.energy_records_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.energy_records_id_seq OWNER TO postgres;

--
-- Name: energy_records_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.energy_records_id_seq OWNED BY public.energy_records.id;


--
-- Name: equ_equipment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.equ_equipment (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    equipment_code character varying(50) NOT NULL,
    equipment_name character varying(100) NOT NULL,
    equipment_type character varying(20),
    brand character varying(50),
    model character varying(50),
    serial_number character varying(100),
    workshop_id bigint,
    workshop_name character varying(100),
    line_id bigint,
    line_name character varying(100),
    station_id bigint,
    station_name character varying(100),
    supplier character varying(100),
    purchase_date timestamp(6) without time zone,
    purchase_price numeric(18,2),
    warranty_end_date timestamp(6) without time zone,
    status integer DEFAULT 1
);


ALTER TABLE public.equ_equipment OWNER TO postgres;

--
-- Name: equ_equipment_check; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.equ_equipment_check (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    equipment_id bigint,
    equipment_code character varying(50),
    equipment_name character varying(100),
    check_plan_id bigint,
    check_user_id bigint,
    check_user_name character varying(50),
    check_date timestamp(6) without time zone,
    check_result integer,
    status integer DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.equ_equipment_check OWNER TO postgres;

--
-- Name: equ_equipment_check_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.equ_equipment_check_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.equ_equipment_check_id_seq OWNER TO postgres;

--
-- Name: equ_equipment_check_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.equ_equipment_check_id_seq OWNED BY public.equ_equipment_check.id;


--
-- Name: equ_equipment_document; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.equ_equipment_document (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    equipment_id bigint NOT NULL,
    equipment_code character varying(50),
    equipment_name character varying(100),
    doc_type character varying(30),
    doc_name character varying(200) NOT NULL,
    doc_code character varying(50),
    file_name character varying(200),
    file_path character varying(500),
    file_size bigint,
    file_type character varying(50),
    file_url character varying(500),
    version character varying(20),
    effective_date timestamp with time zone,
    expiry_date timestamp with time zone,
    description character varying(500),
    uploaded_by bigint,
    uploaded_by_name character varying(50),
    upload_time timestamp with time zone,
    status bigint DEFAULT 1
);


ALTER TABLE public.equ_equipment_document OWNER TO postgres;

--
-- Name: equ_equipment_document_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.equ_equipment_document_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.equ_equipment_document_id_seq OWNER TO postgres;

--
-- Name: equ_equipment_document_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.equ_equipment_document_id_seq OWNED BY public.equ_equipment_document.id;


--
-- Name: equ_equipment_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.equ_equipment_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.equ_equipment_id_seq OWNER TO postgres;

--
-- Name: equ_equipment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.equ_equipment_id_seq OWNED BY public.equ_equipment.id;


--
-- Name: equ_equipment_maintenance; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.equ_equipment_maintenance (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    equipment_id bigint,
    equipment_code character varying(50),
    equipment_name character varying(100),
    maint_type character varying(20),
    maint_plan_id bigint,
    maint_user_id bigint,
    maint_user_name character varying(50),
    maint_date timestamp(6) without time zone,
    start_time timestamp(6) without time zone,
    end_time timestamp(6) without time zone,
    duration integer,
    content text,
    cost numeric(18,2),
    status integer DEFAULT 1
);


ALTER TABLE public.equ_equipment_maintenance OWNER TO postgres;

--
-- Name: equ_equipment_maintenance_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.equ_equipment_maintenance_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.equ_equipment_maintenance_id_seq OWNER TO postgres;

--
-- Name: equ_equipment_maintenance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.equ_equipment_maintenance_id_seq OWNED BY public.equ_equipment_maintenance.id;


--
-- Name: equ_equipment_part; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.equ_equipment_part (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    equipment_id bigint NOT NULL,
    equipment_code character varying(50),
    equipment_name character varying(100),
    part_code character varying(50) NOT NULL,
    part_name character varying(100) NOT NULL,
    spec character varying(100),
    unit character varying(20),
    qty numeric(18,4),
    supplier character varying(100),
    unit_price numeric(18,2),
    total_price numeric(18,2),
    replacement_freq bigint,
    max_stock numeric(18,4),
    min_stock numeric(18,4),
    current_stock numeric(18,4),
    status bigint DEFAULT 1
);


ALTER TABLE public.equ_equipment_part OWNER TO postgres;

--
-- Name: equ_equipment_part_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.equ_equipment_part_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.equ_equipment_part_id_seq OWNER TO postgres;

--
-- Name: equ_equipment_part_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.equ_equipment_part_id_seq OWNED BY public.equ_equipment_part.id;


--
-- Name: equ_equipment_repair; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.equ_equipment_repair (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    equipment_id bigint,
    equipment_code character varying(50),
    equipment_name character varying(100),
    fault_desc text,
    fault_time timestamp(6) without time zone,
    report_user_id bigint,
    report_user_name character varying(50),
    repair_user_id bigint,
    repair_user_name character varying(50),
    start_time timestamp(6) without time zone,
    end_time timestamp(6) without time zone,
    duration integer,
    repair_content text,
    cost numeric(18,2),
    status integer DEFAULT 1
);


ALTER TABLE public.equ_equipment_repair OWNER TO postgres;

--
-- Name: equ_equipment_repair_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.equ_equipment_repair_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.equ_equipment_repair_id_seq OWNER TO postgres;

--
-- Name: equ_equipment_repair_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.equ_equipment_repair_id_seq OWNED BY public.equ_equipment_repair.id;


--
-- Name: equ_gauge; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.equ_gauge (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    gauge_code character varying(50) NOT NULL,
    gauge_name character varying(100) NOT NULL,
    gauge_type character varying(20),
    spec character varying(100),
    "precision" character varying(50),
    measure_range character varying(100),
    cal_cycle bigint,
    last_cal_date timestamp with time zone,
    next_cal_date timestamp with time zone,
    workshop_id bigint,
    workshop_name character varying(100),
    location_id bigint,
    location_name character varying(100),
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.equ_gauge OWNER TO postgres;

--
-- Name: equ_gauge_calibration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.equ_gauge_calibration (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    gauge_id bigint,
    gauge_code character varying(50),
    gauge_name character varying(100),
    cal_date timestamp with time zone,
    cal_type character varying(20),
    cal_result bigint,
    standard character varying(100),
    actual_value character varying(100),
    error_value character varying(100),
    cal_user_id bigint,
    user_name character varying(50),
    agency character varying(100),
    certificate character varying(100),
    cost numeric(18,2),
    remark character varying(500)
);


ALTER TABLE public.equ_gauge_calibration OWNER TO postgres;

--
-- Name: equ_gauge_calibration_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.equ_gauge_calibration_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.equ_gauge_calibration_id_seq OWNER TO postgres;

--
-- Name: equ_gauge_calibration_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.equ_gauge_calibration_id_seq OWNED BY public.equ_gauge_calibration.id;


--
-- Name: equ_gauge_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.equ_gauge_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.equ_gauge_id_seq OWNER TO postgres;

--
-- Name: equ_gauge_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.equ_gauge_id_seq OWNED BY public.equ_gauge.id;


--
-- Name: equ_mold; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.equ_mold (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    mold_code character varying(50) NOT NULL,
    mold_name character varying(100) NOT NULL,
    mold_type character varying(50),
    product_code character varying(50),
    product_name character varying(100),
    cavity_count bigint,
    lifecycle bigint,
    used_count bigint,
    workshop_id bigint,
    workshop_name character varying(100),
    location_id bigint,
    location_name character varying(100),
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.equ_mold OWNER TO postgres;

--
-- Name: equ_mold_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.equ_mold_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.equ_mold_id_seq OWNER TO postgres;

--
-- Name: equ_mold_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.equ_mold_id_seq OWNED BY public.equ_mold.id;


--
-- Name: equ_mold_maintenance; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.equ_mold_maintenance (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    mold_id bigint,
    mold_code character varying(50),
    mold_name character varying(100),
    maint_type character varying(20),
    maint_date timestamp with time zone,
    maint_user_id bigint,
    user_name character varying(50),
    content text,
    duration bigint,
    cost numeric(18,2),
    status bigint DEFAULT 1
);


ALTER TABLE public.equ_mold_maintenance OWNER TO postgres;

--
-- Name: equ_mold_maintenance_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.equ_mold_maintenance_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.equ_mold_maintenance_id_seq OWNER TO postgres;

--
-- Name: equ_mold_maintenance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.equ_mold_maintenance_id_seq OWNED BY public.equ_mold_maintenance.id;


--
-- Name: equ_mold_repair; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.equ_mold_repair (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    mold_id bigint,
    mold_code character varying(50),
    mold_name character varying(100),
    repair_type character varying(20),
    repair_date timestamp with time zone,
    repair_user_id bigint,
    user_name character varying(50),
    reason character varying(500),
    content text,
    parts_used character varying(500),
    duration bigint,
    cost numeric(18,2),
    status bigint DEFAULT 1
);


ALTER TABLE public.equ_mold_repair OWNER TO postgres;

--
-- Name: equ_mold_repair_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.equ_mold_repair_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.equ_mold_repair_id_seq OWNER TO postgres;

--
-- Name: equ_mold_repair_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.equ_mold_repair_id_seq OWNED BY public.equ_mold_repair.id;


--
-- Name: equ_oee; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.equ_oee (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    equipment_id bigint,
    equipment_code character varying(50),
    equipment_name character varying(200),
    workshop_id bigint,
    record_date character varying(10),
    plan_time bigint DEFAULT 0,
    run_time bigint DEFAULT 0,
    down_time bigint DEFAULT 0,
    idle_time bigint DEFAULT 0,
    plan_stop_time bigint DEFAULT 0,
    output_qty bigint DEFAULT 0,
    qualified_qty bigint DEFAULT 0,
    availability numeric(5,2) DEFAULT 0,
    performance numeric(5,2) DEFAULT 0,
    quality numeric(5,2) DEFAULT 0,
    oee numeric(5,2) DEFAULT 0
);


ALTER TABLE public.equ_oee OWNER TO postgres;

--
-- Name: equ_oee_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.equ_oee_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.equ_oee_id_seq OWNER TO postgres;

--
-- Name: equ_oee_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.equ_oee_id_seq OWNED BY public.equ_oee.id;


--
-- Name: equ_oee_record; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.equ_oee_record (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    oee_id bigint,
    event_type character varying(20),
    start_time character varying(19),
    end_time character varying(19),
    duration bigint,
    reason character varying(500)
);


ALTER TABLE public.equ_oee_record OWNER TO postgres;

--
-- Name: equ_oee_record_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.equ_oee_record_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.equ_oee_record_id_seq OWNER TO postgres;

--
-- Name: equ_oee_record_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.equ_oee_record_id_seq OWNED BY public.equ_oee_record.id;


--
-- Name: equ_spare_part; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.equ_spare_part (
    id bigint NOT NULL,
    created_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp(6) without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    spare_part_code character varying(50) NOT NULL,
    spare_part_name character varying(100) NOT NULL,
    spec character varying(100),
    unit character varying(20),
    quantity numeric(18,2) DEFAULT 0,
    min_quantity numeric(18,2),
    price numeric(18,2),
    supplier character varying(100),
    status integer DEFAULT 1
);


ALTER TABLE public.equ_spare_part OWNER TO postgres;

--
-- Name: equ_spare_part_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.equ_spare_part_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.equ_spare_part_id_seq OWNER TO postgres;

--
-- Name: equ_spare_part_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.equ_spare_part_id_seq OWNED BY public.equ_spare_part.id;


--
-- Name: equ_teep_data; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.equ_teep_data (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    equipment_id bigint,
    equipment_code character varying(50),
    equipment_name character varying(100),
    report_date timestamp with time zone,
    plan_time numeric,
    down_time numeric,
    actual_time numeric,
    availability numeric,
    ideal_cycle_time numeric,
    actual_output bigint,
    ideal_output numeric,
    performance numeric,
    pass_output bigint,
    fail_output bigint,
    quality numeric,
    teep numeric,
    oee numeric,
    remark character varying(500)
);


ALTER TABLE public.equ_teep_data OWNER TO postgres;

--
-- Name: equ_teep_data_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.equ_teep_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.equ_teep_data_id_seq OWNER TO postgres;

--
-- Name: equ_teep_data_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.equ_teep_data_id_seq OWNED BY public.equ_teep_data.id;


--
-- Name: fin_payment_request; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.fin_payment_request (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    request_no character varying(50) NOT NULL,
    request_type character varying(20) NOT NULL,
    supplier_customer_id bigint,
    supplier_customer_name character varying(100),
    request_amount numeric(18,2) NOT NULL,
    amount_in_words character varying(200),
    purpose character varying(200),
    bank_name character varying(100),
    bank_account character varying(100),
    settlement_ids jsonb,
    attachment_urls jsonb,
    status character varying(20) DEFAULT 'PENDING'::character varying,
    approval_status character varying(20) DEFAULT 'PENDING'::character varying,
    approved_by bigint,
    approved_time timestamp with time zone,
    approver_comment text,
    paid_by bigint,
    paid_time timestamp with time zone,
    payment_status character varying(20) DEFAULT 'UNPAID'::character varying,
    tenant_id bigint NOT NULL,
    created_by character varying(50)
);


ALTER TABLE public.fin_payment_request OWNER TO postgres;

--
-- Name: fin_payment_request_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.fin_payment_request_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.fin_payment_request_id_seq OWNER TO postgres;

--
-- Name: fin_payment_request_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.fin_payment_request_id_seq OWNED BY public.fin_payment_request.id;


--
-- Name: fin_purchase_advance; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.fin_purchase_advance (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    advance_no character varying(50) NOT NULL,
    supplier_id bigint,
    supplier_name character varying(100),
    advance_amount numeric(18,2) NOT NULL,
    paid_date character varying(10) NOT NULL,
    payment_method character varying(20),
    payment_account character varying(50),
    bank_flow_no character varying(50),
    status character varying(20) DEFAULT 'PENDING'::character varying,
    used_amount numeric(18,2) DEFAULT 0,
    settlement_ids jsonb,
    remark text,
    tenant_id bigint NOT NULL,
    created_by character varying(50)
);


ALTER TABLE public.fin_purchase_advance OWNER TO postgres;

--
-- Name: fin_purchase_advance_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.fin_purchase_advance_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.fin_purchase_advance_id_seq OWNER TO postgres;

--
-- Name: fin_purchase_advance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.fin_purchase_advance_id_seq OWNED BY public.fin_purchase_advance.id;


--
-- Name: fin_purchase_settlement; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.fin_purchase_settlement (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    settlement_no character varying(50) NOT NULL,
    settlement_type character varying(20) NOT NULL,
    related_type character varying(20) NOT NULL,
    related_id bigint,
    related_no character varying(50),
    supplier_id bigint,
    supplier_code character varying(50),
    supplier_name character varying(100),
    invoice_no character varying(50),
    invoice_date character varying(10),
    goods_amount numeric(18,2) DEFAULT 0,
    tax_amount numeric(18,2) DEFAULT 0,
    total_amount numeric(18,2) DEFAULT 0,
    paid_amount numeric(18,2) DEFAULT 0,
    discount_amount numeric(18,2) DEFAULT 0,
    currency character varying(10) DEFAULT 'CNY'::character varying,
    exchange_rate numeric(10,4) DEFAULT 1,
    payment_terms character varying(50),
    payment_due_date character varying(10),
    payment_method character varying(20),
    status character varying(20) DEFAULT 'PENDING'::character varying,
    approved_by bigint,
    approved_time timestamp with time zone,
    settlement_date character varying(10),
    remark text,
    tenant_id bigint NOT NULL,
    created_by character varying(50)
);


ALTER TABLE public.fin_purchase_settlement OWNER TO postgres;

--
-- Name: fin_purchase_settlement_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.fin_purchase_settlement_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.fin_purchase_settlement_id_seq OWNER TO postgres;

--
-- Name: fin_purchase_settlement_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.fin_purchase_settlement_id_seq OWNED BY public.fin_purchase_settlement.id;


--
-- Name: fin_purchase_settlement_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.fin_purchase_settlement_item (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    settlement_id bigint NOT NULL,
    line_no bigint,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    specification character varying(200),
    unit character varying(20),
    invoice_qty numeric(18,3),
    received_qty numeric(18,3),
    settled_qty numeric(18,3),
    this_settle_qty numeric(18,3),
    unit_price numeric(18,4),
    tax_rate numeric(5,2),
    goods_amount numeric(18,2),
    tax_amount numeric(18,2),
    line_amount numeric(18,2),
    batch_no character varying(50),
    warehouse_id bigint,
    warehouse_name character varying(100),
    remark text,
    tenant_id bigint NOT NULL
);


ALTER TABLE public.fin_purchase_settlement_item OWNER TO postgres;

--
-- Name: fin_purchase_settlement_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.fin_purchase_settlement_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.fin_purchase_settlement_item_id_seq OWNER TO postgres;

--
-- Name: fin_purchase_settlement_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.fin_purchase_settlement_item_id_seq OWNED BY public.fin_purchase_settlement_item.id;


--
-- Name: fin_sales_receipt; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.fin_sales_receipt (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    receipt_no character varying(50) NOT NULL,
    customer_id bigint,
    customer_name character varying(100),
    receipt_amount numeric(18,2) NOT NULL,
    receipt_date character varying(10) NOT NULL,
    receipt_method character varying(20),
    receipt_account character varying(50),
    bank_flow_no character varying(50),
    settlement_ids jsonb,
    used_amount numeric(18,2) DEFAULT 0,
    status character varying(20) DEFAULT 'PENDING'::character varying,
    remark text,
    tenant_id bigint NOT NULL,
    created_by character varying(50)
);


ALTER TABLE public.fin_sales_receipt OWNER TO postgres;

--
-- Name: fin_sales_receipt_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.fin_sales_receipt_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.fin_sales_receipt_id_seq OWNER TO postgres;

--
-- Name: fin_sales_receipt_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.fin_sales_receipt_id_seq OWNED BY public.fin_sales_receipt.id;


--
-- Name: fin_sales_settlement; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.fin_sales_settlement (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    settlement_no character varying(50) NOT NULL,
    settlement_type character varying(20) NOT NULL,
    related_type character varying(20) NOT NULL,
    related_id bigint,
    related_no character varying(50),
    customer_id bigint,
    customer_code character varying(50),
    customer_name character varying(100),
    invoice_no character varying(50),
    invoice_date character varying(10),
    goods_amount numeric(18,2) DEFAULT 0,
    tax_amount numeric(18,2) DEFAULT 0,
    total_amount numeric(18,2) DEFAULT 0,
    received_amount numeric(18,2) DEFAULT 0,
    discount_amount numeric(18,2) DEFAULT 0,
    currency character varying(10) DEFAULT 'CNY'::character varying,
    payment_terms character varying(50),
    payment_due_date character varying(10),
    payment_method character varying(20),
    status character varying(20) DEFAULT 'PENDING'::character varying,
    approved_by bigint,
    approved_time timestamp with time zone,
    settlement_date character varying(10),
    remark text,
    tenant_id bigint NOT NULL,
    created_by character varying(50)
);


ALTER TABLE public.fin_sales_settlement OWNER TO postgres;

--
-- Name: fin_sales_settlement_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.fin_sales_settlement_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.fin_sales_settlement_id_seq OWNER TO postgres;

--
-- Name: fin_sales_settlement_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.fin_sales_settlement_id_seq OWNED BY public.fin_sales_settlement.id;


--
-- Name: fin_sales_settlement_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.fin_sales_settlement_item (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    settlement_id bigint NOT NULL,
    line_no bigint,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    specification character varying(200),
    unit character varying(20),
    invoice_qty numeric(18,3),
    shipped_qty numeric(18,3),
    settled_qty numeric(18,3),
    this_settle_qty numeric(18,3),
    unit_price numeric(18,4),
    tax_rate numeric(5,2),
    goods_amount numeric(18,2),
    tax_amount numeric(18,2),
    line_amount numeric(18,2),
    batch_no character varying(50),
    remark text,
    tenant_id bigint NOT NULL
);


ALTER TABLE public.fin_sales_settlement_item OWNER TO postgres;

--
-- Name: fin_sales_settlement_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.fin_sales_settlement_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.fin_sales_settlement_item_id_seq OWNER TO postgres;

--
-- Name: fin_sales_settlement_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.fin_sales_settlement_item_id_seq OWNED BY public.fin_sales_settlement_item.id;


--
-- Name: fin_statement_detail; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.fin_statement_detail (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    statement_id bigint NOT NULL,
    detail_type character varying(20) NOT NULL,
    related_no character varying(50),
    biz_date character varying(10),
    amount numeric(18,2),
    remark text,
    tenant_id bigint NOT NULL
);


ALTER TABLE public.fin_statement_detail OWNER TO postgres;

--
-- Name: fin_statement_detail_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.fin_statement_detail_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.fin_statement_detail_id_seq OWNER TO postgres;

--
-- Name: fin_statement_detail_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.fin_statement_detail_id_seq OWNED BY public.fin_statement_detail.id;


--
-- Name: fin_supplier_statement; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.fin_supplier_statement (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    statement_no character varying(50) NOT NULL,
    supplier_id bigint,
    supplier_name character varying(100),
    statement_period character varying(20) NOT NULL,
    start_date character varying(10) NOT NULL,
    end_date character varying(10) NOT NULL,
    beginning_amount numeric(18,2) DEFAULT 0,
    purchase_amount numeric(18,2) DEFAULT 0,
    other_amount numeric(18,2) DEFAULT 0,
    payment_amount numeric(18,2) DEFAULT 0,
    discount_amount numeric(18,2) DEFAULT 0,
    adjustment_amount numeric(18,2) DEFAULT 0,
    ending_amount numeric(18,2) DEFAULT 0,
    status character varying(20) DEFAULT 'PENDING'::character varying,
    confirmed_by bigint,
    confirmed_time timestamp with time zone,
    dispute_reason text,
    tenant_id bigint NOT NULL,
    created_by character varying(50)
);


ALTER TABLE public.fin_supplier_statement OWNER TO postgres;

--
-- Name: fin_supplier_statement_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.fin_supplier_statement_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.fin_supplier_statement_id_seq OWNER TO postgres;

--
-- Name: fin_supplier_statement_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.fin_supplier_statement_id_seq OWNED BY public.fin_supplier_statement.id;


--
-- Name: fqcs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.fqcs (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    fqc_no character varying(50) NOT NULL,
    order_id bigint,
    order_no character varying(50),
    quantity numeric(18,4),
    sample_size bigint,
    qualified_qty numeric(18,4),
    rejected_qty numeric(18,4) DEFAULT 0,
    check_user_id bigint,
    check_user_name character varying(50),
    check_date timestamp(6) with time zone,
    result bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.fqcs OWNER TO postgres;

--
-- Name: fqcs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.fqcs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.fqcs_id_seq OWNER TO postgres;

--
-- Name: fqcs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.fqcs_id_seq OWNED BY public.fqcs.id;


--
-- Name: integration_erp_mappings; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.integration_erp_mappings (
    id bigint NOT NULL,
    erp_table_name character varying(50) NOT NULL,
    erp_field_name character varying(50) NOT NULL,
    mes_table_name character varying(50) NOT NULL,
    mes_field_name character varying(50) NOT NULL,
    transform_rule character varying(200),
    tenant_id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.integration_erp_mappings OWNER TO postgres;

--
-- Name: integration_erp_mappings_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.integration_erp_mappings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.integration_erp_mappings_id_seq OWNER TO postgres;

--
-- Name: integration_erp_mappings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.integration_erp_mappings_id_seq OWNED BY public.integration_erp_mappings.id;


--
-- Name: integration_erp_sync_logs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.integration_erp_sync_logs (
    id bigint NOT NULL,
    sync_type character varying(50) NOT NULL,
    direction character varying(10) NOT NULL,
    erp_bill_no character varying(100),
    mes_bill_no character varying(100),
    request_body text,
    response_body text,
    status character varying(20),
    error_msg text,
    retry_count bigint DEFAULT 0,
    tenant_id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.integration_erp_sync_logs OWNER TO postgres;

--
-- Name: integration_erp_sync_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.integration_erp_sync_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.integration_erp_sync_logs_id_seq OWNER TO postgres;

--
-- Name: integration_erp_sync_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.integration_erp_sync_logs_id_seq OWNED BY public.integration_erp_sync_logs.id;


--
-- Name: inventories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.inventories (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    warehouse_id bigint,
    location_id bigint,
    quantity numeric(18,4) DEFAULT 0,
    available_qty numeric(18,4) DEFAULT 0,
    allocated_qty numeric(18,4) DEFAULT 0,
    locked_qty numeric(18,4) DEFAULT 0,
    batch_no character varying(50)
);


ALTER TABLE public.inventories OWNER TO postgres;

--
-- Name: inventories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.inventories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.inventories_id_seq OWNER TO postgres;

--
-- Name: inventories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.inventories_id_seq OWNED BY public.inventories.id;


--
-- Name: inventory_records; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.inventory_records (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    record_no character varying(50) NOT NULL,
    record_type character varying(20),
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    warehouse_id bigint,
    location_id bigint,
    quantity numeric(18,4),
    batch_no character varying(50),
    source_no character varying(50),
    operator_id bigint,
    operator_name character varying(50),
    operate_time timestamp(6) with time zone,
    remark character varying(500)
);


ALTER TABLE public.inventory_records OWNER TO postgres;

--
-- Name: inventory_records_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.inventory_records_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.inventory_records_id_seq OWNER TO postgres;

--
-- Name: inventory_records_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.inventory_records_id_seq OWNED BY public.inventory_records.id;


--
-- Name: ip_qcs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ip_qcs (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    ip_qc_no character varying(50) NOT NULL,
    order_id bigint,
    order_no character varying(50),
    process_id bigint,
    process_name character varying(100),
    quantity numeric(18,4),
    sample_size bigint,
    check_user_id bigint,
    check_user_name character varying(50),
    check_date timestamp(6) with time zone,
    result bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.ip_qcs OWNER TO postgres;

--
-- Name: ip_qcs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ip_qcs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.ip_qcs_id_seq OWNER TO postgres;

--
-- Name: ip_qcs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ip_qcs_id_seq OWNED BY public.ip_qcs.id;


--
-- Name: iqc_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.iqc_items (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    iqc_id bigint,
    check_item_id bigint,
    check_item character varying(100),
    check_standard character varying(200),
    check_method character varying(100),
    result bigint DEFAULT 1,
    remark character varying(200)
);


ALTER TABLE public.iqc_items OWNER TO postgres;

--
-- Name: iqc_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.iqc_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.iqc_items_id_seq OWNER TO postgres;

--
-- Name: iqc_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.iqc_items_id_seq OWNED BY public.iqc_items.id;


--
-- Name: iqcs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.iqcs (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    iqc_no character varying(50) NOT NULL,
    supplier_id bigint,
    supplier_name character varying(100),
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    quantity numeric(18,4),
    unit character varying(20),
    check_user_id bigint,
    check_user_name character varying(50),
    check_date timestamp(6) with time zone,
    result bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.iqcs OWNER TO postgres;

--
-- Name: iqcs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.iqcs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.iqcs_id_seq OWNER TO postgres;

--
-- Name: iqcs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.iqcs_id_seq OWNED BY public.iqcs.id;


--
-- Name: lab_calibration; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.lab_calibration (
    id bigint NOT NULL,
    tenant_id bigint DEFAULT 1 NOT NULL,
    instrument_id bigint NOT NULL,
    calibration_date timestamp with time zone,
    calibration_result character varying(10),
    calibrated_by character varying(50),
    certificate_no character varying(100),
    next_calibration_date timestamp with time zone,
    calibration_items text,
    attachment_url character varying(500),
    remark character varying(500),
    created_by character varying(50),
    updated_by character varying(50),
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.lab_calibration OWNER TO postgres;

--
-- Name: lab_calibration_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.lab_calibration_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lab_calibration_id_seq OWNER TO postgres;

--
-- Name: lab_calibration_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.lab_calibration_id_seq OWNED BY public.lab_calibration.id;


--
-- Name: lab_instrument; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.lab_instrument (
    id bigint NOT NULL,
    tenant_id bigint DEFAULT 1 NOT NULL,
    instrument_code character varying(50),
    instrument_name character varying(100),
    instrument_type character varying(20),
    manufacturer character varying(100),
    model character varying(100),
    serial_number character varying(100),
    calibration_cycle bigint,
    last_calibration_date timestamp with time zone,
    next_calibration_date timestamp with time zone,
    calibration_status character varying(20),
    lab_location character varying(100),
    status character varying(20),
    remark character varying(500),
    created_by character varying(50),
    updated_by character varying(50),
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.lab_instrument OWNER TO postgres;

--
-- Name: lab_instrument_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.lab_instrument_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lab_instrument_id_seq OWNER TO postgres;

--
-- Name: lab_instrument_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.lab_instrument_id_seq OWNED BY public.lab_instrument.id;


--
-- Name: lab_report; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.lab_report (
    id bigint NOT NULL,
    tenant_id bigint DEFAULT 1 NOT NULL,
    report_no character varying(50),
    sample_id bigint NOT NULL,
    conclusion character varying(20),
    remarks text,
    attachments text,
    inspector_id bigint,
    inspector_name character varying(50),
    approved_by character varying(50),
    report_date timestamp with time zone,
    created_by character varying(50),
    updated_by character varying(50),
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.lab_report OWNER TO postgres;

--
-- Name: lab_report_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.lab_report_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lab_report_id_seq OWNER TO postgres;

--
-- Name: lab_report_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.lab_report_id_seq OWNED BY public.lab_report.id;


--
-- Name: lab_sample; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.lab_sample (
    id bigint NOT NULL,
    tenant_id bigint DEFAULT 1 NOT NULL,
    sample_code character varying(50),
    sample_name character varying(100),
    inspection_type character varying(20),
    source_type character varying(20),
    source_id bigint,
    source_no character varying(50),
    quantity numeric,
    sample_qty numeric,
    received_by character varying(50),
    received_at timestamp with time zone,
    due_date timestamp with time zone,
    status character varying(20),
    remark character varying(500),
    created_by character varying(50),
    updated_by character varying(50),
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.lab_sample OWNER TO postgres;

--
-- Name: lab_sample_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.lab_sample_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lab_sample_id_seq OWNER TO postgres;

--
-- Name: lab_sample_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.lab_sample_id_seq OWNED BY public.lab_sample.id;


--
-- Name: lab_test_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.lab_test_item (
    id bigint NOT NULL,
    tenant_id bigint DEFAULT 1 NOT NULL,
    sample_id bigint NOT NULL,
    item_name character varying(100),
    test_method character varying(100),
    standard character varying(100),
    upper_limit numeric,
    lower_limit numeric,
    unit character varying(20),
    result character varying(20),
    result_value numeric,
    tester_id bigint,
    tester_name character varying(50),
    tested_at timestamp with time zone,
    remark character varying(200),
    created_at timestamp with time zone
);


ALTER TABLE public.lab_test_item OWNER TO postgres;

--
-- Name: lab_test_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.lab_test_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lab_test_item_id_seq OWNER TO postgres;

--
-- Name: lab_test_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.lab_test_item_id_seq OWNED BY public.lab_test_item.id;


--
-- Name: locations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.locations (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    location_code character varying(50) NOT NULL,
    location_name character varying(100),
    warehouse_id bigint,
    zone_code character varying(20),
    "row" bigint,
    col bigint,
    layer bigint,
    location_type character varying(20),
    capacity bigint,
    status bigint DEFAULT 1
);


ALTER TABLE public.locations OWNER TO postgres;

--
-- Name: locations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.locations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.locations_id_seq OWNER TO postgres;

--
-- Name: locations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.locations_id_seq OWNED BY public.locations.id;


--
-- Name: login_logs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.login_logs (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint,
    username character varying(50),
    ip character varying(50),
    login_location character varying(255),
    browser character varying(50),
    os character varying(50),
    status bigint DEFAULT 0,
    msg character varying(255),
    login_time timestamp(6) with time zone
);


ALTER TABLE public.login_logs OWNER TO postgres;

--
-- Name: login_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.login_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.login_logs_id_seq OWNER TO postgres;

--
-- Name: login_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.login_logs_id_seq OWNED BY public.login_logs.id;


--
-- Name: material_categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.material_categories (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    parent_id bigint DEFAULT 0,
    category_name character varying(100) NOT NULL,
    category_code character varying(50),
    sort bigint DEFAULT 0,
    status bigint DEFAULT 1
);


ALTER TABLE public.material_categories OWNER TO postgres;

--
-- Name: material_categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.material_categories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.material_categories_id_seq OWNER TO postgres;

--
-- Name: material_categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.material_categories_id_seq OWNED BY public.material_categories.id;


--
-- Name: materials; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.materials (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    material_code character varying(50) NOT NULL,
    material_name character varying(100) NOT NULL,
    material_type character varying(20),
    spec character varying(100),
    unit character varying(20),
    unit_name character varying(20),
    weight numeric,
    length numeric,
    width numeric,
    height numeric,
    category_id bigint,
    status bigint DEFAULT 1
);


ALTER TABLE public.materials OWNER TO postgres;

--
-- Name: materials_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.materials_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.materials_id_seq OWNER TO postgres;

--
-- Name: materials_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.materials_id_seq OWNED BY public.materials.id;


--
-- Name: mdm_bom; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mdm_bom (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    bom_code character varying(50) NOT NULL,
    bom_name character varying(200) NOT NULL,
    material_id bigint NOT NULL,
    material_code character varying(50),
    material_name character varying(100),
    version character varying(20),
    status character varying(20) DEFAULT 'DRAFT'::character varying,
    eff_date date,
    exp_date date,
    remark character varying(500),
    product_id bigint NOT NULL,
    product_code character varying(50),
    product_name character varying(100),
    bom_type character varying(20)
);


ALTER TABLE public.mdm_bom OWNER TO postgres;

--
-- Name: mdm_bom_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mdm_bom_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mdm_bom_id_seq OWNER TO postgres;

--
-- Name: mdm_bom_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mdm_bom_id_seq OWNED BY public.mdm_bom.id;


--
-- Name: mdm_bom_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mdm_bom_item (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    bom_id bigint NOT NULL,
    line_no bigint DEFAULT 0,
    material_id bigint NOT NULL,
    material_code character varying(50),
    material_name character varying(100),
    quantity numeric(18,4) DEFAULT 0,
    unit character varying(20),
    scrap_rate numeric(10,4) DEFAULT 0,
    substitute_group character varying(50),
    is_alternative bigint DEFAULT 0
);


ALTER TABLE public.mdm_bom_item OWNER TO postgres;

--
-- Name: mdm_bom_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mdm_bom_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mdm_bom_item_id_seq OWNER TO postgres;

--
-- Name: mdm_bom_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mdm_bom_item_id_seq OWNED BY public.mdm_bom_item.id;


--
-- Name: mdm_customer; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mdm_customer (
    id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    deleted_at timestamp without time zone,
    tenant_id bigint NOT NULL,
    code character varying(50) NOT NULL,
    name character varying(200) NOT NULL,
    type character varying(50),
    contact character varying(100),
    phone character varying(50),
    email character varying(100),
    address character varying(500),
    status integer DEFAULT 1
);


ALTER TABLE public.mdm_customer OWNER TO postgres;

--
-- Name: mdm_customer_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mdm_customer_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mdm_customer_id_seq OWNER TO postgres;

--
-- Name: mdm_customer_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mdm_customer_id_seq OWNED BY public.mdm_customer.id;


--
-- Name: mdm_material; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mdm_material (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    material_code character varying(50) NOT NULL,
    material_name character varying(100) NOT NULL,
    material_type character varying(20),
    spec character varying(100),
    unit character varying(20),
    unit_name character varying(20),
    weight numeric(18,4),
    length numeric(18,4),
    width numeric(18,4),
    height numeric(18,4),
    category_id bigint,
    status bigint DEFAULT 1
);


ALTER TABLE public.mdm_material OWNER TO postgres;

--
-- Name: mdm_material_category; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mdm_material_category (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    parent_id bigint DEFAULT 0,
    category_name character varying(100) NOT NULL,
    category_code character varying(50),
    sort bigint DEFAULT 0,
    status bigint DEFAULT 1
);


ALTER TABLE public.mdm_material_category OWNER TO postgres;

--
-- Name: mdm_material_category_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mdm_material_category_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mdm_material_category_id_seq OWNER TO postgres;

--
-- Name: mdm_material_category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mdm_material_category_id_seq OWNED BY public.mdm_material_category.id;


--
-- Name: mdm_material_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mdm_material_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mdm_material_id_seq OWNER TO postgres;

--
-- Name: mdm_material_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mdm_material_id_seq OWNED BY public.mdm_material.id;


--
-- Name: mdm_operation; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mdm_operation (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    operation_code character varying(50) NOT NULL,
    operation_name character varying(100) NOT NULL,
    workcenter_id bigint,
    workcenter_name character varying(100),
    standard_worktime bigint DEFAULT 0,
    quality_std character varying(500),
    is_key_process bigint DEFAULT 0,
    is_qc_point bigint DEFAULT 0,
    sequence bigint DEFAULT 0,
    remark character varying(500)
);


ALTER TABLE public.mdm_operation OWNER TO postgres;

--
-- Name: mdm_operation_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mdm_operation_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mdm_operation_id_seq OWNER TO postgres;

--
-- Name: mdm_operation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mdm_operation_id_seq OWNED BY public.mdm_operation.id;


--
-- Name: mdm_production_line; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mdm_production_line (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    line_code character varying(50) NOT NULL,
    line_name character varying(100) NOT NULL,
    workshop_id bigint NOT NULL,
    line_type character varying(20),
    status bigint DEFAULT 1
);


ALTER TABLE public.mdm_production_line OWNER TO postgres;

--
-- Name: mdm_production_line_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mdm_production_line_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mdm_production_line_id_seq OWNER TO postgres;

--
-- Name: mdm_production_line_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mdm_production_line_id_seq OWNED BY public.mdm_production_line.id;


--
-- Name: mdm_shift; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mdm_shift (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    shift_code character varying(50) NOT NULL,
    shift_name character varying(100) NOT NULL,
    start_time character varying(10),
    end_time character varying(10),
    break_start character varying(10),
    break_end character varying(10),
    status integer DEFAULT 1,
    work_hours numeric(10,2) DEFAULT 8,
    is_night bigint DEFAULT 0,
    remark character varying(500)
);


ALTER TABLE public.mdm_shift OWNER TO postgres;

--
-- Name: mdm_shift_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mdm_shift_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mdm_shift_id_seq OWNER TO postgres;

--
-- Name: mdm_shift_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mdm_shift_id_seq OWNED BY public.mdm_shift.id;


--
-- Name: mdm_supplier; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mdm_supplier (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    code character varying(50) NOT NULL,
    name character varying(100) NOT NULL,
    type character varying(20),
    contact character varying(100),
    phone character varying(20),
    email character varying(100),
    address character varying(200),
    category character varying(50),
    level bigint DEFAULT 1,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.mdm_supplier OWNER TO postgres;

--
-- Name: mdm_supplier_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mdm_supplier_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mdm_supplier_id_seq OWNER TO postgres;

--
-- Name: mdm_supplier_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mdm_supplier_id_seq OWNED BY public.mdm_supplier.id;


--
-- Name: mdm_supplier_material; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mdm_supplier_material (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    supplier_id bigint NOT NULL,
    material_id bigint NOT NULL,
    material_code character varying(50),
    material_name character varying(100),
    supplier_part_no character varying(50),
    price numeric(18,6),
    currency character varying(10),
    min_order_qty numeric DEFAULT 0,
    lead_time bigint DEFAULT 0,
    is_preferred bigint DEFAULT 0,
    status bigint DEFAULT 1,
    tenant_id bigint NOT NULL
);


ALTER TABLE public.mdm_supplier_material OWNER TO postgres;

--
-- Name: mdm_supplier_material_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mdm_supplier_material_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mdm_supplier_material_id_seq OWNER TO postgres;

--
-- Name: mdm_supplier_material_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mdm_supplier_material_id_seq OWNED BY public.mdm_supplier_material.id;


--
-- Name: mdm_workshop; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mdm_workshop (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    workshop_code character varying(50) NOT NULL,
    workshop_name character varying(100) NOT NULL,
    workshop_type character varying(20),
    manager character varying(50),
    phone character varying(20),
    address character varying(200),
    status bigint DEFAULT 1
);


ALTER TABLE public.mdm_workshop OWNER TO postgres;

--
-- Name: mdm_workshop_config; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mdm_workshop_config (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    workshop_id bigint NOT NULL,
    erp_plant_code character varying(50),
    max_devices bigint DEFAULT 0,
    max_workers bigint DEFAULT 0,
    max_capacity_per_day numeric(18,2),
    time_zone character varying(20) DEFAULT 'Asia/Shanghai'::character varying,
    is_default bigint DEFAULT 0
);


ALTER TABLE public.mdm_workshop_config OWNER TO postgres;

--
-- Name: mdm_workshop_config_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mdm_workshop_config_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mdm_workshop_config_id_seq OWNER TO postgres;

--
-- Name: mdm_workshop_config_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mdm_workshop_config_id_seq OWNED BY public.mdm_workshop_config.id;


--
-- Name: mdm_workshop_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mdm_workshop_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mdm_workshop_id_seq OWNER TO postgres;

--
-- Name: mdm_workshop_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mdm_workshop_id_seq OWNED BY public.mdm_workshop.id;


--
-- Name: mdm_workstation; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mdm_workstation (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    station_code character varying(50) NOT NULL,
    station_name character varying(100) NOT NULL,
    line_id bigint NOT NULL,
    station_type character varying(20),
    status bigint DEFAULT 1
);


ALTER TABLE public.mdm_workstation OWNER TO postgres;

--
-- Name: mdm_workstation_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mdm_workstation_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mdm_workstation_id_seq OWNER TO postgres;

--
-- Name: mdm_workstation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mdm_workstation_id_seq OWNED BY public.mdm_workstation.id;


--
-- Name: mes_code_rule; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_code_rule (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    rule_code character varying(50) NOT NULL,
    rule_name character varying(100) NOT NULL,
    entity_type character varying(50) NOT NULL,
    prefix character varying(20),
    date_format character varying(20),
    seq_length bigint DEFAULT 4,
    seq_start bigint DEFAULT 1,
    seq_current bigint DEFAULT 0,
    mid_fix character varying(20),
    suffix character varying(20),
    reset_type character varying(20),
    last_gen_date character varying(20),
    example character varying(100),
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.mes_code_rule OWNER TO postgres;

--
-- Name: mes_code_rule_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_code_rule_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_code_rule_id_seq OWNER TO postgres;

--
-- Name: mes_code_rule_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_code_rule_id_seq OWNED BY public.mes_code_rule.id;


--
-- Name: mes_code_rule_record; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_code_rule_record (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    rule_id bigint NOT NULL,
    rule_code character varying(50),
    entity_type character varying(50),
    gen_date character varying(20),
    seq_value bigint,
    gen_code character varying(100)
);


ALTER TABLE public.mes_code_rule_record OWNER TO postgres;

--
-- Name: mes_code_rule_record_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_code_rule_record_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_code_rule_record_id_seq OWNER TO postgres;

--
-- Name: mes_code_rule_record_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_code_rule_record_id_seq OWNED BY public.mes_code_rule_record.id;


--
-- Name: mes_container_lifecycle; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_container_lifecycle (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    tenant_id bigint DEFAULT 1 NOT NULL,
    container_id bigint NOT NULL,
    container_code character varying(50),
    event_type character varying(20) NOT NULL,
    event_date timestamp with time zone NOT NULL,
    operator_id bigint,
    operator_name character varying(50),
    location_id bigint,
    location_name character varying(100),
    status character varying(20),
    remark character varying(500)
);


ALTER TABLE public.mes_container_lifecycle OWNER TO postgres;

--
-- Name: mes_container_lifecycle_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_container_lifecycle_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_container_lifecycle_id_seq OWNER TO postgres;

--
-- Name: mes_container_lifecycle_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_container_lifecycle_id_seq OWNED BY public.mes_container_lifecycle.id;


--
-- Name: mes_container_maintenance; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_container_maintenance (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    tenant_id bigint DEFAULT 1 NOT NULL,
    container_id bigint NOT NULL,
    container_code character varying(50),
    maintenance_type character varying(20) NOT NULL,
    maintenance_date timestamp with time zone NOT NULL,
    completed_date timestamp with time zone,
    technician_id bigint,
    technician_name character varying(50),
    fault_description character varying(500),
    maintenance_content character varying(500),
    spare_parts_used jsonb,
    cost numeric(18,2),
    status character varying(20) DEFAULT 'PENDING'::character varying,
    remark character varying(500)
);


ALTER TABLE public.mes_container_maintenance OWNER TO postgres;

--
-- Name: mes_container_maintenance_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_container_maintenance_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_container_maintenance_id_seq OWNER TO postgres;

--
-- Name: mes_container_maintenance_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_container_maintenance_id_seq OWNED BY public.mes_container_maintenance.id;


--
-- Name: mes_electronic_sop; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_electronic_sop (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    sop_no character varying(50) NOT NULL,
    sop_name character varying(200) NOT NULL,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    version character varying(20),
    process_id bigint,
    process_name character varying(100),
    content_type character varying(20),
    content_url character varying(500),
    thumbnail_url character varying(500),
    steps jsonb,
    workstation_id bigint,
    workstation_name character varying(100),
    workshop_id bigint,
    workshop_name character varying(100),
    status bigint DEFAULT 1,
    eff_date timestamp with time zone,
    exp_date timestamp with time zone,
    approved_by character varying(50),
    approved_at timestamp with time zone,
    remark character varying(500)
);


ALTER TABLE public.mes_electronic_sop OWNER TO postgres;

--
-- Name: mes_electronic_sop_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_electronic_sop_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_electronic_sop_id_seq OWNER TO postgres;

--
-- Name: mes_electronic_sop_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_electronic_sop_id_seq OWNED BY public.mes_electronic_sop.id;


--
-- Name: mes_first_last_inspect; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_first_last_inspect (
    id bigint NOT NULL,
    tenant_id bigint NOT NULL,
    workshop_id bigint,
    inspect_no character varying(50) NOT NULL,
    inspect_type character varying(10) NOT NULL,
    production_order_id bigint NOT NULL,
    process_id bigint NOT NULL,
    workstation_id bigint NOT NULL,
    shift_id bigint NOT NULL,
    product_id bigint NOT NULL,
    serial_no character varying(100),
    inspect_items jsonb,
    overall_result character varying(10),
    inspector_id bigint,
    inspector_name character varying(50),
    inspect_time timestamp with time zone,
    bluetooth_device_id character varying(100),
    remark character varying(500),
    status character varying(20) DEFAULT 'PENDING'::character varying,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.mes_first_last_inspect OWNER TO postgres;

--
-- Name: mes_first_last_inspect_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_first_last_inspect_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_first_last_inspect_id_seq OWNER TO postgres;

--
-- Name: mes_first_last_inspect_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_first_last_inspect_id_seq OWNED BY public.mes_first_last_inspect.id;


--
-- Name: mes_flow_card; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_flow_card (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    card_no character varying(50) NOT NULL,
    order_id bigint,
    order_no character varying(50),
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    workshop_id bigint,
    workshop_name character varying(100),
    line_id bigint,
    line_name character varying(100),
    process_id bigint,
    process_name character varying(100),
    station_id bigint,
    station_name character varying(100),
    plan_qty numeric(18,4),
    completed_qty numeric(18,4),
    status bigint DEFAULT 1,
    priority bigint DEFAULT 1,
    plan_start_time timestamp with time zone,
    plan_end_time timestamp with time zone,
    actual_start_time timestamp with time zone,
    actual_end_time timestamp with time zone,
    remark character varying(500)
);


ALTER TABLE public.mes_flow_card OWNER TO postgres;

--
-- Name: mes_flow_card_detail; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_flow_card_detail (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    card_id bigint NOT NULL,
    step_no bigint,
    process_id bigint,
    process_name character varying(100),
    station_id bigint,
    station_name character varying(100),
    work_content character varying(500),
    std_cycle_time bigint,
    seq_qty numeric(18,4)
);


ALTER TABLE public.mes_flow_card_detail OWNER TO postgres;

--
-- Name: mes_flow_card_detail_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_flow_card_detail_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_flow_card_detail_id_seq OWNER TO postgres;

--
-- Name: mes_flow_card_detail_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_flow_card_detail_id_seq OWNED BY public.mes_flow_card_detail.id;


--
-- Name: mes_flow_card_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_flow_card_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_flow_card_id_seq OWNER TO postgres;

--
-- Name: mes_flow_card_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_flow_card_id_seq OWNED BY public.mes_flow_card.id;


--
-- Name: mes_job_report_log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_job_report_log (
    id bigint NOT NULL,
    tenant_id bigint,
    work_order_id bigint,
    work_order_code character varying(50),
    process_code character varying(50),
    process_name character varying(100),
    report_type character varying(20),
    quantity numeric(18,2),
    report_time timestamp with time zone,
    reporter_id bigint,
    reporter_name character varying(100),
    remark character varying(500),
    status character varying(20),
    created_at timestamp with time zone
);


ALTER TABLE public.mes_job_report_log OWNER TO postgres;

--
-- Name: mes_job_report_log_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_job_report_log_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_job_report_log_id_seq OWNER TO postgres;

--
-- Name: mes_job_report_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_job_report_log_id_seq OWNED BY public.mes_job_report_log.id;


--
-- Name: mes_order_day; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_order_day (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    day_plan_no character varying(50) NOT NULL,
    plan_date date NOT NULL,
    month_plan_id bigint,
    month_plan_no character varying(50),
    workshop_id bigint,
    workshop_name character varying(100),
    production_line_id bigint,
    line_name character varying(100),
    total_product_count bigint DEFAULT 0,
    total_plan_qty numeric(18,3) DEFAULT 0,
    total_completed_qty numeric(18,3) DEFAULT 0,
    kit_status character varying(20) DEFAULT 'PENDING'::character varying,
    kit_check_time timestamp with time zone,
    kit_check_by bigint,
    status character varying(20) DEFAULT 'DRAFT'::character varying,
    shift_type character varying(20),
    published_at timestamp with time zone,
    published_by bigint,
    remark text,
    created_by character varying(50)
);


ALTER TABLE public.mes_order_day OWNER TO postgres;

--
-- Name: mes_order_day_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_order_day_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_order_day_id_seq OWNER TO postgres;

--
-- Name: mes_order_day_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_order_day_id_seq OWNED BY public.mes_order_day.id;


--
-- Name: mes_order_day_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_order_day_item (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    day_plan_id bigint NOT NULL,
    line_no bigint NOT NULL,
    product_id bigint NOT NULL,
    product_code character varying(50),
    product_name character varying(100),
    specification character varying(200),
    unit character varying(20),
    plan_qty numeric(18,3) NOT NULL,
    completed_qty numeric(18,3) DEFAULT 0,
    bom_id bigint,
    bom_version character varying(20),
    process_route_id bigint,
    route_version character varying(20),
    production_mode character varying(20) DEFAULT 'BATCH'::character varying,
    kit_status character varying(20) DEFAULT 'PENDING'::character varying,
    kit_check_remark text,
    month_plan_item_id bigint,
    item_status character varying(20) DEFAULT 'PENDING'::character varying,
    priority bigint DEFAULT 5,
    remark text,
    tenant_id bigint NOT NULL
);


ALTER TABLE public.mes_order_day_item OWNER TO postgres;

--
-- Name: mes_order_day_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_order_day_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_order_day_item_id_seq OWNER TO postgres;

--
-- Name: mes_order_day_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_order_day_item_id_seq OWNED BY public.mes_order_day_item.id;


--
-- Name: mes_order_month; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_order_month (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    month_plan_no character varying(50) NOT NULL,
    plan_month character varying(7) NOT NULL,
    title character varying(200),
    source_type character varying(20),
    source_no character varying(50),
    workshop_id bigint,
    workshop_name character varying(100),
    total_product_count bigint DEFAULT 0,
    total_plan_qty numeric(18,3) DEFAULT 0,
    total_completed_qty numeric(18,3) DEFAULT 0,
    approval_status character varying(20) DEFAULT 'DRAFT'::character varying,
    submitted_by bigint,
    submitted_at timestamp with time zone,
    approved_by bigint,
    approved_at timestamp with time zone,
    released_by bigint,
    released_at timestamp with time zone,
    remark text,
    created_by character varying(50)
);


ALTER TABLE public.mes_order_month OWNER TO postgres;

--
-- Name: mes_order_month_audit; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_order_month_audit (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    month_plan_id bigint NOT NULL,
    approval_status character varying(20) NOT NULL,
    approver_id bigint,
    approver_name character varying(50),
    approval_time timestamp with time zone,
    comment text,
    tenant_id bigint NOT NULL
);


ALTER TABLE public.mes_order_month_audit OWNER TO postgres;

--
-- Name: mes_order_month_audit_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_order_month_audit_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_order_month_audit_id_seq OWNER TO postgres;

--
-- Name: mes_order_month_audit_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_order_month_audit_id_seq OWNED BY public.mes_order_month_audit.id;


--
-- Name: mes_order_month_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_order_month_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_order_month_id_seq OWNER TO postgres;

--
-- Name: mes_order_month_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_order_month_id_seq OWNED BY public.mes_order_month.id;


--
-- Name: mes_order_month_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_order_month_item (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    month_plan_id bigint NOT NULL,
    line_no bigint NOT NULL,
    product_id bigint NOT NULL,
    product_code character varying(50),
    product_name character varying(100),
    specification character varying(200),
    unit character varying(20),
    plan_qty numeric(18,3) NOT NULL,
    completed_qty numeric(18,3) DEFAULT 0,
    released_qty numeric(18,3) DEFAULT 0,
    delivery_date timestamp with time zone,
    priority bigint DEFAULT 5,
    remark text,
    tenant_id bigint NOT NULL
);


ALTER TABLE public.mes_order_month_item OWNER TO postgres;

--
-- Name: mes_order_month_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_order_month_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_order_month_item_id_seq OWNER TO postgres;

--
-- Name: mes_order_month_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_order_month_item_id_seq OWNED BY public.mes_order_month_item.id;


--
-- Name: mes_package; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_package (
    id bigint NOT NULL,
    tenant_id bigint NOT NULL,
    workshop_id bigint,
    package_no character varying(100) NOT NULL,
    package_type character varying(20),
    production_order_id bigint,
    product_id bigint NOT NULL,
    product_code character varying(50),
    qty bigint NOT NULL,
    serial_nos jsonb,
    status character varying(20) DEFAULT 'OPEN'::character varying,
    seal_time timestamp with time zone,
    seal_by character varying(50),
    ship_time timestamp with time zone,
    customer_id bigint,
    container_id bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.mes_package OWNER TO postgres;

--
-- Name: mes_package_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_package_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_package_id_seq OWNER TO postgres;

--
-- Name: mes_package_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_package_id_seq OWNED BY public.mes_package.id;


--
-- Name: mes_person_skill; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_person_skill (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    person_id bigint NOT NULL,
    person_code character varying(50) NOT NULL,
    person_name character varying(100) NOT NULL,
    workshop_id bigint NOT NULL,
    workstation_id bigint,
    skill_level character varying(20) NOT NULL,
    certified_date date,
    expiry_date date,
    status character varying(20) DEFAULT 'ACTIVE'::character varying,
    created_by bigint
);


ALTER TABLE public.mes_person_skill OWNER TO postgres;

--
-- Name: mes_person_skill_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_person_skill_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_person_skill_id_seq OWNER TO postgres;

--
-- Name: mes_person_skill_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_person_skill_id_seq OWNED BY public.mes_person_skill.id;


--
-- Name: mes_person_skill_score; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_person_skill_score (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    person_skill_id bigint NOT NULL,
    skill_type character varying(20) NOT NULL,
    score numeric(5,2),
    evaluated_by bigint,
    evaluator_name character varying(50),
    evaluated_at timestamp without time zone,
    remark character varying(500)
);


ALTER TABLE public.mes_person_skill_score OWNER TO postgres;

--
-- Name: mes_person_skill_score_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_person_skill_score_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_person_skill_score_id_seq OWNER TO postgres;

--
-- Name: mes_person_skill_score_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_person_skill_score_id_seq OWNED BY public.mes_person_skill_score.id;


--
-- Name: mes_process; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_process (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    process_code character varying(50) NOT NULL,
    process_name character varying(200) NOT NULL,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    version character varying(20),
    status character varying(20) DEFAULT 'DRAFT'::character varying,
    eff_date date,
    exp_date date,
    remark character varying(500),
    is_current bigint DEFAULT 1
);


ALTER TABLE public.mes_process OWNER TO postgres;

--
-- Name: mes_process_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_process_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_process_id_seq OWNER TO postgres;

--
-- Name: mes_process_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_process_id_seq OWNED BY public.mes_process.id;


--
-- Name: mes_process_operation; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_process_operation (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    process_id bigint NOT NULL,
    operation_id bigint,
    operation_code character varying(50),
    operation_name character varying(100),
    line_no bigint NOT NULL,
    standard_worktime bigint DEFAULT 0,
    workcenter_id bigint,
    workcenter_name character varying(100),
    required_capacity numeric(10,2) DEFAULT 0,
    min_workers bigint DEFAULT 1,
    max_workers bigint DEFAULT 10,
    is_key_process bigint DEFAULT 0,
    is_qc_point bigint DEFAULT 0,
    quality_std character varying(500),
    status character varying(20) DEFAULT 'ACTIVE'::character varying,
    remark character varying(500)
);


ALTER TABLE public.mes_process_operation OWNER TO postgres;

--
-- Name: mes_process_operation_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_process_operation_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_process_operation_id_seq OWNER TO postgres;

--
-- Name: mes_process_operation_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_process_operation_id_seq OWNED BY public.mes_process_operation.id;


--
-- Name: mes_team; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_team (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    team_code character varying(50) NOT NULL,
    team_name character varying(100) NOT NULL,
    workshop_id bigint NOT NULL,
    shift_id bigint,
    leader_id bigint,
    phone character varying(20),
    remark character varying(500),
    status bigint DEFAULT 1
);


ALTER TABLE public.mes_team OWNER TO postgres;

--
-- Name: mes_team_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_team_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_team_id_seq OWNER TO postgres;

--
-- Name: mes_team_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_team_id_seq OWNED BY public.mes_team.id;


--
-- Name: mes_team_member; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_team_member (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    team_id bigint NOT NULL,
    user_id bigint NOT NULL,
    user_name character varying(50),
    role character varying(20),
    join_date timestamp with time zone,
    remark character varying(500),
    status bigint DEFAULT 1
);


ALTER TABLE public.mes_team_member OWNER TO postgres;

--
-- Name: mes_team_member_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_team_member_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_team_member_id_seq OWNER TO postgres;

--
-- Name: mes_team_member_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_team_member_id_seq OWNED BY public.mes_team_member.id;


--
-- Name: mes_team_shift; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_team_shift (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    team_id bigint NOT NULL,
    shift_id bigint NOT NULL,
    shift_date date NOT NULL,
    start_time character varying(10),
    end_time character varying(10),
    leader_id bigint,
    remark character varying(500)
);


ALTER TABLE public.mes_team_shift OWNER TO postgres;

--
-- Name: mes_team_shift_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_team_shift_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_team_shift_id_seq OWNER TO postgres;

--
-- Name: mes_team_shift_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_team_shift_id_seq OWNED BY public.mes_team_shift.id;


--
-- Name: mes_work_scheduling; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_work_scheduling (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    plan_no_day character varying(50),
    scheduling_code character varying(50),
    product_code character varying(50),
    product_name character varying(200),
    status character varying(20),
    quantity numeric(18,2),
    finished_qty numeric(18,2),
    work_mode character varying(20),
    task_mode character varying(20),
    plan_date date,
    workshop_code character varying(50),
    line_code character varying(50),
    created_by bigint,
    updated_by bigint
);


ALTER TABLE public.mes_work_scheduling OWNER TO postgres;

--
-- Name: mes_work_scheduling_detail; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mes_work_scheduling_detail (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint,
    scheduling_id bigint NOT NULL,
    working_node character varying(50),
    working_name character varying(100),
    status character varying(20),
    equipment_id bigint,
    equipment_code character varying(50),
    workstation_id bigint,
    workstation_name character varying(100),
    worker_id bigint,
    worker_name character varying(100),
    start_time timestamp with time zone,
    end_time timestamp with time zone,
    plan_qty numeric(18,2),
    finished_qty numeric(18,2),
    work_minutes bigint,
    created_by bigint,
    updated_by bigint
);


ALTER TABLE public.mes_work_scheduling_detail OWNER TO postgres;

--
-- Name: mes_work_scheduling_detail_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_work_scheduling_detail_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_work_scheduling_detail_id_seq OWNER TO postgres;

--
-- Name: mes_work_scheduling_detail_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_work_scheduling_detail_id_seq OWNED BY public.mes_work_scheduling_detail.id;


--
-- Name: mes_work_scheduling_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mes_work_scheduling_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mes_work_scheduling_id_seq OWNER TO postgres;

--
-- Name: mes_work_scheduling_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mes_work_scheduling_id_seq OWNED BY public.mes_work_scheduling.id;


--
-- Name: mps; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mps (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    mps_no character varying(50) NOT NULL,
    plan_month character varying(10),
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    quantity numeric(18,4),
    status bigint DEFAULT 1
);


ALTER TABLE public.mps OWNER TO postgres;

--
-- Name: mps_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mps_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mps_id_seq OWNER TO postgres;

--
-- Name: mps_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mps_id_seq OWNED BY public.mps.id;


--
-- Name: mrp_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mrp_items (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    mrp_id bigint,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    quantity numeric(18,4),
    stock_qty numeric(18,4),
    allocated_qty numeric(18,4),
    net_qty numeric(18,4),
    source_type character varying(20),
    source_no character varying(50)
);


ALTER TABLE public.mrp_items OWNER TO postgres;

--
-- Name: mrp_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mrp_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mrp_items_id_seq OWNER TO postgres;

--
-- Name: mrp_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mrp_items_id_seq OWNED BY public.mrp_items.id;


--
-- Name: mrps; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mrps (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    mrp_no character varying(50) NOT NULL,
    mrp_type character varying(20),
    plan_date timestamp(6) with time zone,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.mrps OWNER TO postgres;

--
-- Name: mrps_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.mrps_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.mrps_id_seq OWNER TO postgres;

--
-- Name: mrps_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.mrps_id_seq OWNED BY public.mrps.id;


--
-- Name: ncrs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.ncrs (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    ncr_no character varying(50) NOT NULL,
    defect_id bigint,
    source_type character varying(20),
    issue_desc text,
    root_cause text,
    corrective_action text,
    preventive_action text,
    verify_result character varying(200),
    verify_user_id bigint,
    verify_date timestamp(6) with time zone,
    status bigint DEFAULT 1
);


ALTER TABLE public.ncrs OWNER TO postgres;

--
-- Name: ncrs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.ncrs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.ncrs_id_seq OWNER TO postgres;

--
-- Name: ncrs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.ncrs_id_seq OWNED BY public.ncrs.id;


--
-- Name: notification_channel; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.notification_channel (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    channel_code character varying(50) NOT NULL,
    channel_name character varying(100) NOT NULL,
    channel_type character varying(20) NOT NULL,
    config jsonb,
    is_enabled bigint DEFAULT 1,
    priority bigint DEFAULT 0,
    tenant_id bigint NOT NULL,
    created_by character varying(50),
    updated_by character varying(50),
    updated_at timestamp with time zone
);


ALTER TABLE public.notification_channel OWNER TO postgres;

--
-- Name: notification_channel_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.notification_channel_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.notification_channel_id_seq OWNER TO postgres;

--
-- Name: notification_channel_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.notification_channel_id_seq OWNED BY public.notification_channel.id;


--
-- Name: oee_report; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.oee_report (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    report_date timestamp with time zone NOT NULL,
    workshop_id bigint,
    workshop_name character varying(100),
    line_id bigint,
    line_name character varying(100),
    availability numeric,
    performance numeric,
    quality numeric,
    oee numeric,
    planned_production_time bigint,
    actual_production_time bigint,
    down_time bigint,
    speed_loss numeric,
    defect_loss numeric,
    remark character varying(500)
);


ALTER TABLE public.oee_report OWNER TO postgres;

--
-- Name: oee_report_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.oee_report_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.oee_report_id_seq OWNER TO postgres;

--
-- Name: oee_report_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.oee_report_id_seq OWNED BY public.oee_report.id;


--
-- Name: oper_logs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.oper_logs (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint,
    title character varying(200),
    business_type character varying(20),
    method character varying(100),
    request_method character varying(10),
    operator_type bigint DEFAULT 1,
    oper_name character varying(50),
    dept_name character varying(100),
    oper_url character varying(255),
    oper_ip character varying(50),
    oper_location character varying(255),
    oper_param text,
    json_result text,
    status bigint DEFAULT 0,
    error_msg text,
    oper_time timestamp(6) with time zone
);


ALTER TABLE public.oper_logs OWNER TO postgres;

--
-- Name: oper_logs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.oper_logs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.oper_logs_id_seq OWNER TO postgres;

--
-- Name: oper_logs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.oper_logs_id_seq OWNED BY public.oper_logs.id;


--
-- Name: oqcs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.oqcs (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    oqc_no character varying(50) NOT NULL,
    shipping_no character varying(50),
    customer_id bigint,
    customer_name character varying(100),
    quantity numeric(18,4),
    check_user_id bigint,
    check_user_name character varying(50),
    check_date timestamp(6) with time zone,
    result bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.oqcs OWNER TO postgres;

--
-- Name: oqcs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.oqcs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.oqcs_id_seq OWNER TO postgres;

--
-- Name: oqcs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.oqcs_id_seq OWNED BY public.oqcs.id;


--
-- Name: pro_dispatch; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pro_dispatch (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    order_id bigint NOT NULL,
    order_no character varying(50),
    process_id bigint,
    process_name character varying(100),
    station_id bigint,
    station_name character varying(100),
    assign_user_id bigint,
    assign_user_name character varying(50),
    quantity numeric(18,4),
    status bigint DEFAULT 1
);


ALTER TABLE public.pro_dispatch OWNER TO postgres;

--
-- Name: pro_dispatch_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.pro_dispatch_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.pro_dispatch_id_seq OWNER TO postgres;

--
-- Name: pro_dispatch_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.pro_dispatch_id_seq OWNED BY public.pro_dispatch.id;


--
-- Name: pro_production_order; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pro_production_order (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    order_no character varying(50) NOT NULL,
    sales_order_no character varying(50),
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    material_spec character varying(100),
    unit character varying(20),
    quantity numeric(18,4),
    completed_qty numeric(18,4) DEFAULT 0,
    rejected_qty numeric(18,4) DEFAULT 0,
    workshop_id bigint,
    workshop_name character varying(100),
    line_id bigint,
    line_name character varying(100),
    route_id bigint,
    bom_id bigint,
    plan_start_date timestamp(6) without time zone,
    plan_end_date timestamp(6) without time zone,
    actual_start_date timestamp(6) without time zone,
    actual_end_date timestamp(6) without time zone,
    priority bigint DEFAULT 1,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.pro_production_order OWNER TO postgres;

--
-- Name: pro_production_order_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.pro_production_order_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.pro_production_order_id_seq OWNER TO postgres;

--
-- Name: pro_production_order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.pro_production_order_id_seq OWNED BY public.pro_production_order.id;


--
-- Name: pro_production_report; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pro_production_report (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    order_id bigint NOT NULL,
    order_no character varying(50),
    process_id bigint,
    process_name character varying(100),
    station_id bigint,
    station_name character varying(100),
    report_user_id bigint,
    report_user_name character varying(50),
    report_date timestamp(6) without time zone,
    quantity numeric(18,4),
    qualified_qty numeric(18,4),
    rejected_qty numeric(18,4) DEFAULT 0,
    work_time bigint,
    remark character varying(500)
);


ALTER TABLE public.pro_production_report OWNER TO postgres;

--
-- Name: pro_production_report_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.pro_production_report_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.pro_production_report_id_seq OWNER TO postgres;

--
-- Name: pro_production_report_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.pro_production_report_id_seq OWNED BY public.pro_production_report.id;


--
-- Name: pro_sales_order; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pro_sales_order (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    order_no character varying(50) NOT NULL,
    customer_id bigint,
    customer_name character varying(100),
    order_date timestamp(6) without time zone,
    delivery_date timestamp(6) without time zone,
    order_type character varying(20),
    priority bigint DEFAULT 1,
    status bigint DEFAULT 1,
    remark character varying(500),
    order_id bigint NOT NULL,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    quantity numeric(18,4),
    unit character varying(20),
    price numeric(18,2),
    amount numeric(18,2),
    shipped_qty numeric(18,4) DEFAULT 0
);


ALTER TABLE public.pro_sales_order OWNER TO postgres;

--
-- Name: pro_sales_order_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.pro_sales_order_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.pro_sales_order_id_seq OWNER TO postgres;

--
-- Name: pro_sales_order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.pro_sales_order_id_seq OWNED BY public.pro_sales_order.id;


--
-- Name: pro_sales_order_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.pro_sales_order_item (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    order_id bigint NOT NULL,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    quantity numeric(18,4),
    unit character varying(20),
    price numeric(18,2),
    amount numeric(18,2),
    shipped_qty numeric(18,4) DEFAULT 0
);


ALTER TABLE public.pro_sales_order_item OWNER TO postgres;

--
-- Name: pro_sales_order_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.pro_sales_order_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.pro_sales_order_item_id_seq OWNER TO postgres;

--
-- Name: pro_sales_order_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.pro_sales_order_item_id_seq OWNED BY public.pro_sales_order_item.id;


--
-- Name: processes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.processes (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    process_code character varying(50) NOT NULL,
    process_name character varying(100) NOT NULL,
    process_type character varying(20),
    sequence bigint DEFAULT 0,
    station_id bigint,
    status bigint DEFAULT 1
);


ALTER TABLE public.processes OWNER TO postgres;

--
-- Name: processes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.processes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.processes_id_seq OWNER TO postgres;

--
-- Name: processes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.processes_id_seq OWNED BY public.processes.id;


--
-- Name: production_daily_report; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.production_daily_report (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    report_date timestamp with time zone NOT NULL,
    workshop_id bigint,
    workshop_name character varying(100),
    production_order_count bigint,
    completed_order_count bigint,
    total_output_qty numeric,
    qualified_qty numeric,
    defect_qty numeric,
    pass_rate numeric,
    first_pass_rate numeric,
    oee numeric,
    output_per_hour numeric,
    worker_count bigint,
    working_hours numeric,
    remark character varying(500)
);


ALTER TABLE public.production_daily_report OWNER TO postgres;

--
-- Name: production_daily_report_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.production_daily_report_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.production_daily_report_id_seq OWNER TO postgres;

--
-- Name: production_daily_report_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.production_daily_report_id_seq OWNED BY public.production_daily_report.id;


--
-- Name: production_lines; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.production_lines (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    line_code character varying(50) NOT NULL,
    line_name character varying(100) NOT NULL,
    workshop_id bigint NOT NULL,
    line_type character varying(20),
    status bigint DEFAULT 1
);


ALTER TABLE public.production_lines OWNER TO postgres;

--
-- Name: production_lines_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.production_lines_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.production_lines_id_seq OWNER TO postgres;

--
-- Name: production_lines_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.production_lines_id_seq OWNED BY public.production_lines.id;


--
-- Name: production_offline; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.production_offline (
    id bigint NOT NULL,
    tenant_id bigint DEFAULT 1 NOT NULL,
    offline_code character varying(50),
    work_order_id bigint,
    work_order_code character varying(50),
    product_id bigint,
    product_code character varying(50),
    product_name character varying(100),
    offline_type character varying(20),
    offline_reason character varying(200),
    offline_qty numeric,
    process_route_id bigint,
    current_op_id bigint,
    current_op_name character varying(50),
    handle_method character varying(20),
    handle_qty numeric,
    handle_result character varying(20),
    rework_order_id bigint,
    scrap_qty numeric,
    downgrade_qty numeric,
    status character varying(20),
    operator_id bigint,
    operator_name character varying(50),
    remark character varying(500),
    created_by character varying(50),
    updated_by character varying(50),
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.production_offline OWNER TO postgres;

--
-- Name: production_offline_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.production_offline_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.production_offline_id_seq OWNER TO postgres;

--
-- Name: production_offline_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.production_offline_id_seq OWNED BY public.production_offline.id;


--
-- Name: production_offline_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.production_offline_item (
    id bigint NOT NULL,
    tenant_id bigint DEFAULT 1 NOT NULL,
    offline_id bigint NOT NULL,
    serial_no character varying(50),
    batch_no character varying(50),
    offline_qty numeric,
    handle_method character varying(20),
    handle_qty numeric,
    handle_result character varying(20),
    remark character varying(200),
    created_at timestamp with time zone
);


ALTER TABLE public.production_offline_item OWNER TO postgres;

--
-- Name: production_offline_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.production_offline_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.production_offline_item_id_seq OWNER TO postgres;

--
-- Name: production_offline_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.production_offline_item_id_seq OWNED BY public.production_offline_item.id;


--
-- Name: production_orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.production_orders (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    order_no character varying(50) NOT NULL,
    sales_order_no character varying(50),
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    material_spec character varying(100),
    unit character varying(20),
    quantity numeric(18,4),
    completed_qty numeric(18,4) DEFAULT 0,
    rejected_qty numeric(18,4) DEFAULT 0,
    workshop_id bigint,
    workshop_name character varying(100),
    line_id bigint,
    line_name character varying(100),
    route_id bigint,
    bom_id bigint,
    plan_start_date timestamp(6) with time zone,
    plan_end_date timestamp(6) with time zone,
    actual_start_date timestamp(6) with time zone,
    actual_end_date timestamp(6) with time zone,
    priority bigint DEFAULT 1,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.production_orders OWNER TO postgres;

--
-- Name: production_orders_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.production_orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.production_orders_id_seq OWNER TO postgres;

--
-- Name: production_orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.production_orders_id_seq OWNED BY public.production_orders.id;


--
-- Name: production_reports; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.production_reports (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    order_id bigint NOT NULL,
    order_no character varying(50),
    process_id bigint,
    process_name character varying(100),
    station_id bigint,
    station_name character varying(100),
    report_user_id bigint,
    report_user_name character varying(50),
    report_date timestamp(6) with time zone,
    quantity numeric(18,4),
    qualified_qty numeric(18,4),
    rejected_qty numeric(18,4) DEFAULT 0,
    work_time bigint,
    remark character varying(500)
);


ALTER TABLE public.production_reports OWNER TO postgres;

--
-- Name: production_reports_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.production_reports_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.production_reports_id_seq OWNER TO postgres;

--
-- Name: production_reports_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.production_reports_id_seq OWNED BY public.production_reports.id;


--
-- Name: qc_aql_level; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.qc_aql_level (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    level character varying(10) NOT NULL,
    name character varying(50) NOT NULL,
    type character varying(20),
    "order" bigint DEFAULT 0,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.qc_aql_level OWNER TO postgres;

--
-- Name: qc_aql_level_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.qc_aql_level_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.qc_aql_level_id_seq OWNER TO postgres;

--
-- Name: qc_aql_level_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.qc_aql_level_id_seq OWNED BY public.qc_aql_level.id;


--
-- Name: qc_aql_table_row; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.qc_aql_table_row (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    aql_level_id bigint NOT NULL,
    aql_value character varying(10),
    batch_min bigint,
    batch_max bigint,
    sample_size bigint,
    ac bigint,
    re bigint
);


ALTER TABLE public.qc_aql_table_row OWNER TO postgres;

--
-- Name: qc_aql_table_row_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.qc_aql_table_row_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.qc_aql_table_row_id_seq OWNER TO postgres;

--
-- Name: qc_aql_table_row_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.qc_aql_table_row_id_seq OWNED BY public.qc_aql_table_row.id;


--
-- Name: qc_defect_code; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.qc_defect_code (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    defect_code character varying(20) NOT NULL,
    defect_name character varying(100) NOT NULL,
    defect_type character varying(20),
    severity bigint DEFAULT 1,
    status bigint DEFAULT 1
);


ALTER TABLE public.qc_defect_code OWNER TO postgres;

--
-- Name: qc_defect_code_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.qc_defect_code_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.qc_defect_code_id_seq OWNER TO postgres;

--
-- Name: qc_defect_code_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.qc_defect_code_id_seq OWNED BY public.qc_defect_code.id;


--
-- Name: qc_defect_record; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.qc_defect_record (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    record_no character varying(50) NOT NULL,
    order_id bigint,
    order_no character varying(50),
    process_id bigint,
    process_name character varying(100),
    defect_code_id bigint,
    defect_code character varying(20),
    defect_name character varying(100),
    quantity numeric(18,4),
    handle_method bigint DEFAULT 1,
    handle_user_id bigint,
    handle_date timestamp(6) without time zone,
    status bigint DEFAULT 1
);


ALTER TABLE public.qc_defect_record OWNER TO postgres;

--
-- Name: qc_defect_record_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.qc_defect_record_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.qc_defect_record_id_seq OWNER TO postgres;

--
-- Name: qc_defect_record_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.qc_defect_record_id_seq OWNED BY public.qc_defect_record.id;


--
-- Name: qc_fqc; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.qc_fqc (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    fqc_no character varying(50) NOT NULL,
    order_id bigint,
    order_no character varying(50),
    quantity numeric(18,4),
    sample_size bigint,
    qualified_qty numeric(18,4),
    rejected_qty numeric(18,4) DEFAULT 0,
    check_user_id bigint,
    check_user_name character varying(50),
    check_date timestamp(6) without time zone,
    result bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.qc_fqc OWNER TO postgres;

--
-- Name: qc_fqc_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.qc_fqc_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.qc_fqc_id_seq OWNER TO postgres;

--
-- Name: qc_fqc_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.qc_fqc_id_seq OWNED BY public.qc_fqc.id;


--
-- Name: qc_inspection_characteristic; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.qc_inspection_characteristic (
    id bigint NOT NULL,
    code character varying(50) NOT NULL,
    name character varying(100) NOT NULL,
    type character varying(20),
    spec_lower numeric(18,4),
    spec_upper numeric(18,4),
    usl numeric(18,4),
    lsl numeric(18,4),
    target numeric(18,4),
    unit character varying(20),
    aql numeric(5,2),
    inspection_method character varying(100),
    tenant_id bigint,
    status bigint DEFAULT 1
);


ALTER TABLE public.qc_inspection_characteristic OWNER TO postgres;

--
-- Name: qc_inspection_characteristic_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.qc_inspection_characteristic_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.qc_inspection_characteristic_id_seq OWNER TO postgres;

--
-- Name: qc_inspection_characteristic_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.qc_inspection_characteristic_id_seq OWNED BY public.qc_inspection_characteristic.id;


--
-- Name: qc_inspection_feature; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.qc_inspection_feature (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    tenant_id bigint DEFAULT 1 NOT NULL,
    feature_code character varying(50),
    feature_name character varying(100) NOT NULL,
    product_id bigint,
    product_code character varying(50),
    product_name character varying(100),
    inspection_type character varying(20),
    feature_type character varying(20),
    technical_spec character varying(500),
    lower_limit numeric,
    upper_limit numeric,
    unit character varying(20),
    sample_size bigint,
    gauges_method character varying(100),
    aql_level character varying(20),
    status character varying(20),
    remark character varying(500),
    created_by character varying(50),
    updated_by character varying(50)
);


ALTER TABLE public.qc_inspection_feature OWNER TO postgres;

--
-- Name: qc_inspection_feature_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.qc_inspection_feature_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.qc_inspection_feature_id_seq OWNER TO postgres;

--
-- Name: qc_inspection_feature_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.qc_inspection_feature_id_seq OWNED BY public.qc_inspection_feature.id;


--
-- Name: qc_ipqc; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.qc_ipqc (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    ipqc_no character varying(50) NOT NULL,
    order_id bigint,
    order_no character varying(50),
    process_id bigint,
    process_name character varying(100),
    quantity numeric(18,4),
    sample_size bigint,
    check_user_id bigint,
    check_user_name character varying(50),
    check_date timestamp(6) without time zone,
    result bigint DEFAULT 1,
    remark character varying(500),
    ip_qc_no character varying(50) NOT NULL
);


ALTER TABLE public.qc_ipqc OWNER TO postgres;

--
-- Name: qc_ipqc_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.qc_ipqc_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.qc_ipqc_id_seq OWNER TO postgres;

--
-- Name: qc_ipqc_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.qc_ipqc_id_seq OWNED BY public.qc_ipqc.id;


--
-- Name: qc_iqc; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.qc_iqc (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    iqc_no character varying(50) NOT NULL,
    supplier_id bigint,
    supplier_name character varying(100),
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    quantity numeric(18,4),
    unit character varying(20),
    check_user_id bigint,
    check_user_name character varying(50),
    check_date timestamp(6) without time zone,
    result bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.qc_iqc OWNER TO postgres;

--
-- Name: qc_iqc_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.qc_iqc_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.qc_iqc_id_seq OWNER TO postgres;

--
-- Name: qc_iqc_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.qc_iqc_id_seq OWNED BY public.qc_iqc.id;


--
-- Name: qc_iqc_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.qc_iqc_item (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    iqc_id bigint,
    check_item_id bigint,
    check_item character varying(100),
    check_standard character varying(200),
    check_method character varying(100),
    result bigint DEFAULT 1,
    remark character varying(200)
);


ALTER TABLE public.qc_iqc_item OWNER TO postgres;

--
-- Name: qc_iqc_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.qc_iqc_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.qc_iqc_item_id_seq OWNER TO postgres;

--
-- Name: qc_iqc_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.qc_iqc_item_id_seq OWNED BY public.qc_iqc_item.id;


--
-- Name: qc_ncr; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.qc_ncr (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    ncr_no character varying(50) NOT NULL,
    defect_id bigint,
    source_type character varying(20),
    issue_desc text,
    root_cause text,
    corrective_action text,
    preventive_action text,
    verify_result character varying(200),
    verify_user_id bigint,
    verify_date timestamp(6) without time zone,
    status bigint DEFAULT 1
);


ALTER TABLE public.qc_ncr OWNER TO postgres;

--
-- Name: qc_ncr_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.qc_ncr_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.qc_ncr_id_seq OWNER TO postgres;

--
-- Name: qc_ncr_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.qc_ncr_id_seq OWNED BY public.qc_ncr.id;


--
-- Name: qc_oqc; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.qc_oqc (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    oqc_no character varying(50) NOT NULL,
    shipping_no character varying(50),
    customer_id bigint,
    customer_name character varying(100),
    quantity numeric(18,4),
    check_user_id bigint,
    check_user_name character varying(50),
    check_date timestamp(6) without time zone,
    result bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.qc_oqc OWNER TO postgres;

--
-- Name: qc_oqc_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.qc_oqc_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.qc_oqc_id_seq OWNER TO postgres;

--
-- Name: qc_oqc_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.qc_oqc_id_seq OWNED BY public.qc_oqc.id;


--
-- Name: qc_sampling_plan; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.qc_sampling_plan (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    code character varying(50) NOT NULL,
    name character varying(100) NOT NULL,
    inspection_type character varying(20),
    aql_level_id bigint,
    default_aql numeric(5,2),
    min_batch_size bigint DEFAULT 0,
    max_batch_size bigint DEFAULT 0,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.qc_sampling_plan OWNER TO postgres;

--
-- Name: qc_sampling_plan_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.qc_sampling_plan_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.qc_sampling_plan_id_seq OWNER TO postgres;

--
-- Name: qc_sampling_plan_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.qc_sampling_plan_id_seq OWNED BY public.qc_sampling_plan.id;


--
-- Name: qc_spc_data; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.qc_spc_data (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    equipment_id bigint,
    station_id bigint,
    process_id bigint,
    process_name character varying(100),
    check_item character varying(100),
    check_value numeric(18,4),
    usl numeric(18,4),
    lsl numeric(18,4),
    cl numeric(18,4),
    ucl numeric(18,4),
    lcl numeric(18,4),
    check_time timestamp(6) without time zone NOT NULL
);


ALTER TABLE public.qc_spc_data OWNER TO postgres;

--
-- Name: qc_spc_data_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.qc_spc_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.qc_spc_data_id_seq OWNER TO postgres;

--
-- Name: qc_spc_data_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.qc_spc_data_id_seq OWNED BY public.qc_spc_data.id;


--
-- Name: qms_sampling_plan; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.qms_sampling_plan (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    plan_code character varying(50) NOT NULL,
    plan_name character varying(100) NOT NULL,
    inspection_level character varying(20),
    sample_type character varying(20),
    aql numeric(5,2),
    status character varying(20) DEFAULT 'ACTIVE'::character varying,
    remark text
);


ALTER TABLE public.qms_sampling_plan OWNER TO postgres;

--
-- Name: qms_sampling_plan_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.qms_sampling_plan_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.qms_sampling_plan_id_seq OWNER TO postgres;

--
-- Name: qms_sampling_plan_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.qms_sampling_plan_id_seq OWNED BY public.qms_sampling_plan.id;


--
-- Name: qms_sampling_record; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.qms_sampling_record (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    plan_id bigint NOT NULL,
    plan_code character varying(50),
    inspection_id bigint NOT NULL,
    batch_qty numeric(18,3),
    sample_size bigint,
    defect_count bigint,
    ac_result boolean DEFAULT false,
    inspector character varying(50),
    inspect_time timestamp with time zone
);


ALTER TABLE public.qms_sampling_record OWNER TO postgres;

--
-- Name: qms_sampling_record_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.qms_sampling_record_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.qms_sampling_record_id_seq OWNER TO postgres;

--
-- Name: qms_sampling_record_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.qms_sampling_record_id_seq OWNED BY public.qms_sampling_record.id;


--
-- Name: qms_sampling_rule; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.qms_sampling_rule (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    plan_id bigint NOT NULL,
    batch_qty_from numeric(18,3),
    batch_qty_to numeric(18,3),
    sample_size bigint,
    ac_accept bigint,
    re_reject bigint
);


ALTER TABLE public.qms_sampling_rule OWNER TO postgres;

--
-- Name: qms_sampling_rule_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.qms_sampling_rule_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.qms_sampling_rule_id_seq OWNER TO postgres;

--
-- Name: qms_sampling_rule_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.qms_sampling_rule_id_seq OWNED BY public.qms_sampling_rule.id;


--
-- Name: quality_weekly_report; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.quality_weekly_report (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    report_year bigint,
    report_week bigint,
    start_date timestamp with time zone,
    end_date timestamp with time zone,
    workshop_id bigint,
    workshop_name character varying(100),
    total_inspection_qty bigint,
    qualified_qty bigint,
    defect_qty bigint,
    pass_rate numeric,
    iqc_insp_qty bigint,
    iqc_qualified_qty bigint,
    iqc_defect_qty bigint,
    ip_qc_insp_qty bigint,
    ip_qc_qualified_qty bigint,
    ip_qc_defect_qty bigint,
    fqc_insp_qty bigint,
    fqc_qualified_qty bigint,
    fqc_defect_qty bigint,
    oqc_insp_qty bigint,
    oqc_qualified_qty bigint,
    oqc_defect_qty bigint,
    ncr_count bigint,
    customer_complaint_count bigint,
    remark character varying(500)
);


ALTER TABLE public.quality_weekly_report OWNER TO postgres;

--
-- Name: quality_weekly_report_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.quality_weekly_report_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.quality_weekly_report_id_seq OWNER TO postgres;

--
-- Name: quality_weekly_report_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.quality_weekly_report_id_seq OWNED BY public.quality_weekly_report.id;


--
-- Name: receive_order_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.receive_order_items (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    receive_id bigint,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    quantity numeric(18,4),
    received_qty numeric(18,4) DEFAULT 0,
    unit character varying(20),
    batch_no character varying(50)
);


ALTER TABLE public.receive_order_items OWNER TO postgres;

--
-- Name: receive_order_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.receive_order_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.receive_order_items_id_seq OWNER TO postgres;

--
-- Name: receive_order_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.receive_order_items_id_seq OWNED BY public.receive_order_items.id;


--
-- Name: receive_orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.receive_orders (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    receive_no character varying(50) NOT NULL,
    supplier_id bigint,
    supplier_name character varying(100),
    warehouse_id bigint,
    receive_date timestamp(6) with time zone,
    receive_user_id bigint,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.receive_orders OWNER TO postgres;

--
-- Name: receive_orders_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.receive_orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.receive_orders_id_seq OWNER TO postgres;

--
-- Name: receive_orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.receive_orders_id_seq OWNED BY public.receive_orders.id;


--
-- Name: resources; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.resources (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    resource_code character varying(50) NOT NULL,
    resource_name character varying(100) NOT NULL,
    resource_type character varying(20),
    workshop_id bigint,
    capacity numeric(18,2),
    unit character varying(20),
    efficiency numeric DEFAULT 100,
    status bigint DEFAULT 1
);


ALTER TABLE public.resources OWNER TO postgres;

--
-- Name: resources_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.resources_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.resources_id_seq OWNER TO postgres;

--
-- Name: resources_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.resources_id_seq OWNED BY public.resources.id;


--
-- Name: route_operations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.route_operations (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    route_id bigint NOT NULL,
    process_id bigint NOT NULL,
    process_name character varying(100),
    sequence bigint DEFAULT 0,
    station_id bigint,
    station_name character varying(100),
    standard_time bigint,
    queue_time bigint,
    transport_time bigint
);


ALTER TABLE public.route_operations OWNER TO postgres;

--
-- Name: route_operations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.route_operations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.route_operations_id_seq OWNER TO postgres;

--
-- Name: route_operations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.route_operations_id_seq OWNED BY public.route_operations.id;


--
-- Name: routes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.routes (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    route_code character varying(50) NOT NULL,
    route_name character varying(100) NOT NULL,
    material_id bigint NOT NULL,
    version character varying(20),
    status bigint DEFAULT 1
);


ALTER TABLE public.routes OWNER TO postgres;

--
-- Name: routes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.routes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.routes_id_seq OWNER TO postgres;

--
-- Name: routes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.routes_id_seq OWNED BY public.routes.id;


--
-- Name: sales_order_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sales_order_items (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    order_id bigint NOT NULL,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    quantity numeric(18,4),
    unit character varying(20),
    price numeric(18,2),
    amount numeric(18,2),
    shipped_qty numeric(18,4) DEFAULT 0
);


ALTER TABLE public.sales_order_items OWNER TO postgres;

--
-- Name: sales_order_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sales_order_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sales_order_items_id_seq OWNER TO postgres;

--
-- Name: sales_order_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sales_order_items_id_seq OWNED BY public.sales_order_items.id;


--
-- Name: sales_orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sales_orders (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    order_no character varying(50) NOT NULL,
    customer_id bigint,
    customer_name character varying(100),
    order_date timestamp(6) with time zone,
    delivery_date timestamp(6) with time zone,
    order_type character varying(20),
    priority bigint DEFAULT 1,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.sales_orders OWNER TO postgres;

--
-- Name: sales_orders_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sales_orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sales_orders_id_seq OWNER TO postgres;

--
-- Name: sales_orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sales_orders_id_seq OWNED BY public.sales_orders.id;


--
-- Name: schedule_plans; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schedule_plans (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    plan_no character varying(50) NOT NULL,
    plan_type character varying(20),
    start_date timestamp(6) with time zone,
    end_date timestamp(6) with time zone,
    algorithm character varying(20),
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.schedule_plans OWNER TO postgres;

--
-- Name: schedule_plans_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.schedule_plans_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.schedule_plans_id_seq OWNER TO postgres;

--
-- Name: schedule_plans_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.schedule_plans_id_seq OWNED BY public.schedule_plans.id;


--
-- Name: schedule_results; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schedule_results (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    plan_id bigint,
    order_id bigint,
    order_no character varying(50),
    sequence bigint,
    line_id bigint,
    line_name character varying(100),
    station_id bigint,
    station_name character varying(100),
    plan_start_time timestamp(6) with time zone,
    plan_end_time timestamp(6) with time zone,
    actual_start_time timestamp(6) with time zone,
    actual_end_time timestamp(6) with time zone,
    status bigint DEFAULT 1
);


ALTER TABLE public.schedule_results OWNER TO postgres;

--
-- Name: schedule_results_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.schedule_results_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.schedule_results_id_seq OWNER TO postgres;

--
-- Name: schedule_results_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.schedule_results_id_seq OWNED BY public.schedule_results.id;


--
-- Name: scp_customer_inquiry; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_customer_inquiry (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    inquiry_no character varying(50) NOT NULL,
    inquiry_date date NOT NULL,
    customer_id bigint,
    customer_name character varying(100),
    contact_person character varying(50),
    contact_phone character varying(20),
    contact_email character varying(100),
    expected_date date,
    valid_until date,
    currency character varying(10) DEFAULT 'CNY'::character varying,
    status character varying(20) DEFAULT 'DRAFT'::character varying,
    quoted_amount numeric(18,2),
    winner_supplier_id bigint,
    remark text,
    created_by character varying(50)
);


ALTER TABLE public.scp_customer_inquiry OWNER TO postgres;

--
-- Name: scp_customer_inquiry_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_customer_inquiry_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_customer_inquiry_id_seq OWNER TO postgres;

--
-- Name: scp_customer_inquiry_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_customer_inquiry_id_seq OWNED BY public.scp_customer_inquiry.id;


--
-- Name: scp_inquiry_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_inquiry_item (
    id bigint NOT NULL,
    inquiry_id bigint NOT NULL,
    line_no bigint NOT NULL,
    material_id bigint,
    material_code character varying(50) NOT NULL,
    material_name character varying(100),
    specification character varying(200),
    unit character varying(20) DEFAULT 'PCS'::character varying,
    required_qty numeric(18,3),
    target_price numeric(18,4),
    quoted_price numeric(18,4),
    quoted_supplier_id bigint,
    lead_time_days bigint,
    remark text,
    created_at timestamp with time zone
);


ALTER TABLE public.scp_inquiry_item OWNER TO postgres;

--
-- Name: scp_inquiry_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_inquiry_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_inquiry_item_id_seq OWNER TO postgres;

--
-- Name: scp_inquiry_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_inquiry_item_id_seq OWNED BY public.scp_inquiry_item.id;


--
-- Name: scp_mrs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_mrs (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    mrs_no character varying(50) NOT NULL,
    plan_month character varying(7),
    status character varying(20) DEFAULT 'DRAFT'::character varying,
    source_type character varying(20),
    source_no character varying(50),
    total_items bigint DEFAULT 0,
    total_qty numeric(18,3),
    published_by bigint,
    published_at timestamp with time zone,
    remark text
);


ALTER TABLE public.scp_mrs OWNER TO postgres;

--
-- Name: scp_mrs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_mrs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_mrs_id_seq OWNER TO postgres;

--
-- Name: scp_mrs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_mrs_id_seq OWNED BY public.scp_mrs.id;


--
-- Name: scp_mrs_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_mrs_item (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    mrs_id bigint NOT NULL,
    mrs_no character varying(50) NOT NULL,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    spec character varying(200),
    unit character varying(20),
    req_qty numeric(18,3),
    on_hand_qty numeric(18,3),
    short_qty numeric(18,3),
    supplier_id bigint,
    supplier_name character varying(100),
    promise_date date,
    status character varying(20) DEFAULT 'PENDING'::character varying
);


ALTER TABLE public.scp_mrs_item OWNER TO postgres;

--
-- Name: scp_mrs_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_mrs_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_mrs_item_id_seq OWNER TO postgres;

--
-- Name: scp_mrs_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_mrs_item_id_seq OWNED BY public.scp_mrs_item.id;


--
-- Name: scp_po_change_log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_po_change_log (
    id bigint NOT NULL,
    po_id bigint NOT NULL,
    change_type character varying(30) NOT NULL,
    change_field character varying(50),
    old_value character varying(200),
    new_value character varying(200),
    changed_by bigint,
    changed_by_name character varying(50),
    change_time timestamp with time zone,
    reason character varying(200),
    tenant_id bigint NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.scp_po_change_log OWNER TO postgres;

--
-- Name: scp_po_change_log_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_po_change_log_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_po_change_log_id_seq OWNER TO postgres;

--
-- Name: scp_po_change_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_po_change_log_id_seq OWNED BY public.scp_po_change_log.id;


--
-- Name: scp_purchase_order; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_purchase_order (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    po_no character varying(50) NOT NULL,
    po_type character varying(20) DEFAULT 'STANDARD'::character varying,
    supplier_id bigint NOT NULL,
    supplier_code character varying(50),
    supplier_name character varying(100),
    contact_person character varying(50),
    contact_phone character varying(20),
    contact_email character varying(100),
    order_date date NOT NULL,
    promised_date date,
    currency character varying(10) DEFAULT 'CNY'::character varying,
    payment_terms character varying(50),
    tax_rate numeric(5,2) DEFAULT 13,
    total_amount numeric(18,2) DEFAULT 0,
    total_qty numeric(18,3) DEFAULT 0,
    approved_by bigint,
    approved_time timestamp with time zone,
    approval_status character varying(20) DEFAULT 'PENDING'::character varying,
    status character varying(20) DEFAULT 'DRAFT'::character varying,
    source_type character varying(20),
    source_no character varying(50),
    remark text,
    close_reason character varying(500)
);


ALTER TABLE public.scp_purchase_order OWNER TO postgres;

--
-- Name: scp_purchase_order_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_purchase_order_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_purchase_order_id_seq OWNER TO postgres;

--
-- Name: scp_purchase_order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_purchase_order_id_seq OWNED BY public.scp_purchase_order.id;


--
-- Name: scp_purchase_order_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_purchase_order_item (
    id bigint NOT NULL,
    po_id bigint NOT NULL,
    line_no bigint NOT NULL,
    material_id bigint,
    material_code character varying(50) NOT NULL,
    material_name character varying(100),
    specification character varying(200),
    unit character varying(20) DEFAULT 'PCS'::character varying,
    unit_price numeric(18,4) DEFAULT 0,
    order_qty numeric(18,3) NOT NULL,
    delivered_qty numeric(18,3) DEFAULT 0,
    received_qty numeric(18,3) DEFAULT 0,
    tax_amount numeric(18,2) DEFAULT 0,
    line_amount numeric(18,2) DEFAULT 0,
    promised_date date,
    actual_delivery_date date,
    batch_no character varying(50),
    quality_require text,
    package_require character varying(200),
    is_gifted bigint DEFAULT 0,
    status character varying(20) DEFAULT 'PENDING'::character varying,
    remark text,
    created_at timestamp with time zone
);


ALTER TABLE public.scp_purchase_order_item OWNER TO postgres;

--
-- Name: scp_purchase_order_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_purchase_order_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_purchase_order_item_id_seq OWNER TO postgres;

--
-- Name: scp_purchase_order_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_purchase_order_item_id_seq OWNED BY public.scp_purchase_order_item.id;


--
-- Name: scp_purchase_plan; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_purchase_plan (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    plan_no character varying(50) NOT NULL,
    title character varying(200),
    plan_type character varying(20),
    plan_year bigint,
    plan_month bigint,
    quarter bigint,
    status character varying(20) DEFAULT 'DRAFT'::character varying,
    total_items bigint DEFAULT 0,
    total_amount numeric(18,2),
    currency character varying(10) DEFAULT 'CNY'::character varying,
    department character varying(100),
    submitter_id bigint,
    submitted_at timestamp with time zone,
    confirmed_by bigint,
    confirmed_at timestamp with time zone,
    published_by bigint,
    published_at timestamp with time zone,
    closed_by bigint,
    closed_at timestamp with time zone,
    close_reason character varying(500),
    remark text
);


ALTER TABLE public.scp_purchase_plan OWNER TO postgres;

--
-- Name: scp_purchase_plan_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_purchase_plan_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_purchase_plan_id_seq OWNER TO postgres;

--
-- Name: scp_purchase_plan_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_purchase_plan_id_seq OWNED BY public.scp_purchase_plan.id;


--
-- Name: scp_purchase_plan_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_purchase_plan_item (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    plan_id bigint NOT NULL,
    plan_no character varying(50) NOT NULL,
    line_no bigint,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    spec character varying(200),
    unit character varying(20),
    req_qty numeric(18,3),
    unit_price numeric(18,6),
    line_amount numeric(18,2),
    supplier_id bigint,
    supplier_code character varying(50),
    supplier_name character varying(100),
    req_delivery_date date,
    promise_date date,
    mrs_no character varying(50),
    mrs_line_no bigint,
    status character varying(20) DEFAULT 'PENDING'::character varying,
    remark text
);


ALTER TABLE public.scp_purchase_plan_item OWNER TO postgres;

--
-- Name: scp_purchase_plan_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_purchase_plan_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_purchase_plan_item_id_seq OWNER TO postgres;

--
-- Name: scp_purchase_plan_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_purchase_plan_item_id_seq OWNED BY public.scp_purchase_plan_item.id;


--
-- Name: scp_qad_sync_log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_qad_sync_log (
    id bigint NOT NULL,
    tenant_id bigint,
    sync_type character varying(20),
    sync_direction character varying(10),
    qad_doc_no character varying(50),
    mom_doc_no character varying(50),
    status character varying(20),
    request_content text,
    response_content text,
    error_msg character varying(500),
    sync_time timestamp with time zone,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.scp_qad_sync_log OWNER TO postgres;

--
-- Name: scp_qad_sync_log_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_qad_sync_log_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_qad_sync_log_id_seq OWNER TO postgres;

--
-- Name: scp_qad_sync_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_qad_sync_log_id_seq OWNED BY public.scp_qad_sync_log.id;


--
-- Name: scp_quote_comparison; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_quote_comparison (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    rfq_id bigint NOT NULL,
    comparison_no character varying(50) NOT NULL,
    compared_by bigint,
    compared_at timestamp with time zone,
    summary_data jsonb,
    recommendation character varying(100),
    decision_remark text
);


ALTER TABLE public.scp_quote_comparison OWNER TO postgres;

--
-- Name: scp_quote_comparison_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_quote_comparison_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_quote_comparison_id_seq OWNER TO postgres;

--
-- Name: scp_quote_comparison_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_quote_comparison_id_seq OWNED BY public.scp_quote_comparison.id;


--
-- Name: scp_quote_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_quote_item (
    id bigint NOT NULL,
    quote_id bigint NOT NULL,
    rfq_line_id bigint,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    unit character varying(20),
    quoted_qty numeric(18,3),
    unit_price numeric(18,4) NOT NULL,
    line_amount numeric(18,2) DEFAULT 0,
    delivery_date date,
    lead_time_days bigint,
    remark text,
    created_at timestamp with time zone
);


ALTER TABLE public.scp_quote_item OWNER TO postgres;

--
-- Name: scp_quote_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_quote_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_quote_item_id_seq OWNER TO postgres;

--
-- Name: scp_quote_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_quote_item_id_seq OWNED BY public.scp_quote_item.id;


--
-- Name: scp_rfq; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_rfq (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    rfq_no character varying(50) NOT NULL,
    rfq_name character varying(200),
    rfq_type character varying(20) DEFAULT 'STANDARD'::character varying,
    inquiry_date date NOT NULL,
    deadline_date date NOT NULL,
    currency character varying(10) DEFAULT 'CNY'::character varying,
    payment_terms character varying(50),
    delivery_terms character varying(100),
    quality_standard character varying(100),
    status character varying(20) DEFAULT 'DRAFT'::character varying,
    awarded_supplier_id bigint,
    awarded_total_amount numeric(18,2),
    total_bids bigint DEFAULT 0,
    is_evaluated bigint DEFAULT 0,
    evaluation_by bigint,
    evaluation_time timestamp with time zone,
    remark text,
    created_by character varying(50)
);


ALTER TABLE public.scp_rfq OWNER TO postgres;

--
-- Name: scp_rfq_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_rfq_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_rfq_id_seq OWNER TO postgres;

--
-- Name: scp_rfq_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_rfq_id_seq OWNED BY public.scp_rfq.id;


--
-- Name: scp_rfq_invite; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_rfq_invite (
    id bigint NOT NULL,
    rfq_id bigint NOT NULL,
    supplier_id bigint NOT NULL,
    supplier_code character varying(50),
    supplier_name character varying(100),
    contact_person character varying(50),
    contact_email character varying(100),
    invite_date date,
    response_status character varying(20) DEFAULT 'PENDING'::character varying,
    quoted_date date,
    tenant_id bigint NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.scp_rfq_invite OWNER TO postgres;

--
-- Name: scp_rfq_invite_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_rfq_invite_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_rfq_invite_id_seq OWNER TO postgres;

--
-- Name: scp_rfq_invite_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_rfq_invite_id_seq OWNED BY public.scp_rfq_invite.id;


--
-- Name: scp_rfq_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_rfq_item (
    id bigint NOT NULL,
    rfq_id bigint NOT NULL,
    line_no bigint NOT NULL,
    material_id bigint,
    material_code character varying(50) NOT NULL,
    material_name character varying(100),
    specification character varying(200),
    unit character varying(20) DEFAULT 'PCS'::character varying,
    required_qty numeric(18,3),
    target_price numeric(18,4),
    market_price numeric(18,4),
    requested_date date,
    quality_require text,
    remark text,
    created_at timestamp with time zone
);


ALTER TABLE public.scp_rfq_item OWNER TO postgres;

--
-- Name: scp_rfq_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_rfq_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_rfq_item_id_seq OWNER TO postgres;

--
-- Name: scp_rfq_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_rfq_item_id_seq OWNED BY public.scp_rfq_item.id;


--
-- Name: scp_sales_order; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_sales_order (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    so_no character varying(50) NOT NULL,
    so_type character varying(20) DEFAULT 'STANDARD'::character varying,
    customer_id bigint NOT NULL,
    customer_code character varying(50),
    customer_name character varying(100),
    contact_person character varying(50),
    contact_phone character varying(20),
    contact_email character varying(100),
    sales_person_id bigint,
    sales_person_name character varying(50),
    order_date date NOT NULL,
    promised_date date,
    currency character varying(10) DEFAULT 'CNY'::character varying,
    payment_terms character varying(50),
    tax_rate numeric(5,2) DEFAULT 13,
    total_amount numeric(18,2) DEFAULT 0,
    total_qty numeric(18,3) DEFAULT 0,
    delivered_amount numeric(18,2) DEFAULT 0,
    delivered_qty numeric(18,3) DEFAULT 0,
    approved_by bigint,
    approved_time timestamp with time zone,
    approval_status character varying(20) DEFAULT 'PENDING'::character varying,
    status character varying(20) DEFAULT 'DRAFT'::character varying,
    source_type character varying(20),
    source_no character varying(50),
    delivery_address text,
    delivery_warehouse_id bigint,
    remark text
);


ALTER TABLE public.scp_sales_order OWNER TO postgres;

--
-- Name: scp_sales_order_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_sales_order_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_sales_order_id_seq OWNER TO postgres;

--
-- Name: scp_sales_order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_sales_order_id_seq OWNED BY public.scp_sales_order.id;


--
-- Name: scp_sales_order_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_sales_order_item (
    id bigint NOT NULL,
    so_id bigint NOT NULL,
    line_no bigint NOT NULL,
    material_id bigint,
    material_code character varying(50) NOT NULL,
    material_name character varying(100),
    specification character varying(200),
    unit character varying(20) DEFAULT 'PCS'::character varying,
    unit_price numeric(18,4) DEFAULT 0,
    order_qty numeric(18,3) NOT NULL,
    delivered_qty numeric(18,3) DEFAULT 0,
    shipped_qty numeric(18,3) DEFAULT 0,
    tax_amount numeric(18,2) DEFAULT 0,
    line_amount numeric(18,2) DEFAULT 0,
    promised_date date,
    actual_delivery_date date,
    production_order_id bigint,
    status character varying(20) DEFAULT 'PENDING'::character varying,
    remark text,
    created_at timestamp with time zone
);


ALTER TABLE public.scp_sales_order_item OWNER TO postgres;

--
-- Name: scp_sales_order_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_sales_order_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_sales_order_item_id_seq OWNER TO postgres;

--
-- Name: scp_sales_order_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_sales_order_item_id_seq OWNED BY public.scp_sales_order_item.id;


--
-- Name: scp_so_change_log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_so_change_log (
    id bigint NOT NULL,
    so_id bigint NOT NULL,
    change_type character varying(30) NOT NULL,
    change_field character varying(50),
    old_value character varying(200),
    new_value character varying(200),
    changed_by bigint,
    changed_by_name character varying(50),
    change_time timestamp with time zone,
    reason character varying(200),
    tenant_id bigint NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.scp_so_change_log OWNER TO postgres;

--
-- Name: scp_so_change_log_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_so_change_log_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_so_change_log_id_seq OWNER TO postgres;

--
-- Name: scp_so_change_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_so_change_log_id_seq OWNED BY public.scp_so_change_log.id;


--
-- Name: scp_supplier_bank; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_supplier_bank (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    supplier_id bigint NOT NULL,
    supplier_code character varying(50),
    supplier_name character varying(100),
    bank_name character varying(100) NOT NULL,
    bank_code character varying(50),
    branch_name character varying(100),
    branch_code character varying(50),
    account_type character varying(20),
    account_no character varying(50) NOT NULL,
    account_name character varying(100) NOT NULL,
    currency character varying(10) DEFAULT 'CNY'::character varying,
    is_primary boolean DEFAULT false,
    is_active boolean DEFAULT true,
    status character varying(20) DEFAULT 'ACTIVE'::character varying,
    remark text
);


ALTER TABLE public.scp_supplier_bank OWNER TO postgres;

--
-- Name: scp_supplier_bank_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_supplier_bank_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_supplier_bank_id_seq OWNER TO postgres;

--
-- Name: scp_supplier_bank_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_supplier_bank_id_seq OWNED BY public.scp_supplier_bank.id;


--
-- Name: scp_supplier_contact; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_supplier_contact (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    supplier_id bigint NOT NULL,
    supplier_code character varying(50),
    supplier_name character varying(100),
    name character varying(50) NOT NULL,
    gender character varying(10),
    department character varying(50),
    "position" character varying(50),
    phone character varying(20),
    mobile character varying(20),
    email character varying(100),
    wechat character varying(50),
    qq character varying(20),
    is_primary boolean DEFAULT false,
    is_active boolean DEFAULT true,
    remark text
);


ALTER TABLE public.scp_supplier_contact OWNER TO postgres;

--
-- Name: scp_supplier_contact_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_supplier_contact_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_supplier_contact_id_seq OWNER TO postgres;

--
-- Name: scp_supplier_contact_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_supplier_contact_id_seq OWNED BY public.scp_supplier_contact.id;


--
-- Name: scp_supplier_delivery_record; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_supplier_delivery_record (
    id bigint NOT NULL,
    supplier_id bigint NOT NULL,
    po_no character varying(50),
    po_line_id bigint,
    promised_date date,
    actual_delivery_date date,
    delivery_status character varying(20),
    delay_days bigint DEFAULT 0,
    is_penalty_applied bigint DEFAULT 0,
    penalty_amount numeric(18,2),
    remark character varying(200),
    tenant_id bigint NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.scp_supplier_delivery_record OWNER TO postgres;

--
-- Name: scp_supplier_delivery_record_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_supplier_delivery_record_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_supplier_delivery_record_id_seq OWNER TO postgres;

--
-- Name: scp_supplier_delivery_record_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_supplier_delivery_record_id_seq OWNED BY public.scp_supplier_delivery_record.id;


--
-- Name: scp_supplier_grade_standard; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_supplier_grade_standard (
    id bigint NOT NULL,
    grade character varying(10) NOT NULL,
    grade_name character varying(50),
    min_score numeric(5,2) NOT NULL,
    max_score numeric(5,2) NOT NULL,
    on_time_rate_threshold numeric(5,2),
    quality_rate_threshold numeric(5,2),
    is_active bigint DEFAULT 1,
    tenant_id bigint NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.scp_supplier_grade_standard OWNER TO postgres;

--
-- Name: scp_supplier_grade_standard_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_supplier_grade_standard_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_supplier_grade_standard_id_seq OWNER TO postgres;

--
-- Name: scp_supplier_grade_standard_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_supplier_grade_standard_id_seq OWNED BY public.scp_supplier_grade_standard.id;


--
-- Name: scp_supplier_kpi; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_supplier_kpi (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    supplier_id bigint NOT NULL,
    supplier_code character varying(50),
    supplier_name character varying(100),
    evaluation_month character varying(7) NOT NULL,
    evaluation_date date NOT NULL,
    evaluated_by bigint,
    evaluated_by_name character varying(50),
    on_time_delivery_rate numeric(5,2),
    total_delivery_orders bigint DEFAULT 0,
    on_time_delivery_count bigint DEFAULT 0,
    avg_delay_days numeric(5,2),
    quality_pass_rate numeric(5,2),
    total_iqc_orders bigint DEFAULT 0,
    passed_iqc_orders bigint DEFAULT 0,
    defect_parts_count bigint DEFAULT 0,
    defect_rate numeric(5,4),
    price_competitiveness numeric(5,2),
    last_purchase_price numeric(18,4),
    market_avg_price numeric(18,4),
    total_score numeric(5,2) DEFAULT 0,
    grade character varying(10),
    rank_position bigint,
    evaluation_remark text
);


ALTER TABLE public.scp_supplier_kpi OWNER TO postgres;

--
-- Name: scp_supplier_kpi_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_supplier_kpi_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_supplier_kpi_id_seq OWNER TO postgres;

--
-- Name: scp_supplier_kpi_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_supplier_kpi_id_seq OWNED BY public.scp_supplier_kpi.id;


--
-- Name: scp_supplier_purchase_info; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_supplier_purchase_info (
    id bigint NOT NULL,
    supplier_id bigint NOT NULL,
    payment_terms character varying(50),
    credit_limit numeric(18,2),
    tax_rate numeric(5,2) DEFAULT 13,
    min_order_amount numeric(18,2),
    lead_time_days bigint,
    supplier_grade character varying(10),
    is_preferred bigint DEFAULT 0,
    is_blacklist bigint DEFAULT 0,
    blacklist_reason character varying(200),
    cooperation_start_date date,
    cooperation_end_date date,
    total_cooperation_amount numeric(18,2),
    tenant_id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.scp_supplier_purchase_info OWNER TO postgres;

--
-- Name: scp_supplier_purchase_info_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_supplier_purchase_info_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_supplier_purchase_info_id_seq OWNER TO postgres;

--
-- Name: scp_supplier_purchase_info_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_supplier_purchase_info_id_seq OWNED BY public.scp_supplier_purchase_info.id;


--
-- Name: scp_supplier_quality_record; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_supplier_quality_record (
    id bigint NOT NULL,
    supplier_id bigint NOT NULL,
    material_code character varying(50),
    iqc_record_id bigint,
    iqc_no character varying(50),
    inspect_date date,
    inspect_qty numeric(18,3),
    qualified_qty numeric(18,3),
    defect_qty numeric(18,3),
    defect_rate numeric(5,4),
    defect_types jsonb,
    is_chargeback bigint DEFAULT 0,
    chargeback_amount numeric(18,2),
    remark character varying(200),
    tenant_id bigint NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.scp_supplier_quality_record OWNER TO postgres;

--
-- Name: scp_supplier_quality_record_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_supplier_quality_record_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_supplier_quality_record_id_seq OWNER TO postgres;

--
-- Name: scp_supplier_quality_record_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_supplier_quality_record_id_seq OWNED BY public.scp_supplier_quality_record.id;


--
-- Name: scp_supplier_quote; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scp_supplier_quote (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    rfq_id bigint NOT NULL,
    rfq_no character varying(50),
    supplier_id bigint NOT NULL,
    supplier_code character varying(50),
    supplier_name character varying(100),
    quote_no character varying(50) NOT NULL,
    quote_date date NOT NULL,
    valid_until date,
    currency character varying(10) DEFAULT 'CNY'::character varying,
    payment_terms character varying(50),
    delivery_days bigint,
    total_amount numeric(18,2) DEFAULT 0,
    is_accepted bigint DEFAULT 0,
    is_lowest bigint DEFAULT 0,
    rank_position bigint,
    quote_status character varying(20) DEFAULT 'SUBMITTED'::character varying,
    evaluation_score numeric(5,2),
    evaluation_result character varying(20),
    evaluator_id bigint,
    evaluation_time timestamp with time zone,
    evaluation_remark text,
    remark text
);


ALTER TABLE public.scp_supplier_quote OWNER TO postgres;

--
-- Name: scp_supplier_quote_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.scp_supplier_quote_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.scp_supplier_quote_id_seq OWNER TO postgres;

--
-- Name: scp_supplier_quote_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.scp_supplier_quote_id_seq OWNED BY public.scp_supplier_quote.id;


--
-- Name: serial_numbers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.serial_numbers (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    serial_number character varying(50) NOT NULL,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    batch_no character varying(50),
    line_id bigint,
    line_name character varying(100),
    order_id bigint,
    order_no character varying(50),
    production_date timestamp(6) with time zone,
    status bigint DEFAULT 1
);


ALTER TABLE public.serial_numbers OWNER TO postgres;

--
-- Name: serial_numbers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.serial_numbers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.serial_numbers_id_seq OWNER TO postgres;

--
-- Name: serial_numbers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.serial_numbers_id_seq OWNED BY public.serial_numbers.id;


--
-- Name: shifts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.shifts (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    shift_code character varying(50) NOT NULL,
    shift_name character varying(100) NOT NULL,
    start_time character varying(10),
    end_time character varying(10),
    break_start character varying(10),
    break_end character varying(10),
    status bigint DEFAULT 1
);


ALTER TABLE public.shifts OWNER TO postgres;

--
-- Name: shifts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.shifts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.shifts_id_seq OWNER TO postgres;

--
-- Name: shifts_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.shifts_id_seq OWNED BY public.shifts.id;


--
-- Name: spc_data; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.spc_data (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    equipment_id bigint,
    station_id bigint,
    process_id bigint,
    process_name character varying(100),
    check_item character varying(100),
    check_value numeric(18,4),
    usl numeric,
    lsl numeric,
    cl numeric,
    ucl numeric,
    lcl numeric,
    check_time timestamp(6) with time zone
);


ALTER TABLE public.spc_data OWNER TO postgres;

--
-- Name: spc_data_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.spc_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.spc_data_id_seq OWNER TO postgres;

--
-- Name: spc_data_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.spc_data_id_seq OWNED BY public.spc_data.id;


--
-- Name: stock_checks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.stock_checks (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    check_no character varying(50) NOT NULL,
    warehouse_id bigint,
    check_date timestamp(6) with time zone,
    check_user_id bigint,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.stock_checks OWNER TO postgres;

--
-- Name: stock_checks_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.stock_checks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.stock_checks_id_seq OWNER TO postgres;

--
-- Name: stock_checks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.stock_checks_id_seq OWNED BY public.stock_checks.id;


--
-- Name: supplier_asn_items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.supplier_asn_items (
    id bigint NOT NULL,
    asn_id bigint NOT NULL,
    line_no bigint,
    material_code character varying(50) NOT NULL,
    material_name character varying(100),
    spec character varying(100),
    unit character varying(20),
    batch_no character varying(50),
    qty numeric(18,3),
    qualified_qty numeric(18,3),
    price numeric(18,4),
    amount numeric(18,2),
    packing_qty numeric(18,3),
    packing_unit character varying(20),
    received_qty numeric(18,3),
    tenant_id bigint NOT NULL,
    created_at timestamp with time zone
);


ALTER TABLE public.supplier_asn_items OWNER TO postgres;

--
-- Name: supplier_asn_items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.supplier_asn_items_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.supplier_asn_items_id_seq OWNER TO postgres;

--
-- Name: supplier_asn_items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.supplier_asn_items_id_seq OWNED BY public.supplier_asn_items.id;


--
-- Name: supplier_asns; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.supplier_asns (
    id bigint NOT NULL,
    asn_no character varying(50) NOT NULL,
    supplier_id bigint NOT NULL,
    supplier_code character varying(50) NOT NULL,
    supplier_name character varying(100),
    delivery_type character varying(20) DEFAULT 'NORMAL'::character varying,
    delivery_date timestamp with time zone,
    delivery_time_start character varying(20),
    delivery_time_end character varying(20),
    warehouse_code character varying(50),
    contact_person character varying(50),
    contact_phone character varying(20),
    status character varying(20) DEFAULT 'DRAFT'::character varying,
    total_qty numeric(18,3),
    total_amount numeric(18,2),
    remark text,
    tenant_id bigint NOT NULL,
    created_by bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.supplier_asns OWNER TO postgres;

--
-- Name: supplier_asns_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.supplier_asns_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.supplier_asns_id_seq OWNER TO postgres;

--
-- Name: supplier_asns_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.supplier_asns_id_seq OWNED BY public.supplier_asns.id;


--
-- Name: sys_dept; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_dept (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    parent_id bigint DEFAULT 0,
    dept_name character varying(50) NOT NULL,
    dept_code character varying(50),
    dept_sort bigint DEFAULT 0,
    leader character varying(50),
    phone character varying(20),
    email character varying(100),
    status bigint DEFAULT 1
);


ALTER TABLE public.sys_dept OWNER TO postgres;

--
-- Name: sys_dept_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_dept_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_dept_id_seq OWNER TO postgres;

--
-- Name: sys_dept_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_dept_id_seq OWNED BY public.sys_dept.id;


--
-- Name: sys_dict_data; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_dict_data (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    dict_sort bigint DEFAULT 0,
    dict_label character varying(100) NOT NULL,
    dict_value character varying(100) NOT NULL,
    dict_type character varying(100) NOT NULL,
    dict_key character varying(100),
    css_class character varying(100),
    list_class character varying(100),
    is_default bigint DEFAULT 0,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.sys_dict_data OWNER TO postgres;

--
-- Name: sys_dict_data_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_dict_data_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_dict_data_id_seq OWNER TO postgres;

--
-- Name: sys_dict_data_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_dict_data_id_seq OWNED BY public.sys_dict_data.id;


--
-- Name: sys_dict_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_dict_type (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    dict_name character varying(100) NOT NULL,
    dict_type character varying(100) NOT NULL,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.sys_dict_type OWNER TO postgres;

--
-- Name: sys_dict_type_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_dict_type_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_dict_type_id_seq OWNER TO postgres;

--
-- Name: sys_dict_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_dict_type_id_seq OWNED BY public.sys_dict_type.id;


--
-- Name: sys_interface_config; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_interface_config (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    name character varying(100) NOT NULL,
    code character varying(50) NOT NULL,
    category character varying(20),
    description character varying(500),
    direction character varying(20),
    method character varying(10),
    base_url character varying(500),
    path character varying(500),
    auth_type character varying(20),
    auth_config text,
    request_content_type character varying(20),
    request_body_template text,
    response_format character varying(20),
    timeout bigint DEFAULT 30,
    retry_count bigint DEFAULT 0,
    retry_interval bigint DEFAULT 3,
    source_type character varying(20),
    source_table character varying(100),
    source_api character varying(200),
    source_filter text,
    source_fields character varying(500),
    primary_key character varying(50),
    incremental_field character varying(50),
    incremental_window bigint DEFAULT 0,
    status character varying(20) DEFAULT 'ENABLE'::character varying,
    remark character varying(500)
);


ALTER TABLE public.sys_interface_config OWNER TO postgres;

--
-- Name: sys_interface_config_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_interface_config_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_interface_config_id_seq OWNER TO postgres;

--
-- Name: sys_interface_config_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_interface_config_id_seq OWNED BY public.sys_interface_config.id;


--
-- Name: sys_interface_exec_log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_interface_exec_log (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    interface_config_id bigint,
    config_name character varying(100),
    trigger_type character varying(20),
    start_time timestamp with time zone,
    end_time timestamp with time zone,
    duration bigint,
    request_url character varying(500),
    request_method character varying(10),
    request_headers text,
    request_body text,
    response_body text,
    response_code bigint,
    status character varying(20),
    error_message character varying(1000),
    retry_count bigint DEFAULT 0,
    records_processed bigint
);


ALTER TABLE public.sys_interface_exec_log OWNER TO postgres;

--
-- Name: sys_interface_exec_log_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_interface_exec_log_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_interface_exec_log_id_seq OWNER TO postgres;

--
-- Name: sys_interface_exec_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_interface_exec_log_id_seq OWNED BY public.sys_interface_exec_log.id;


--
-- Name: sys_interface_field_map; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_interface_field_map (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    interface_config_id bigint NOT NULL,
    field_name character varying(100) NOT NULL,
    field_type character varying(20),
    map_type character varying(20),
    map_value character varying(500),
    required boolean DEFAULT false,
    default_value character varying(200),
    transform_func character varying(100),
    sort_order bigint DEFAULT 0
);


ALTER TABLE public.sys_interface_field_map OWNER TO postgres;

--
-- Name: sys_interface_field_map_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_interface_field_map_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_interface_field_map_id_seq OWNER TO postgres;

--
-- Name: sys_interface_field_map_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_interface_field_map_id_seq OWNED BY public.sys_interface_field_map.id;


--
-- Name: sys_interface_trigger; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_interface_trigger (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    interface_config_id bigint NOT NULL,
    trigger_type character varying(20) NOT NULL,
    cron_expr character varying(50),
    event_source character varying(50),
    payload_template text,
    condition character varying(500),
    fallback_minutes bigint DEFAULT 0,
    status character varying(20) DEFAULT 'ENABLE'::character varying
);


ALTER TABLE public.sys_interface_trigger OWNER TO postgres;

--
-- Name: sys_interface_trigger_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_interface_trigger_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_interface_trigger_id_seq OWNER TO postgres;

--
-- Name: sys_interface_trigger_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_interface_trigger_id_seq OWNED BY public.sys_interface_trigger.id;


--
-- Name: sys_login_log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_login_log (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint,
    username character varying(50),
    ip character varying(50),
    login_location character varying(255),
    browser character varying(50),
    os character varying(50),
    status bigint DEFAULT 0,
    msg character varying(255),
    login_time timestamp(6) without time zone NOT NULL
);


ALTER TABLE public.sys_login_log OWNER TO postgres;

--
-- Name: sys_login_log_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_login_log_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_login_log_id_seq OWNER TO postgres;

--
-- Name: sys_login_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_login_log_id_seq OWNED BY public.sys_login_log.id;


--
-- Name: sys_menu; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_menu (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    parent_id bigint DEFAULT 0,
    menu_name character varying(50) NOT NULL,
    menu_type character varying(1),
    path character varying(200),
    component character varying(200),
    perms character varying(200),
    icon character varying(100),
    sort bigint DEFAULT 0,
    visible bigint DEFAULT 1,
    status bigint DEFAULT 1,
    is_frame bigint DEFAULT 1,
    is_cache bigint DEFAULT 0
);


ALTER TABLE public.sys_menu OWNER TO postgres;

--
-- Name: sys_menu_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_menu_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_menu_id_seq OWNER TO postgres;

--
-- Name: sys_menu_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_menu_id_seq OWNED BY public.sys_menu.id;


--
-- Name: sys_notice; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_notice (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    title character varying(200) NOT NULL,
    content text,
    notice_type character varying(20),
    priority bigint DEFAULT 1,
    publish_dept character varying(100),
    publisher_id bigint,
    publisher_name character varying(50),
    publish_time timestamp with time zone,
    effect_time timestamp with time zone,
    expire_time timestamp with time zone,
    target_type character varying(20),
    target_ids character varying(500),
    is_top bigint DEFAULT 0,
    status bigint DEFAULT 1,
    view_count bigint DEFAULT 0,
    remark character varying(500)
);


ALTER TABLE public.sys_notice OWNER TO postgres;

--
-- Name: sys_notice_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_notice_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_notice_id_seq OWNER TO postgres;

--
-- Name: sys_notice_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_notice_id_seq OWNED BY public.sys_notice.id;


--
-- Name: sys_notice_read_record; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_notice_read_record (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    notice_id bigint NOT NULL,
    user_id bigint NOT NULL,
    user_name character varying(50),
    read_time character varying(30)
);


ALTER TABLE public.sys_notice_read_record OWNER TO postgres;

--
-- Name: sys_notice_read_record_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_notice_read_record_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_notice_read_record_id_seq OWNER TO postgres;

--
-- Name: sys_notice_read_record_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_notice_read_record_id_seq OWNED BY public.sys_notice_read_record.id;


--
-- Name: sys_oper_log; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_oper_log (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint,
    title character varying(200),
    business_type character varying(20),
    method character varying(100),
    request_method character varying(10),
    operator_type bigint DEFAULT 1,
    oper_name character varying(50),
    dept_name character varying(100),
    oper_url character varying(255),
    oper_ip character varying(50),
    oper_location character varying(255),
    oper_param text,
    json_result text,
    status bigint DEFAULT 0,
    error_msg text,
    oper_time timestamp(6) without time zone NOT NULL
);


ALTER TABLE public.sys_oper_log OWNER TO postgres;

--
-- Name: sys_oper_log_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_oper_log_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_oper_log_id_seq OWNER TO postgres;

--
-- Name: sys_oper_log_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_oper_log_id_seq OWNED BY public.sys_oper_log.id;


--
-- Name: sys_post; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_post (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    post_code character varying(50) NOT NULL,
    post_name character varying(100) NOT NULL,
    post_sort bigint DEFAULT 0,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.sys_post OWNER TO postgres;

--
-- Name: sys_post_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_post_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_post_id_seq OWNER TO postgres;

--
-- Name: sys_post_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_post_id_seq OWNED BY public.sys_post.id;


--
-- Name: sys_print_template; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_print_template (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    template_code character varying(50) NOT NULL,
    template_name character varying(100) NOT NULL,
    template_type character varying(30) NOT NULL,
    entity_type character varying(50),
    content text,
    paper_type character varying(20),
    paper_width numeric,
    paper_height numeric,
    is_default bigint DEFAULT 0,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.sys_print_template OWNER TO postgres;

--
-- Name: sys_print_template_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_print_template_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_print_template_id_seq OWNER TO postgres;

--
-- Name: sys_print_template_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_print_template_id_seq OWNED BY public.sys_print_template.id;


--
-- Name: sys_role; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_role (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    role_name character varying(50) NOT NULL,
    role_key character varying(100) NOT NULL,
    role_sort bigint DEFAULT 0,
    data_scope bigint DEFAULT 1,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.sys_role OWNER TO postgres;

--
-- Name: sys_role_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_role_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_role_id_seq OWNER TO postgres;

--
-- Name: sys_role_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_role_id_seq OWNED BY public.sys_role.id;


--
-- Name: sys_role_menu; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_role_menu (
    role_id bigint NOT NULL,
    menu_id bigint NOT NULL
);


ALTER TABLE public.sys_role_menu OWNER TO postgres;

--
-- Name: sys_role_perm; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_role_perm (
    role_id bigint NOT NULL,
    perm character varying(100) NOT NULL
);


ALTER TABLE public.sys_role_perm OWNER TO postgres;

--
-- Name: sys_tenant; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_tenant (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_name character varying(100) NOT NULL,
    tenant_key character varying(50) NOT NULL,
    contact character varying(50),
    phone character varying(20),
    email character varying(100),
    status bigint DEFAULT 1,
    expire_time timestamp(6) without time zone,
    package_id bigint,
    province character varying(50),
    city character varying(50),
    district character varying(50),
    address character varying(255),
    manager character varying(50),
    contact_name character varying(50),
    contact_phone character varying(20),
    contact_email character varying(100),
    factory_type character varying(50),
    employee_count bigint,
    area_size numeric,
    annual_capacity numeric,
    remark text
);


ALTER TABLE public.sys_tenant OWNER TO postgres;

--
-- Name: sys_tenant_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_tenant_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_tenant_id_seq OWNER TO postgres;

--
-- Name: sys_tenant_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_tenant_id_seq OWNED BY public.sys_tenant.id;


--
-- Name: sys_user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_user (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp(6) with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    username character varying(50) NOT NULL,
    nickname character varying(50),
    password character varying(200) NOT NULL,
    email character varying(100),
    phone character varying(20),
    avatar character varying(500),
    dept_id bigint,
    status bigint DEFAULT 1,
    login_ip character varying(128),
    login_date timestamp(6) without time zone
);


ALTER TABLE public.sys_user OWNER TO postgres;

--
-- Name: sys_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_user_id_seq OWNER TO postgres;

--
-- Name: sys_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_user_id_seq OWNED BY public.sys_user.id;


--
-- Name: sys_user_role; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_user_role (
    user_id bigint NOT NULL,
    role_id bigint NOT NULL
);


ALTER TABLE public.sys_user_role OWNER TO postgres;

--
-- Name: tra_andon_call; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tra_andon_call (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    call_no character varying(50) NOT NULL,
    line_id bigint,
    line_name character varying(100),
    station_id bigint,
    station_name character varying(100),
    call_type character varying(20),
    call_level bigint DEFAULT 1,
    call_desc character varying(500),
    call_user_id bigint,
    call_user_name character varying(50),
    call_time timestamp(6) without time zone NOT NULL,
    response_user_id bigint,
    response_time timestamp(6) without time zone,
    resolve_time timestamp(6) without time zone,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.tra_andon_call OWNER TO postgres;

--
-- Name: tra_andon_call_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tra_andon_call_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tra_andon_call_id_seq OWNER TO postgres;

--
-- Name: tra_andon_call_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tra_andon_call_id_seq OWNED BY public.tra_andon_call.id;


--
-- Name: tra_data_collection; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tra_data_collection (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    equipment_id bigint,
    equipment_code character varying(50),
    station_id bigint,
    data_type character varying(20),
    data_key character varying(50),
    data_value character varying(200),
    unit character varying(20),
    collect_time timestamp(6) without time zone NOT NULL
);


ALTER TABLE public.tra_data_collection OWNER TO postgres;

--
-- Name: tra_data_collection_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tra_data_collection_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tra_data_collection_id_seq OWNER TO postgres;

--
-- Name: tra_data_collection_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tra_data_collection_id_seq OWNED BY public.tra_data_collection.id;


--
-- Name: tra_serial_number; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tra_serial_number (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    serial_number character varying(50) NOT NULL,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    batch_no character varying(50),
    line_id bigint,
    line_name character varying(100),
    order_id bigint,
    order_no character varying(50),
    production_date timestamp(6) without time zone,
    status bigint DEFAULT 1
);


ALTER TABLE public.tra_serial_number OWNER TO postgres;

--
-- Name: tra_serial_number_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tra_serial_number_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tra_serial_number_id_seq OWNER TO postgres;

--
-- Name: tra_serial_number_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tra_serial_number_id_seq OWNED BY public.tra_serial_number.id;


--
-- Name: tra_trace_record; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tra_trace_record (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    trace_no character varying(50) NOT NULL,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    serial_number character varying(50),
    batch_no character varying(50),
    process_id bigint,
    process_name character varying(100),
    station_id bigint,
    station_name character varying(100),
    operator_id bigint,
    operator_name character varying(50),
    operate_time timestamp(6) without time zone NOT NULL,
    operate_type character varying(20),
    input_qty numeric(18,4),
    output_qty numeric(18,4),
    reject_qty numeric(18,4) DEFAULT 0
);


ALTER TABLE public.tra_trace_record OWNER TO postgres;

--
-- Name: tra_trace_record_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tra_trace_record_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tra_trace_record_id_seq OWNER TO postgres;

--
-- Name: tra_trace_record_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tra_trace_record_id_seq OWNED BY public.tra_trace_record.id;


--
-- Name: trace_records; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.trace_records (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    trace_no character varying(50) NOT NULL,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    serial_number character varying(50),
    batch_no character varying(50),
    process_id bigint,
    process_name character varying(100),
    station_id bigint,
    station_name character varying(100),
    operator_id bigint,
    operator_name character varying(50),
    operate_time timestamp(6) with time zone,
    operate_type character varying(20),
    input_qty numeric(18,4),
    output_qty numeric(18,4),
    reject_qty numeric(18,4) DEFAULT 0
);


ALTER TABLE public.trace_records OWNER TO postgres;

--
-- Name: trace_records_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.trace_records_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.trace_records_id_seq OWNER TO postgres;

--
-- Name: trace_records_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.trace_records_id_seq OWNED BY public.trace_records.id;


--
-- Name: transfer_orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transfer_orders (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    transfer_no character varying(50) NOT NULL,
    from_warehouse_id bigint,
    to_warehouse_id bigint,
    transfer_date timestamp(6) with time zone,
    transfer_user_id bigint,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.transfer_orders OWNER TO postgres;

--
-- Name: transfer_orders_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.transfer_orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.transfer_orders_id_seq OWNER TO postgres;

--
-- Name: transfer_orders_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.transfer_orders_id_seq OWNED BY public.transfer_orders.id;


--
-- Name: visual_inspection_result; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.visual_inspection_result (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    task_id bigint NOT NULL,
    detection_time timestamp with time zone,
    result character varying(10) NOT NULL,
    confidence numeric(5,4),
    defect_type character varying(50),
    defect_location json,
    defect_image_url character varying(500),
    ai_analysis json,
    manual_review_result character varying(20),
    manual_review_by bigint,
    manual_review_at timestamp with time zone,
    remark character varying(500),
    tenant_id bigint NOT NULL
);


ALTER TABLE public.visual_inspection_result OWNER TO postgres;

--
-- Name: visual_inspection_result_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.visual_inspection_result_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.visual_inspection_result_id_seq OWNER TO postgres;

--
-- Name: visual_inspection_result_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.visual_inspection_result_id_seq OWNED BY public.visual_inspection_result.id;


--
-- Name: visual_inspection_task; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.visual_inspection_task (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    task_no character varying(50) NOT NULL,
    task_type character varying(30) NOT NULL,
    product_id bigint,
    product_code character varying(50),
    product_name character varying(100),
    production_order_id bigint,
    workshop_id bigint,
    image_url character varying(500),
    image_hash character varying(64),
    detection_standard json,
    ai_model_version character varying(50),
    status character varying(20) DEFAULT 'PENDING'::character varying,
    priority character varying(10) DEFAULT 'NORMAL'::character varying,
    requested_by character varying(50),
    requested_at timestamp with time zone,
    completed_at timestamp with time zone,
    remark character varying(500),
    tenant_id bigint NOT NULL,
    created_by bigint
);


ALTER TABLE public.visual_inspection_task OWNER TO postgres;

--
-- Name: visual_inspection_task_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.visual_inspection_task_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.visual_inspection_task_id_seq OWNER TO postgres;

--
-- Name: visual_inspection_task_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.visual_inspection_task_id_seq OWNED BY public.visual_inspection_task.id;


--
-- Name: warehouses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.warehouses (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    warehouse_code character varying(50) NOT NULL,
    warehouse_name character varying(100) NOT NULL,
    warehouse_type character varying(20),
    address character varying(200),
    manager character varying(50),
    phone character varying(20),
    status bigint DEFAULT 1
);


ALTER TABLE public.warehouses OWNER TO postgres;

--
-- Name: warehouses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.warehouses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.warehouses_id_seq OWNER TO postgres;

--
-- Name: warehouses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.warehouses_id_seq OWNED BY public.warehouses.id;


--
-- Name: wms_agv_device; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_agv_device (
    id bigint NOT NULL,
    tenant_id bigint DEFAULT 1,
    agv_code character varying(50) NOT NULL,
    agv_name character varying(100),
    agv_type character varying(30),
    status character varying(20) DEFAULT 'OFFLINE'::character varying,
    current_location character varying(100),
    battery_level numeric(5,2),
    max_load numeric(18,6),
    ext_config jsonb,
    last_heartbeat timestamp with time zone,
    remark text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.wms_agv_device OWNER TO postgres;

--
-- Name: wms_agv_device_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_agv_device_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_agv_device_id_seq OWNER TO postgres;

--
-- Name: wms_agv_device_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_agv_device_id_seq OWNED BY public.wms_agv_device.id;


--
-- Name: wms_agv_location_mapping; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_agv_location_mapping (
    id bigint NOT NULL,
    tenant_id bigint DEFAULT 1,
    location_code character varying(50) NOT NULL,
    location_name character varying(100),
    location_type character varying(30),
    agv_location_code character varying(50),
    x_coord numeric(10,2),
    y_coord numeric(10,2),
    priority bigint DEFAULT 0,
    enabled boolean DEFAULT true,
    ext_data jsonb,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.wms_agv_location_mapping OWNER TO postgres;

--
-- Name: wms_agv_location_mapping_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_agv_location_mapping_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_agv_location_mapping_id_seq OWNER TO postgres;

--
-- Name: wms_agv_location_mapping_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_agv_location_mapping_id_seq OWNED BY public.wms_agv_location_mapping.id;


--
-- Name: wms_agv_task; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_agv_task (
    id bigint NOT NULL,
    tenant_id bigint DEFAULT 1,
    task_no character varying(50) NOT NULL,
    task_type character varying(20) NOT NULL,
    status character varying(20) DEFAULT 'PENDING'::character varying,
    priority bigint DEFAULT 0,
    source_location_id bigint,
    source_location character varying(100),
    target_location_id bigint,
    target_location character varying(100),
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    quantity numeric(18,6),
    unit character varying(20),
    assigned_agv_code character varying(50),
    assigned_agv_name character varying(100),
    started_at timestamp with time zone,
    completed_at timestamp with time zone,
    error_message text,
    related_order_no character varying(50),
    related_order_type character varying(50),
    ext_data jsonb,
    created_by character varying(50),
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.wms_agv_task OWNER TO postgres;

--
-- Name: wms_agv_task_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_agv_task_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_agv_task_id_seq OWNER TO postgres;

--
-- Name: wms_agv_task_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_agv_task_id_seq OWNED BY public.wms_agv_task.id;


--
-- Name: wms_delivery_order; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_delivery_order (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    delivery_no character varying(50) NOT NULL,
    customer_id bigint,
    customer_name character varying(100),
    warehouse_id bigint,
    delivery_date timestamp(6) without time zone,
    delivery_user_id bigint,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.wms_delivery_order OWNER TO postgres;

--
-- Name: wms_delivery_order_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_delivery_order_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_delivery_order_id_seq OWNER TO postgres;

--
-- Name: wms_delivery_order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_delivery_order_id_seq OWNED BY public.wms_delivery_order.id;


--
-- Name: wms_delivery_order_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_delivery_order_item (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    delivery_id bigint,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    quantity numeric(18,4),
    shipped_qty numeric(18,4) DEFAULT 0,
    unit character varying(20),
    batch_no character varying(50)
);


ALTER TABLE public.wms_delivery_order_item OWNER TO postgres;

--
-- Name: wms_delivery_order_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_delivery_order_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_delivery_order_item_id_seq OWNER TO postgres;

--
-- Name: wms_delivery_order_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_delivery_order_item_id_seq OWNED BY public.wms_delivery_order_item.id;


--
-- Name: wms_inventory; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_inventory (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    warehouse_id bigint,
    location_id bigint,
    quantity numeric(18,4) DEFAULT 0,
    available_qty numeric(18,4) DEFAULT 0,
    allocated_qty numeric(18,4) DEFAULT 0,
    locked_qty numeric(18,4) DEFAULT 0,
    batch_no character varying(50)
);


ALTER TABLE public.wms_inventory OWNER TO postgres;

--
-- Name: wms_inventory_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_inventory_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_inventory_id_seq OWNER TO postgres;

--
-- Name: wms_inventory_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_inventory_id_seq OWNED BY public.wms_inventory.id;


--
-- Name: wms_inventory_record; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_inventory_record (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    record_no character varying(50) NOT NULL,
    record_type character varying(20),
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    warehouse_id bigint,
    location_id bigint,
    quantity numeric(18,4),
    batch_no character varying(50),
    source_no character varying(50),
    operator_id bigint,
    operator_name character varying(50),
    operate_time timestamp(6) without time zone,
    remark character varying(500)
);


ALTER TABLE public.wms_inventory_record OWNER TO postgres;

--
-- Name: wms_inventory_record_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_inventory_record_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_inventory_record_id_seq OWNER TO postgres;

--
-- Name: wms_inventory_record_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_inventory_record_id_seq OWNED BY public.wms_inventory_record.id;


--
-- Name: wms_kanban_pull; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_kanban_pull (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    kanban_no character varying(50) NOT NULL,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    from_location_id bigint,
    to_location_id bigint,
    kanban_qty numeric,
    trigger_qty numeric,
    current_qty numeric,
    lead_time bigint,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.wms_kanban_pull OWNER TO postgres;

--
-- Name: wms_kanban_pull_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_kanban_pull_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_kanban_pull_id_seq OWNER TO postgres;

--
-- Name: wms_kanban_pull_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_kanban_pull_id_seq OWNED BY public.wms_kanban_pull.id;


--
-- Name: wms_location; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_location (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    location_code character varying(50) NOT NULL,
    location_name character varying(100),
    warehouse_id bigint,
    zone_code character varying(20),
    "row" bigint,
    col bigint,
    layer bigint,
    location_type character varying(20),
    capacity bigint,
    status bigint DEFAULT 1
);


ALTER TABLE public.wms_location OWNER TO postgres;

--
-- Name: wms_location_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_location_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_location_id_seq OWNER TO postgres;

--
-- Name: wms_location_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_location_id_seq OWNED BY public.wms_location.id;


--
-- Name: wms_production_complete; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_production_complete (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    complete_no character varying(50) NOT NULL,
    production_order_id bigint,
    order_no character varying(50),
    workshop_id bigint,
    workstation_id bigint,
    complete_qty numeric(18,3) DEFAULT 0 NOT NULL,
    qualified_qty numeric(18,3) DEFAULT 0,
    status character varying(20) DEFAULT 'PENDING'::character varying NOT NULL,
    complete_time timestamp with time zone,
    operator_id bigint,
    operator_name character varying(50),
    remark text,
    tenant_id bigint NOT NULL
);


ALTER TABLE public.wms_production_complete OWNER TO postgres;

--
-- Name: wms_production_complete_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_production_complete_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_production_complete_id_seq OWNER TO postgres;

--
-- Name: wms_production_complete_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_production_complete_id_seq OWNED BY public.wms_production_complete.id;


--
-- Name: wms_production_complete_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_production_complete_item (
    id bigint NOT NULL,
    complete_id bigint NOT NULL,
    line_no bigint NOT NULL,
    material_id bigint NOT NULL,
    material_code character varying(50),
    material_name character varying(100),
    unit character varying(20),
    complete_qty numeric(18,3) DEFAULT 0 NOT NULL,
    qualified_qty numeric(18,3) DEFAULT 0,
    warehouse_id bigint,
    location_id bigint,
    batch_no character varying(50),
    tenant_id bigint NOT NULL
);


ALTER TABLE public.wms_production_complete_item OWNER TO postgres;

--
-- Name: wms_production_complete_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_production_complete_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_production_complete_item_id_seq OWNER TO postgres;

--
-- Name: wms_production_complete_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_production_complete_item_id_seq OWNED BY public.wms_production_complete_item.id;


--
-- Name: wms_production_issue; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_production_issue (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    issue_no character varying(50) NOT NULL,
    issue_type character varying(20) NOT NULL,
    production_order_id bigint NOT NULL,
    order_no character varying(50),
    workstation_id bigint,
    workshop_id bigint,
    status character varying(20) DEFAULT 'PENDING'::character varying NOT NULL,
    pick_status character varying(20) DEFAULT 'PENDING'::character varying,
    request_by bigint,
    request_time timestamp with time zone,
    issued_by bigint,
    issued_time timestamp with time zone,
    remark text,
    tenant_id bigint NOT NULL,
    created_by character varying(50)
);


ALTER TABLE public.wms_production_issue OWNER TO postgres;

--
-- Name: wms_production_issue_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_production_issue_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_production_issue_id_seq OWNER TO postgres;

--
-- Name: wms_production_issue_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_production_issue_id_seq OWNED BY public.wms_production_issue.id;


--
-- Name: wms_production_issue_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_production_issue_item (
    id bigint NOT NULL,
    issue_id bigint NOT NULL,
    line_no bigint NOT NULL,
    material_id bigint NOT NULL,
    material_code character varying(50),
    material_name character varying(100),
    unit character varying(20),
    required_qty numeric(18,3) NOT NULL,
    picked_qty numeric(18,3) DEFAULT 0,
    issued_qty numeric(18,3) DEFAULT 0,
    warehouse_id bigint,
    location_id bigint,
    batch_no character varying(50),
    remark text,
    tenant_id bigint NOT NULL
);


ALTER TABLE public.wms_production_issue_item OWNER TO postgres;

--
-- Name: wms_production_issue_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_production_issue_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_production_issue_item_id_seq OWNER TO postgres;

--
-- Name: wms_production_issue_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_production_issue_item_id_seq OWNED BY public.wms_production_issue_item.id;


--
-- Name: wms_production_return; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_production_return (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    return_no character varying(50) NOT NULL,
    production_order_id bigint,
    order_no character varying(50),
    workstation_id bigint,
    workshop_id bigint,
    status character varying(20) DEFAULT 'PENDING'::character varying NOT NULL,
    return_type character varying(20) NOT NULL,
    request_by bigint,
    request_time timestamp with time zone,
    approved_by bigint,
    approved_time timestamp with time zone,
    returned_by bigint,
    returned_time timestamp with time zone,
    remark text,
    tenant_id bigint NOT NULL,
    created_by character varying(50)
);


ALTER TABLE public.wms_production_return OWNER TO postgres;

--
-- Name: wms_production_return_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_production_return_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_production_return_id_seq OWNER TO postgres;

--
-- Name: wms_production_return_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_production_return_id_seq OWNED BY public.wms_production_return.id;


--
-- Name: wms_production_return_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_production_return_item (
    id bigint NOT NULL,
    return_id bigint NOT NULL,
    line_no bigint NOT NULL,
    material_id bigint NOT NULL,
    material_code character varying(50),
    material_name character varying(100),
    unit character varying(20),
    issued_qty numeric(18,3) NOT NULL,
    return_qty numeric(18,3) DEFAULT 0,
    warehouse_id bigint,
    location_id bigint,
    batch_no character varying(50),
    reason character varying(200),
    remark text,
    tenant_id bigint NOT NULL
);


ALTER TABLE public.wms_production_return_item OWNER TO postgres;

--
-- Name: wms_production_return_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_production_return_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_production_return_item_id_seq OWNER TO postgres;

--
-- Name: wms_production_return_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_production_return_item_id_seq OWNED BY public.wms_production_return_item.id;


--
-- Name: wms_production_stock_in; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_production_stock_in (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    stock_in_no character varying(50) NOT NULL,
    complete_id bigint NOT NULL,
    complete_no character varying(50),
    warehouse_id bigint,
    location_id bigint,
    status character varying(20) DEFAULT 'PENDING'::character varying NOT NULL,
    stock_in_time timestamp with time zone,
    operator_id bigint,
    operator_name character varying(50),
    tenant_id bigint NOT NULL
);


ALTER TABLE public.wms_production_stock_in OWNER TO postgres;

--
-- Name: wms_production_stock_in_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_production_stock_in_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_production_stock_in_id_seq OWNER TO postgres;

--
-- Name: wms_production_stock_in_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_production_stock_in_id_seq OWNED BY public.wms_production_stock_in.id;


--
-- Name: wms_production_stock_in_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_production_stock_in_item (
    id bigint NOT NULL,
    stock_in_id bigint NOT NULL,
    line_no bigint NOT NULL,
    material_id bigint NOT NULL,
    material_code character varying(50),
    material_name character varying(100),
    unit character varying(20),
    stock_in_qty numeric(18,3) DEFAULT 0 NOT NULL,
    warehouse_id bigint,
    location_id bigint,
    batch_no character varying(50),
    tenant_id bigint NOT NULL
);


ALTER TABLE public.wms_production_stock_in_item OWNER TO postgres;

--
-- Name: wms_production_stock_in_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_production_stock_in_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_production_stock_in_item_id_seq OWNER TO postgres;

--
-- Name: wms_production_stock_in_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_production_stock_in_item_id_seq OWNED BY public.wms_production_stock_in_item.id;


--
-- Name: wms_purchase_return; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_purchase_return (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    return_no character varying(50) NOT NULL,
    purchase_order_id bigint,
    supplier_id bigint,
    supplier_name character varying(100),
    warehouse_id bigint,
    return_date date,
    status character varying(20) DEFAULT 'PENDING'::character varying NOT NULL,
    return_type character varying(20) NOT NULL,
    request_by bigint,
    request_time timestamp with time zone,
    approved_by bigint,
    approved_time timestamp with time zone,
    returned_by bigint,
    returned_time timestamp with time zone,
    remark text,
    tenant_id bigint NOT NULL,
    created_by character varying(50)
);


ALTER TABLE public.wms_purchase_return OWNER TO postgres;

--
-- Name: wms_purchase_return_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_purchase_return_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_purchase_return_id_seq OWNER TO postgres;

--
-- Name: wms_purchase_return_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_purchase_return_id_seq OWNED BY public.wms_purchase_return.id;


--
-- Name: wms_purchase_return_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_purchase_return_item (
    id bigint NOT NULL,
    return_id bigint NOT NULL,
    line_no bigint NOT NULL,
    material_id bigint NOT NULL,
    material_code character varying(50),
    material_name character varying(100),
    unit character varying(20),
    return_qty numeric(18,3) DEFAULT 0 NOT NULL,
    reason character varying(200),
    tenant_id bigint NOT NULL
);


ALTER TABLE public.wms_purchase_return_item OWNER TO postgres;

--
-- Name: wms_purchase_return_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_purchase_return_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_purchase_return_item_id_seq OWNER TO postgres;

--
-- Name: wms_purchase_return_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_purchase_return_item_id_seq OWNED BY public.wms_purchase_return_item.id;


--
-- Name: wms_receive_order; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_receive_order (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    receive_no character varying(50) NOT NULL,
    supplier_id bigint,
    supplier_name character varying(100),
    warehouse_id bigint,
    receive_date timestamp(6) without time zone,
    receive_user_id bigint,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.wms_receive_order OWNER TO postgres;

--
-- Name: wms_receive_order_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_receive_order_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_receive_order_id_seq OWNER TO postgres;

--
-- Name: wms_receive_order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_receive_order_id_seq OWNED BY public.wms_receive_order.id;


--
-- Name: wms_receive_order_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_receive_order_item (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    receive_id bigint,
    material_id bigint,
    material_code character varying(50),
    material_name character varying(100),
    quantity numeric(18,4),
    received_qty numeric(18,4) DEFAULT 0,
    unit character varying(20),
    batch_no character varying(50)
);


ALTER TABLE public.wms_receive_order_item OWNER TO postgres;

--
-- Name: wms_receive_order_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_receive_order_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_receive_order_item_id_seq OWNER TO postgres;

--
-- Name: wms_receive_order_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_receive_order_item_id_seq OWNED BY public.wms_receive_order_item.id;


--
-- Name: wms_sales_return; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_sales_return (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    return_no character varying(50) NOT NULL,
    sales_order_id bigint,
    customer_id bigint,
    customer_name character varying(100),
    warehouse_id bigint,
    return_date date,
    status character varying(20) DEFAULT 'PENDING'::character varying NOT NULL,
    return_type character varying(20) NOT NULL,
    request_by bigint,
    request_time timestamp with time zone,
    approved_by bigint,
    approved_time timestamp with time zone,
    returned_by bigint,
    returned_time timestamp with time zone,
    remark text,
    tenant_id bigint NOT NULL,
    created_by character varying(50)
);


ALTER TABLE public.wms_sales_return OWNER TO postgres;

--
-- Name: wms_sales_return_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_sales_return_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_sales_return_id_seq OWNER TO postgres;

--
-- Name: wms_sales_return_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_sales_return_id_seq OWNED BY public.wms_sales_return.id;


--
-- Name: wms_sales_return_item; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_sales_return_item (
    id bigint NOT NULL,
    return_id bigint NOT NULL,
    line_no bigint NOT NULL,
    material_id bigint NOT NULL,
    material_code character varying(50),
    material_name character varying(100),
    unit character varying(20),
    return_qty numeric(18,3) DEFAULT 0 NOT NULL,
    reason character varying(200),
    tenant_id bigint NOT NULL
);


ALTER TABLE public.wms_sales_return_item OWNER TO postgres;

--
-- Name: wms_sales_return_item_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_sales_return_item_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_sales_return_item_id_seq OWNER TO postgres;

--
-- Name: wms_sales_return_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_sales_return_item_id_seq OWNED BY public.wms_sales_return_item.id;


--
-- Name: wms_side_location; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_side_location (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    tenant_id bigint NOT NULL,
    location_code character varying(50) NOT NULL,
    location_name character varying(100),
    workshop_id bigint,
    workshop_name character varying(100),
    line_id bigint,
    line_name character varying(100),
    station_id bigint,
    station_name character varying(100),
    location_type character varying(20),
    max_capacity numeric,
    current_qty numeric,
    status bigint DEFAULT 1,
    remark character varying(500)
);


ALTER TABLE public.wms_side_location OWNER TO postgres;

--
-- Name: wms_side_location_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_side_location_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_side_location_id_seq OWNER TO postgres;

--
-- Name: wms_side_location_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_side_location_id_seq OWNED BY public.wms_side_location.id;


--
-- Name: wms_stock_check; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_stock_check (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    check_no character varying(50) NOT NULL,
    warehouse_id bigint,
    check_date timestamp(6) without time zone,
    check_user_id bigint,
    status character varying(20) DEFAULT 'DRAFT'::character varying,
    remark text,
    check_type character varying(20),
    warehouse_name character varying(100),
    area_ids jsonb,
    location_ids jsonb,
    plan_start_date character varying(10),
    plan_end_date character varying(10),
    actual_start_date character varying(10),
    actual_end_date character varying(10),
    include_zero_stock bigint DEFAULT 1,
    include_expired_stock bigint DEFAULT 0,
    is_blind_mode bigint DEFAULT 0,
    checker_id bigint,
    checker_name character varying(50),
    audit_id bigint,
    audit_name character varying(50),
    total_locations bigint DEFAULT 0,
    counted_locations bigint DEFAULT 0,
    total_materials bigint DEFAULT 0,
    variance_count bigint DEFAULT 0,
    variance_rate numeric(5,2),
    approval_status character varying(20) DEFAULT 'PENDING'::character varying,
    approved_by bigint,
    approved_time timestamp with time zone,
    approval_comment text,
    workshop_id bigint
);


ALTER TABLE public.wms_stock_check OWNER TO postgres;

--
-- Name: wms_stock_check_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_stock_check_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_stock_check_id_seq OWNER TO postgres;

--
-- Name: wms_stock_check_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_stock_check_id_seq OWNED BY public.wms_stock_check.id;


--
-- Name: wms_transfer_order; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_transfer_order (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    transfer_no character varying(50) NOT NULL,
    from_warehouse_id bigint,
    to_warehouse_id bigint,
    transfer_date timestamp(6) without time zone,
    transfer_user_id bigint,
    status character varying(20) DEFAULT 'DRAFT'::character varying,
    remark text,
    transfer_type character varying(20),
    from_warehouse_name character varying(100),
    from_warehouse_type character varying(20),
    to_warehouse_name character varying(100),
    to_warehouse_type character varying(20),
    from_workstation_id bigint,
    from_workstation_name character varying(100),
    to_workstation_id bigint,
    to_workstation_name character varying(100),
    transfer_reason character varying(100),
    source_order_id bigint,
    source_order_no character varying(50),
    requester_id bigint,
    requester_name character varying(50),
    request_time timestamp with time zone,
    approver_id bigint,
    approved_time timestamp with time zone,
    approval_comment text,
    actual_start_time timestamp with time zone,
    actual_complete_time timestamp with time zone,
    operator_id bigint,
    operator_name character varying(50),
    logistics_provider character varying(100),
    tracking_no character varying(100),
    total_amount numeric(18,2) DEFAULT 0,
    currency character varying(10) DEFAULT 'CNY'::character varying,
    workshop_id bigint
);


ALTER TABLE public.wms_transfer_order OWNER TO postgres;

--
-- Name: wms_transfer_order_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_transfer_order_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_transfer_order_id_seq OWNER TO postgres;

--
-- Name: wms_transfer_order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_transfer_order_id_seq OWNED BY public.wms_transfer_order.id;


--
-- Name: wms_warehouse; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.wms_warehouse (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp(6) without time zone,
    tenant_id bigint NOT NULL,
    warehouse_code character varying(50) NOT NULL,
    warehouse_name character varying(100) NOT NULL,
    warehouse_type character varying(20),
    address character varying(200),
    manager character varying(50),
    phone character varying(20),
    status bigint DEFAULT 1
);


ALTER TABLE public.wms_warehouse OWNER TO postgres;

--
-- Name: wms_warehouse_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.wms_warehouse_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.wms_warehouse_id_seq OWNER TO postgres;

--
-- Name: wms_warehouse_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.wms_warehouse_id_seq OWNED BY public.wms_warehouse.id;


--
-- Name: work_centers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.work_centers (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    work_center_code character varying(50) NOT NULL,
    work_center_name character varying(100) NOT NULL,
    workshop_id bigint,
    capacity numeric(18,2),
    status bigint DEFAULT 1
);


ALTER TABLE public.work_centers OWNER TO postgres;

--
-- Name: work_centers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.work_centers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.work_centers_id_seq OWNER TO postgres;

--
-- Name: work_centers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.work_centers_id_seq OWNED BY public.work_centers.id;


--
-- Name: workshops; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.workshops (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    workshop_code character varying(50) NOT NULL,
    workshop_name character varying(100) NOT NULL,
    workshop_type character varying(20),
    manager character varying(50),
    phone character varying(20),
    address character varying(200),
    status bigint DEFAULT 1
);


ALTER TABLE public.workshops OWNER TO postgres;

--
-- Name: workshops_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.workshops_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.workshops_id_seq OWNER TO postgres;

--
-- Name: workshops_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.workshops_id_seq OWNED BY public.workshops.id;


--
-- Name: workstations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.workstations (
    id bigint NOT NULL,
    created_at timestamp(6) with time zone,
    updated_at timestamp(6) with time zone,
    deleted_at timestamp(6) with time zone,
    tenant_id bigint NOT NULL,
    station_code character varying(50) NOT NULL,
    station_name character varying(100) NOT NULL,
    line_id bigint NOT NULL,
    station_type character varying(20),
    status bigint DEFAULT 1
);


ALTER TABLE public.workstations OWNER TO postgres;

--
-- Name: workstations_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.workstations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.workstations_id_seq OWNER TO postgres;

--
-- Name: workstations_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.workstations_id_seq OWNED BY public.workstations.id;


--
-- Name: ai_chat_conversation id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ai_chat_conversation ALTER COLUMN id SET DEFAULT nextval('public.ai_chat_conversation_id_seq'::regclass);


--
-- Name: ai_chat_message id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ai_chat_message ALTER COLUMN id SET DEFAULT nextval('public.ai_chat_message_id_seq'::regclass);


--
-- Name: ai_config id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ai_config ALTER COLUMN id SET DEFAULT nextval('public.ai_config_id_seq'::regclass);


--
-- Name: alert_escalation_rule id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alert_escalation_rule ALTER COLUMN id SET DEFAULT nextval('public.alert_escalation_rule_id_seq'::regclass);


--
-- Name: alert_notification_log id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alert_notification_log ALTER COLUMN id SET DEFAULT nextval('public.alert_notification_log_id_seq'::regclass);


--
-- Name: alert_record id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alert_record ALTER COLUMN id SET DEFAULT nextval('public.alert_record_id_seq'::regclass);


--
-- Name: alert_rule id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alert_rule ALTER COLUMN id SET DEFAULT nextval('public.alert_rule_id_seq'::regclass);


--
-- Name: andon_call id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.andon_call ALTER COLUMN id SET DEFAULT nextval('public.andon_call_id_seq'::regclass);


--
-- Name: andon_calls id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.andon_calls ALTER COLUMN id SET DEFAULT nextval('public.andon_calls_id_seq'::regclass);


--
-- Name: andon_escalation_log id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.andon_escalation_log ALTER COLUMN id SET DEFAULT nextval('public.andon_escalation_log_id_seq'::regclass);


--
-- Name: andon_escalation_rule id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.andon_escalation_rule ALTER COLUMN id SET DEFAULT nextval('public.andon_escalation_rule_id_seq'::regclass);


--
-- Name: andon_notification_log id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.andon_notification_log ALTER COLUMN id SET DEFAULT nextval('public.andon_notification_log_id_seq'::regclass);


--
-- Name: andon_report id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.andon_report ALTER COLUMN id SET DEFAULT nextval('public.andon_report_id_seq'::regclass);


--
-- Name: aps_capacity_analysis id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_capacity_analysis ALTER COLUMN id SET DEFAULT nextval('public.aps_capacity_analysis_id_seq'::regclass);


--
-- Name: aps_changeover_matrix id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_changeover_matrix ALTER COLUMN id SET DEFAULT nextval('public.aps_changeover_matrix_id_seq'::regclass);


--
-- Name: aps_delivery_rate id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_delivery_rate ALTER COLUMN id SET DEFAULT nextval('public.aps_delivery_rate_id_seq'::regclass);


--
-- Name: aps_jit_demand id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_jit_demand ALTER COLUMN id SET DEFAULT nextval('public.aps_jit_demand_id_seq'::regclass);


--
-- Name: aps_mps id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_mps ALTER COLUMN id SET DEFAULT nextval('public.aps_mps_id_seq'::regclass);


--
-- Name: aps_mrp id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_mrp ALTER COLUMN id SET DEFAULT nextval('public.aps_mrp_id_seq'::regclass);


--
-- Name: aps_mrp_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_mrp_item ALTER COLUMN id SET DEFAULT nextval('public.aps_mrp_item_id_seq'::regclass);


--
-- Name: aps_resource id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_resource ALTER COLUMN id SET DEFAULT nextval('public.aps_resource_id_seq'::regclass);


--
-- Name: aps_rolling_schedule id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_rolling_schedule ALTER COLUMN id SET DEFAULT nextval('public.aps_rolling_schedule_id_seq'::regclass);


--
-- Name: aps_schedule_plan id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_schedule_plan ALTER COLUMN id SET DEFAULT nextval('public.aps_schedule_plan_id_seq'::regclass);


--
-- Name: aps_schedule_result id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_schedule_result ALTER COLUMN id SET DEFAULT nextval('public.aps_schedule_result_id_seq'::regclass);


--
-- Name: aps_work_center id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_work_center ALTER COLUMN id SET DEFAULT nextval('public.aps_work_center_id_seq'::regclass);


--
-- Name: aps_working_calendar id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_working_calendar ALTER COLUMN id SET DEFAULT nextval('public.aps_working_calendar_id_seq'::regclass);


--
-- Name: bom_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bom_items ALTER COLUMN id SET DEFAULT nextval('public.bom_items_id_seq'::regclass);


--
-- Name: boms id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.boms ALTER COLUMN id SET DEFAULT nextval('public.boms_id_seq'::regclass);


--
-- Name: bpm_approval_record id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_approval_record ALTER COLUMN id SET DEFAULT nextval('public.bpm_approval_record_id_seq'::regclass);


--
-- Name: bpm_biz_mapping id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_biz_mapping ALTER COLUMN id SET DEFAULT nextval('public.bpm_biz_mapping_id_seq'::regclass);


--
-- Name: bpm_delegate_record id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_delegate_record ALTER COLUMN id SET DEFAULT nextval('public.bpm_delegate_record_id_seq'::regclass);


--
-- Name: bpm_form_definition id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_form_definition ALTER COLUMN id SET DEFAULT nextval('public.bpm_form_definition_id_seq'::regclass);


--
-- Name: bpm_form_field id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_form_field ALTER COLUMN id SET DEFAULT nextval('public.bpm_form_field_id_seq'::regclass);


--
-- Name: bpm_form_version id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_form_version ALTER COLUMN id SET DEFAULT nextval('public.bpm_form_version_id_seq'::regclass);


--
-- Name: bpm_model_version id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_model_version ALTER COLUMN id SET DEFAULT nextval('public.bpm_model_version_id_seq'::regclass);


--
-- Name: bpm_node_definition id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_node_definition ALTER COLUMN id SET DEFAULT nextval('public.bpm_node_definition_id_seq'::regclass);


--
-- Name: bpm_node_timeout id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_node_timeout ALTER COLUMN id SET DEFAULT nextval('public.bpm_node_timeout_id_seq'::regclass);


--
-- Name: bpm_order_day_workorder_map id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_order_day_workorder_map ALTER COLUMN id SET DEFAULT nextval('public.bpm_order_day_workorder_map_id_seq'::regclass);


--
-- Name: bpm_process_instance id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_process_instance ALTER COLUMN id SET DEFAULT nextval('public.bpm_process_instance_id_seq'::regclass);


--
-- Name: bpm_process_model id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_process_model ALTER COLUMN id SET DEFAULT nextval('public.bpm_process_model_id_seq'::regclass);


--
-- Name: bpm_reminder_record id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_reminder_record ALTER COLUMN id SET DEFAULT nextval('public.bpm_reminder_record_id_seq'::regclass);


--
-- Name: bpm_sequence_flow id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_sequence_flow ALTER COLUMN id SET DEFAULT nextval('public.bpm_sequence_flow_id_seq'::regclass);


--
-- Name: bpm_task_assignment id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_task_assignment ALTER COLUMN id SET DEFAULT nextval('public.bpm_task_assignment_id_seq'::regclass);


--
-- Name: bpm_task_instance id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_task_instance ALTER COLUMN id SET DEFAULT nextval('public.bpm_task_instance_id_seq'::regclass);


--
-- Name: container_master id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.container_master ALTER COLUMN id SET DEFAULT nextval('public.container_master_id_seq'::regclass);


--
-- Name: container_movement id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.container_movement ALTER COLUMN id SET DEFAULT nextval('public.container_movement_id_seq'::regclass);


--
-- Name: data_collections id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.data_collections ALTER COLUMN id SET DEFAULT nextval('public.data_collections_id_seq'::regclass);


--
-- Name: dc_collect_record id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dc_collect_record ALTER COLUMN id SET DEFAULT nextval('public.dc_collect_record_id_seq'::regclass);


--
-- Name: dc_data_point id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dc_data_point ALTER COLUMN id SET DEFAULT nextval('public.dc_data_point_id_seq'::regclass);


--
-- Name: dc_scan_log id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dc_scan_log ALTER COLUMN id SET DEFAULT nextval('public.dc_scan_log_id_seq'::regclass);


--
-- Name: defect_codes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.defect_codes ALTER COLUMN id SET DEFAULT nextval('public.defect_codes_id_seq'::regclass);


--
-- Name: defect_records id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.defect_records ALTER COLUMN id SET DEFAULT nextval('public.defect_records_id_seq'::regclass);


--
-- Name: delivery_order_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.delivery_order_items ALTER COLUMN id SET DEFAULT nextval('public.delivery_order_items_id_seq'::regclass);


--
-- Name: delivery_orders id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.delivery_orders ALTER COLUMN id SET DEFAULT nextval('public.delivery_orders_id_seq'::regclass);


--
-- Name: delivery_report id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.delivery_report ALTER COLUMN id SET DEFAULT nextval('public.delivery_report_id_seq'::regclass);


--
-- Name: dispatches id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dispatches ALTER COLUMN id SET DEFAULT nextval('public.dispatches_id_seq'::regclass);


--
-- Name: eam_equipment_downtime id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.eam_equipment_downtime ALTER COLUMN id SET DEFAULT nextval('public.eam_equipment_downtime_id_seq'::regclass);


--
-- Name: eam_repair_flow id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.eam_repair_flow ALTER COLUMN id SET DEFAULT nextval('public.eam_repair_flow_id_seq'::regclass);


--
-- Name: eam_repair_job id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.eam_repair_job ALTER COLUMN id SET DEFAULT nextval('public.eam_repair_job_id_seq'::regclass);


--
-- Name: eam_repair_std id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.eam_repair_std ALTER COLUMN id SET DEFAULT nextval('public.eam_repair_std_id_seq'::regclass);


--
-- Name: ene_energy_record id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ene_energy_record ALTER COLUMN id SET DEFAULT nextval('public.ene_energy_record_id_seq'::regclass);


--
-- Name: energy_records id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.energy_records ALTER COLUMN id SET DEFAULT nextval('public.energy_records_id_seq'::regclass);


--
-- Name: equ_equipment id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_equipment ALTER COLUMN id SET DEFAULT nextval('public.equ_equipment_id_seq'::regclass);


--
-- Name: equ_equipment_check id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_equipment_check ALTER COLUMN id SET DEFAULT nextval('public.equ_equipment_check_id_seq'::regclass);


--
-- Name: equ_equipment_document id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_equipment_document ALTER COLUMN id SET DEFAULT nextval('public.equ_equipment_document_id_seq'::regclass);


--
-- Name: equ_equipment_maintenance id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_equipment_maintenance ALTER COLUMN id SET DEFAULT nextval('public.equ_equipment_maintenance_id_seq'::regclass);


--
-- Name: equ_equipment_part id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_equipment_part ALTER COLUMN id SET DEFAULT nextval('public.equ_equipment_part_id_seq'::regclass);


--
-- Name: equ_equipment_repair id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_equipment_repair ALTER COLUMN id SET DEFAULT nextval('public.equ_equipment_repair_id_seq'::regclass);


--
-- Name: equ_gauge id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_gauge ALTER COLUMN id SET DEFAULT nextval('public.equ_gauge_id_seq'::regclass);


--
-- Name: equ_gauge_calibration id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_gauge_calibration ALTER COLUMN id SET DEFAULT nextval('public.equ_gauge_calibration_id_seq'::regclass);


--
-- Name: equ_mold id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_mold ALTER COLUMN id SET DEFAULT nextval('public.equ_mold_id_seq'::regclass);


--
-- Name: equ_mold_maintenance id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_mold_maintenance ALTER COLUMN id SET DEFAULT nextval('public.equ_mold_maintenance_id_seq'::regclass);


--
-- Name: equ_mold_repair id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_mold_repair ALTER COLUMN id SET DEFAULT nextval('public.equ_mold_repair_id_seq'::regclass);


--
-- Name: equ_oee id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_oee ALTER COLUMN id SET DEFAULT nextval('public.equ_oee_id_seq'::regclass);


--
-- Name: equ_oee_record id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_oee_record ALTER COLUMN id SET DEFAULT nextval('public.equ_oee_record_id_seq'::regclass);


--
-- Name: equ_spare_part id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_spare_part ALTER COLUMN id SET DEFAULT nextval('public.equ_spare_part_id_seq'::regclass);


--
-- Name: equ_teep_data id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_teep_data ALTER COLUMN id SET DEFAULT nextval('public.equ_teep_data_id_seq'::regclass);


--
-- Name: fin_payment_request id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fin_payment_request ALTER COLUMN id SET DEFAULT nextval('public.fin_payment_request_id_seq'::regclass);


--
-- Name: fin_purchase_advance id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fin_purchase_advance ALTER COLUMN id SET DEFAULT nextval('public.fin_purchase_advance_id_seq'::regclass);


--
-- Name: fin_purchase_settlement id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fin_purchase_settlement ALTER COLUMN id SET DEFAULT nextval('public.fin_purchase_settlement_id_seq'::regclass);


--
-- Name: fin_purchase_settlement_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fin_purchase_settlement_item ALTER COLUMN id SET DEFAULT nextval('public.fin_purchase_settlement_item_id_seq'::regclass);


--
-- Name: fin_sales_receipt id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fin_sales_receipt ALTER COLUMN id SET DEFAULT nextval('public.fin_sales_receipt_id_seq'::regclass);


--
-- Name: fin_sales_settlement id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fin_sales_settlement ALTER COLUMN id SET DEFAULT nextval('public.fin_sales_settlement_id_seq'::regclass);


--
-- Name: fin_sales_settlement_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fin_sales_settlement_item ALTER COLUMN id SET DEFAULT nextval('public.fin_sales_settlement_item_id_seq'::regclass);


--
-- Name: fin_statement_detail id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fin_statement_detail ALTER COLUMN id SET DEFAULT nextval('public.fin_statement_detail_id_seq'::regclass);


--
-- Name: fin_supplier_statement id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fin_supplier_statement ALTER COLUMN id SET DEFAULT nextval('public.fin_supplier_statement_id_seq'::regclass);


--
-- Name: fqcs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fqcs ALTER COLUMN id SET DEFAULT nextval('public.fqcs_id_seq'::regclass);


--
-- Name: integration_erp_mappings id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.integration_erp_mappings ALTER COLUMN id SET DEFAULT nextval('public.integration_erp_mappings_id_seq'::regclass);


--
-- Name: integration_erp_sync_logs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.integration_erp_sync_logs ALTER COLUMN id SET DEFAULT nextval('public.integration_erp_sync_logs_id_seq'::regclass);


--
-- Name: inventories id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.inventories ALTER COLUMN id SET DEFAULT nextval('public.inventories_id_seq'::regclass);


--
-- Name: inventory_records id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.inventory_records ALTER COLUMN id SET DEFAULT nextval('public.inventory_records_id_seq'::regclass);


--
-- Name: ip_qcs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ip_qcs ALTER COLUMN id SET DEFAULT nextval('public.ip_qcs_id_seq'::regclass);


--
-- Name: iqc_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.iqc_items ALTER COLUMN id SET DEFAULT nextval('public.iqc_items_id_seq'::regclass);


--
-- Name: iqcs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.iqcs ALTER COLUMN id SET DEFAULT nextval('public.iqcs_id_seq'::regclass);


--
-- Name: lab_calibration id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lab_calibration ALTER COLUMN id SET DEFAULT nextval('public.lab_calibration_id_seq'::regclass);


--
-- Name: lab_instrument id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lab_instrument ALTER COLUMN id SET DEFAULT nextval('public.lab_instrument_id_seq'::regclass);


--
-- Name: lab_report id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lab_report ALTER COLUMN id SET DEFAULT nextval('public.lab_report_id_seq'::regclass);


--
-- Name: lab_sample id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lab_sample ALTER COLUMN id SET DEFAULT nextval('public.lab_sample_id_seq'::regclass);


--
-- Name: lab_test_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lab_test_item ALTER COLUMN id SET DEFAULT nextval('public.lab_test_item_id_seq'::regclass);


--
-- Name: locations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.locations ALTER COLUMN id SET DEFAULT nextval('public.locations_id_seq'::regclass);


--
-- Name: login_logs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.login_logs ALTER COLUMN id SET DEFAULT nextval('public.login_logs_id_seq'::regclass);


--
-- Name: material_categories id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.material_categories ALTER COLUMN id SET DEFAULT nextval('public.material_categories_id_seq'::regclass);


--
-- Name: materials id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.materials ALTER COLUMN id SET DEFAULT nextval('public.materials_id_seq'::regclass);


--
-- Name: mdm_bom id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_bom ALTER COLUMN id SET DEFAULT nextval('public.mdm_bom_id_seq'::regclass);


--
-- Name: mdm_bom_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_bom_item ALTER COLUMN id SET DEFAULT nextval('public.mdm_bom_item_id_seq'::regclass);


--
-- Name: mdm_customer id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_customer ALTER COLUMN id SET DEFAULT nextval('public.mdm_customer_id_seq'::regclass);


--
-- Name: mdm_material id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_material ALTER COLUMN id SET DEFAULT nextval('public.mdm_material_id_seq'::regclass);


--
-- Name: mdm_material_category id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_material_category ALTER COLUMN id SET DEFAULT nextval('public.mdm_material_category_id_seq'::regclass);


--
-- Name: mdm_operation id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_operation ALTER COLUMN id SET DEFAULT nextval('public.mdm_operation_id_seq'::regclass);


--
-- Name: mdm_production_line id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_production_line ALTER COLUMN id SET DEFAULT nextval('public.mdm_production_line_id_seq'::regclass);


--
-- Name: mdm_shift id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_shift ALTER COLUMN id SET DEFAULT nextval('public.mdm_shift_id_seq'::regclass);


--
-- Name: mdm_supplier id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_supplier ALTER COLUMN id SET DEFAULT nextval('public.mdm_supplier_id_seq'::regclass);


--
-- Name: mdm_supplier_material id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_supplier_material ALTER COLUMN id SET DEFAULT nextval('public.mdm_supplier_material_id_seq'::regclass);


--
-- Name: mdm_workshop id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_workshop ALTER COLUMN id SET DEFAULT nextval('public.mdm_workshop_id_seq'::regclass);


--
-- Name: mdm_workshop_config id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_workshop_config ALTER COLUMN id SET DEFAULT nextval('public.mdm_workshop_config_id_seq'::regclass);


--
-- Name: mdm_workstation id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_workstation ALTER COLUMN id SET DEFAULT nextval('public.mdm_workstation_id_seq'::regclass);


--
-- Name: mes_code_rule id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_code_rule ALTER COLUMN id SET DEFAULT nextval('public.mes_code_rule_id_seq'::regclass);


--
-- Name: mes_code_rule_record id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_code_rule_record ALTER COLUMN id SET DEFAULT nextval('public.mes_code_rule_record_id_seq'::regclass);


--
-- Name: mes_container_lifecycle id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_container_lifecycle ALTER COLUMN id SET DEFAULT nextval('public.mes_container_lifecycle_id_seq'::regclass);


--
-- Name: mes_container_maintenance id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_container_maintenance ALTER COLUMN id SET DEFAULT nextval('public.mes_container_maintenance_id_seq'::regclass);


--
-- Name: mes_electronic_sop id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_electronic_sop ALTER COLUMN id SET DEFAULT nextval('public.mes_electronic_sop_id_seq'::regclass);


--
-- Name: mes_first_last_inspect id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_first_last_inspect ALTER COLUMN id SET DEFAULT nextval('public.mes_first_last_inspect_id_seq'::regclass);


--
-- Name: mes_flow_card id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_flow_card ALTER COLUMN id SET DEFAULT nextval('public.mes_flow_card_id_seq'::regclass);


--
-- Name: mes_flow_card_detail id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_flow_card_detail ALTER COLUMN id SET DEFAULT nextval('public.mes_flow_card_detail_id_seq'::regclass);


--
-- Name: mes_job_report_log id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_job_report_log ALTER COLUMN id SET DEFAULT nextval('public.mes_job_report_log_id_seq'::regclass);


--
-- Name: mes_order_day id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_order_day ALTER COLUMN id SET DEFAULT nextval('public.mes_order_day_id_seq'::regclass);


--
-- Name: mes_order_day_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_order_day_item ALTER COLUMN id SET DEFAULT nextval('public.mes_order_day_item_id_seq'::regclass);


--
-- Name: mes_order_month id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_order_month ALTER COLUMN id SET DEFAULT nextval('public.mes_order_month_id_seq'::regclass);


--
-- Name: mes_order_month_audit id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_order_month_audit ALTER COLUMN id SET DEFAULT nextval('public.mes_order_month_audit_id_seq'::regclass);


--
-- Name: mes_order_month_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_order_month_item ALTER COLUMN id SET DEFAULT nextval('public.mes_order_month_item_id_seq'::regclass);


--
-- Name: mes_package id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_package ALTER COLUMN id SET DEFAULT nextval('public.mes_package_id_seq'::regclass);


--
-- Name: mes_person_skill id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_person_skill ALTER COLUMN id SET DEFAULT nextval('public.mes_person_skill_id_seq'::regclass);


--
-- Name: mes_person_skill_score id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_person_skill_score ALTER COLUMN id SET DEFAULT nextval('public.mes_person_skill_score_id_seq'::regclass);


--
-- Name: mes_process id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_process ALTER COLUMN id SET DEFAULT nextval('public.mes_process_id_seq'::regclass);


--
-- Name: mes_process_operation id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_process_operation ALTER COLUMN id SET DEFAULT nextval('public.mes_process_operation_id_seq'::regclass);


--
-- Name: mes_team id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_team ALTER COLUMN id SET DEFAULT nextval('public.mes_team_id_seq'::regclass);


--
-- Name: mes_team_member id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_team_member ALTER COLUMN id SET DEFAULT nextval('public.mes_team_member_id_seq'::regclass);


--
-- Name: mes_team_shift id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_team_shift ALTER COLUMN id SET DEFAULT nextval('public.mes_team_shift_id_seq'::regclass);


--
-- Name: mes_work_scheduling id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_work_scheduling ALTER COLUMN id SET DEFAULT nextval('public.mes_work_scheduling_id_seq'::regclass);


--
-- Name: mes_work_scheduling_detail id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_work_scheduling_detail ALTER COLUMN id SET DEFAULT nextval('public.mes_work_scheduling_detail_id_seq'::regclass);


--
-- Name: mps id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mps ALTER COLUMN id SET DEFAULT nextval('public.mps_id_seq'::regclass);


--
-- Name: mrp_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mrp_items ALTER COLUMN id SET DEFAULT nextval('public.mrp_items_id_seq'::regclass);


--
-- Name: mrps id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mrps ALTER COLUMN id SET DEFAULT nextval('public.mrps_id_seq'::regclass);


--
-- Name: ncrs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ncrs ALTER COLUMN id SET DEFAULT nextval('public.ncrs_id_seq'::regclass);


--
-- Name: notification_channel id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notification_channel ALTER COLUMN id SET DEFAULT nextval('public.notification_channel_id_seq'::regclass);


--
-- Name: oee_report id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.oee_report ALTER COLUMN id SET DEFAULT nextval('public.oee_report_id_seq'::regclass);


--
-- Name: oper_logs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.oper_logs ALTER COLUMN id SET DEFAULT nextval('public.oper_logs_id_seq'::regclass);


--
-- Name: oqcs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.oqcs ALTER COLUMN id SET DEFAULT nextval('public.oqcs_id_seq'::regclass);


--
-- Name: pro_dispatch id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pro_dispatch ALTER COLUMN id SET DEFAULT nextval('public.pro_dispatch_id_seq'::regclass);


--
-- Name: pro_production_order id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pro_production_order ALTER COLUMN id SET DEFAULT nextval('public.pro_production_order_id_seq'::regclass);


--
-- Name: pro_production_report id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pro_production_report ALTER COLUMN id SET DEFAULT nextval('public.pro_production_report_id_seq'::regclass);


--
-- Name: pro_sales_order id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pro_sales_order ALTER COLUMN id SET DEFAULT nextval('public.pro_sales_order_id_seq'::regclass);


--
-- Name: pro_sales_order_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pro_sales_order_item ALTER COLUMN id SET DEFAULT nextval('public.pro_sales_order_item_id_seq'::regclass);


--
-- Name: processes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.processes ALTER COLUMN id SET DEFAULT nextval('public.processes_id_seq'::regclass);


--
-- Name: production_daily_report id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.production_daily_report ALTER COLUMN id SET DEFAULT nextval('public.production_daily_report_id_seq'::regclass);


--
-- Name: production_lines id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.production_lines ALTER COLUMN id SET DEFAULT nextval('public.production_lines_id_seq'::regclass);


--
-- Name: production_offline id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.production_offline ALTER COLUMN id SET DEFAULT nextval('public.production_offline_id_seq'::regclass);


--
-- Name: production_offline_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.production_offline_item ALTER COLUMN id SET DEFAULT nextval('public.production_offline_item_id_seq'::regclass);


--
-- Name: production_orders id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.production_orders ALTER COLUMN id SET DEFAULT nextval('public.production_orders_id_seq'::regclass);


--
-- Name: production_reports id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.production_reports ALTER COLUMN id SET DEFAULT nextval('public.production_reports_id_seq'::regclass);


--
-- Name: qc_aql_level id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_aql_level ALTER COLUMN id SET DEFAULT nextval('public.qc_aql_level_id_seq'::regclass);


--
-- Name: qc_aql_table_row id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_aql_table_row ALTER COLUMN id SET DEFAULT nextval('public.qc_aql_table_row_id_seq'::regclass);


--
-- Name: qc_defect_code id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_defect_code ALTER COLUMN id SET DEFAULT nextval('public.qc_defect_code_id_seq'::regclass);


--
-- Name: qc_defect_record id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_defect_record ALTER COLUMN id SET DEFAULT nextval('public.qc_defect_record_id_seq'::regclass);


--
-- Name: qc_fqc id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_fqc ALTER COLUMN id SET DEFAULT nextval('public.qc_fqc_id_seq'::regclass);


--
-- Name: qc_inspection_characteristic id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_inspection_characteristic ALTER COLUMN id SET DEFAULT nextval('public.qc_inspection_characteristic_id_seq'::regclass);


--
-- Name: qc_inspection_feature id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_inspection_feature ALTER COLUMN id SET DEFAULT nextval('public.qc_inspection_feature_id_seq'::regclass);


--
-- Name: qc_ipqc id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_ipqc ALTER COLUMN id SET DEFAULT nextval('public.qc_ipqc_id_seq'::regclass);


--
-- Name: qc_iqc id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_iqc ALTER COLUMN id SET DEFAULT nextval('public.qc_iqc_id_seq'::regclass);


--
-- Name: qc_iqc_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_iqc_item ALTER COLUMN id SET DEFAULT nextval('public.qc_iqc_item_id_seq'::regclass);


--
-- Name: qc_ncr id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_ncr ALTER COLUMN id SET DEFAULT nextval('public.qc_ncr_id_seq'::regclass);


--
-- Name: qc_oqc id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_oqc ALTER COLUMN id SET DEFAULT nextval('public.qc_oqc_id_seq'::regclass);


--
-- Name: qc_sampling_plan id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_sampling_plan ALTER COLUMN id SET DEFAULT nextval('public.qc_sampling_plan_id_seq'::regclass);


--
-- Name: qc_spc_data id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_spc_data ALTER COLUMN id SET DEFAULT nextval('public.qc_spc_data_id_seq'::regclass);


--
-- Name: qms_sampling_plan id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qms_sampling_plan ALTER COLUMN id SET DEFAULT nextval('public.qms_sampling_plan_id_seq'::regclass);


--
-- Name: qms_sampling_record id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qms_sampling_record ALTER COLUMN id SET DEFAULT nextval('public.qms_sampling_record_id_seq'::regclass);


--
-- Name: qms_sampling_rule id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qms_sampling_rule ALTER COLUMN id SET DEFAULT nextval('public.qms_sampling_rule_id_seq'::regclass);


--
-- Name: quality_weekly_report id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.quality_weekly_report ALTER COLUMN id SET DEFAULT nextval('public.quality_weekly_report_id_seq'::regclass);


--
-- Name: receive_order_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.receive_order_items ALTER COLUMN id SET DEFAULT nextval('public.receive_order_items_id_seq'::regclass);


--
-- Name: receive_orders id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.receive_orders ALTER COLUMN id SET DEFAULT nextval('public.receive_orders_id_seq'::regclass);


--
-- Name: resources id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.resources ALTER COLUMN id SET DEFAULT nextval('public.resources_id_seq'::regclass);


--
-- Name: route_operations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.route_operations ALTER COLUMN id SET DEFAULT nextval('public.route_operations_id_seq'::regclass);


--
-- Name: routes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.routes ALTER COLUMN id SET DEFAULT nextval('public.routes_id_seq'::regclass);


--
-- Name: sales_order_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sales_order_items ALTER COLUMN id SET DEFAULT nextval('public.sales_order_items_id_seq'::regclass);


--
-- Name: sales_orders id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sales_orders ALTER COLUMN id SET DEFAULT nextval('public.sales_orders_id_seq'::regclass);


--
-- Name: schedule_plans id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedule_plans ALTER COLUMN id SET DEFAULT nextval('public.schedule_plans_id_seq'::regclass);


--
-- Name: schedule_results id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedule_results ALTER COLUMN id SET DEFAULT nextval('public.schedule_results_id_seq'::regclass);


--
-- Name: scp_customer_inquiry id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_customer_inquiry ALTER COLUMN id SET DEFAULT nextval('public.scp_customer_inquiry_id_seq'::regclass);


--
-- Name: scp_inquiry_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_inquiry_item ALTER COLUMN id SET DEFAULT nextval('public.scp_inquiry_item_id_seq'::regclass);


--
-- Name: scp_mrs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_mrs ALTER COLUMN id SET DEFAULT nextval('public.scp_mrs_id_seq'::regclass);


--
-- Name: scp_mrs_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_mrs_item ALTER COLUMN id SET DEFAULT nextval('public.scp_mrs_item_id_seq'::regclass);


--
-- Name: scp_po_change_log id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_po_change_log ALTER COLUMN id SET DEFAULT nextval('public.scp_po_change_log_id_seq'::regclass);


--
-- Name: scp_purchase_order id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_purchase_order ALTER COLUMN id SET DEFAULT nextval('public.scp_purchase_order_id_seq'::regclass);


--
-- Name: scp_purchase_order_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_purchase_order_item ALTER COLUMN id SET DEFAULT nextval('public.scp_purchase_order_item_id_seq'::regclass);


--
-- Name: scp_purchase_plan id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_purchase_plan ALTER COLUMN id SET DEFAULT nextval('public.scp_purchase_plan_id_seq'::regclass);


--
-- Name: scp_purchase_plan_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_purchase_plan_item ALTER COLUMN id SET DEFAULT nextval('public.scp_purchase_plan_item_id_seq'::regclass);


--
-- Name: scp_qad_sync_log id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_qad_sync_log ALTER COLUMN id SET DEFAULT nextval('public.scp_qad_sync_log_id_seq'::regclass);


--
-- Name: scp_quote_comparison id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_quote_comparison ALTER COLUMN id SET DEFAULT nextval('public.scp_quote_comparison_id_seq'::regclass);


--
-- Name: scp_quote_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_quote_item ALTER COLUMN id SET DEFAULT nextval('public.scp_quote_item_id_seq'::regclass);


--
-- Name: scp_rfq id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_rfq ALTER COLUMN id SET DEFAULT nextval('public.scp_rfq_id_seq'::regclass);


--
-- Name: scp_rfq_invite id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_rfq_invite ALTER COLUMN id SET DEFAULT nextval('public.scp_rfq_invite_id_seq'::regclass);


--
-- Name: scp_rfq_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_rfq_item ALTER COLUMN id SET DEFAULT nextval('public.scp_rfq_item_id_seq'::regclass);


--
-- Name: scp_sales_order id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_sales_order ALTER COLUMN id SET DEFAULT nextval('public.scp_sales_order_id_seq'::regclass);


--
-- Name: scp_sales_order_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_sales_order_item ALTER COLUMN id SET DEFAULT nextval('public.scp_sales_order_item_id_seq'::regclass);


--
-- Name: scp_so_change_log id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_so_change_log ALTER COLUMN id SET DEFAULT nextval('public.scp_so_change_log_id_seq'::regclass);


--
-- Name: scp_supplier_bank id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_supplier_bank ALTER COLUMN id SET DEFAULT nextval('public.scp_supplier_bank_id_seq'::regclass);


--
-- Name: scp_supplier_contact id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_supplier_contact ALTER COLUMN id SET DEFAULT nextval('public.scp_supplier_contact_id_seq'::regclass);


--
-- Name: scp_supplier_delivery_record id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_supplier_delivery_record ALTER COLUMN id SET DEFAULT nextval('public.scp_supplier_delivery_record_id_seq'::regclass);


--
-- Name: scp_supplier_grade_standard id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_supplier_grade_standard ALTER COLUMN id SET DEFAULT nextval('public.scp_supplier_grade_standard_id_seq'::regclass);


--
-- Name: scp_supplier_kpi id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_supplier_kpi ALTER COLUMN id SET DEFAULT nextval('public.scp_supplier_kpi_id_seq'::regclass);


--
-- Name: scp_supplier_purchase_info id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_supplier_purchase_info ALTER COLUMN id SET DEFAULT nextval('public.scp_supplier_purchase_info_id_seq'::regclass);


--
-- Name: scp_supplier_quality_record id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_supplier_quality_record ALTER COLUMN id SET DEFAULT nextval('public.scp_supplier_quality_record_id_seq'::regclass);


--
-- Name: scp_supplier_quote id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_supplier_quote ALTER COLUMN id SET DEFAULT nextval('public.scp_supplier_quote_id_seq'::regclass);


--
-- Name: serial_numbers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.serial_numbers ALTER COLUMN id SET DEFAULT nextval('public.serial_numbers_id_seq'::regclass);


--
-- Name: shifts id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shifts ALTER COLUMN id SET DEFAULT nextval('public.shifts_id_seq'::regclass);


--
-- Name: spc_data id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.spc_data ALTER COLUMN id SET DEFAULT nextval('public.spc_data_id_seq'::regclass);


--
-- Name: stock_checks id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stock_checks ALTER COLUMN id SET DEFAULT nextval('public.stock_checks_id_seq'::regclass);


--
-- Name: supplier_asn_items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.supplier_asn_items ALTER COLUMN id SET DEFAULT nextval('public.supplier_asn_items_id_seq'::regclass);


--
-- Name: supplier_asns id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.supplier_asns ALTER COLUMN id SET DEFAULT nextval('public.supplier_asns_id_seq'::regclass);


--
-- Name: sys_dept id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_dept ALTER COLUMN id SET DEFAULT nextval('public.sys_dept_id_seq'::regclass);


--
-- Name: sys_dict_data id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_dict_data ALTER COLUMN id SET DEFAULT nextval('public.sys_dict_data_id_seq'::regclass);


--
-- Name: sys_dict_type id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_dict_type ALTER COLUMN id SET DEFAULT nextval('public.sys_dict_type_id_seq'::regclass);


--
-- Name: sys_interface_config id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_interface_config ALTER COLUMN id SET DEFAULT nextval('public.sys_interface_config_id_seq'::regclass);


--
-- Name: sys_interface_exec_log id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_interface_exec_log ALTER COLUMN id SET DEFAULT nextval('public.sys_interface_exec_log_id_seq'::regclass);


--
-- Name: sys_interface_field_map id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_interface_field_map ALTER COLUMN id SET DEFAULT nextval('public.sys_interface_field_map_id_seq'::regclass);


--
-- Name: sys_interface_trigger id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_interface_trigger ALTER COLUMN id SET DEFAULT nextval('public.sys_interface_trigger_id_seq'::regclass);


--
-- Name: sys_login_log id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_login_log ALTER COLUMN id SET DEFAULT nextval('public.sys_login_log_id_seq'::regclass);


--
-- Name: sys_menu id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_menu ALTER COLUMN id SET DEFAULT nextval('public.sys_menu_id_seq'::regclass);


--
-- Name: sys_notice id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_notice ALTER COLUMN id SET DEFAULT nextval('public.sys_notice_id_seq'::regclass);


--
-- Name: sys_notice_read_record id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_notice_read_record ALTER COLUMN id SET DEFAULT nextval('public.sys_notice_read_record_id_seq'::regclass);


--
-- Name: sys_oper_log id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_oper_log ALTER COLUMN id SET DEFAULT nextval('public.sys_oper_log_id_seq'::regclass);


--
-- Name: sys_post id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_post ALTER COLUMN id SET DEFAULT nextval('public.sys_post_id_seq'::regclass);


--
-- Name: sys_print_template id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_print_template ALTER COLUMN id SET DEFAULT nextval('public.sys_print_template_id_seq'::regclass);


--
-- Name: sys_role id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_role ALTER COLUMN id SET DEFAULT nextval('public.sys_role_id_seq'::regclass);


--
-- Name: sys_tenant id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_tenant ALTER COLUMN id SET DEFAULT nextval('public.sys_tenant_id_seq'::regclass);


--
-- Name: sys_user id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_user ALTER COLUMN id SET DEFAULT nextval('public.sys_user_id_seq'::regclass);


--
-- Name: tra_andon_call id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tra_andon_call ALTER COLUMN id SET DEFAULT nextval('public.tra_andon_call_id_seq'::regclass);


--
-- Name: tra_data_collection id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tra_data_collection ALTER COLUMN id SET DEFAULT nextval('public.tra_data_collection_id_seq'::regclass);


--
-- Name: tra_serial_number id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tra_serial_number ALTER COLUMN id SET DEFAULT nextval('public.tra_serial_number_id_seq'::regclass);


--
-- Name: tra_trace_record id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tra_trace_record ALTER COLUMN id SET DEFAULT nextval('public.tra_trace_record_id_seq'::regclass);


--
-- Name: trace_records id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trace_records ALTER COLUMN id SET DEFAULT nextval('public.trace_records_id_seq'::regclass);


--
-- Name: transfer_orders id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transfer_orders ALTER COLUMN id SET DEFAULT nextval('public.transfer_orders_id_seq'::regclass);


--
-- Name: visual_inspection_result id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.visual_inspection_result ALTER COLUMN id SET DEFAULT nextval('public.visual_inspection_result_id_seq'::regclass);


--
-- Name: visual_inspection_task id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.visual_inspection_task ALTER COLUMN id SET DEFAULT nextval('public.visual_inspection_task_id_seq'::regclass);


--
-- Name: warehouses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.warehouses ALTER COLUMN id SET DEFAULT nextval('public.warehouses_id_seq'::regclass);


--
-- Name: wms_agv_device id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_agv_device ALTER COLUMN id SET DEFAULT nextval('public.wms_agv_device_id_seq'::regclass);


--
-- Name: wms_agv_location_mapping id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_agv_location_mapping ALTER COLUMN id SET DEFAULT nextval('public.wms_agv_location_mapping_id_seq'::regclass);


--
-- Name: wms_agv_task id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_agv_task ALTER COLUMN id SET DEFAULT nextval('public.wms_agv_task_id_seq'::regclass);


--
-- Name: wms_delivery_order id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_delivery_order ALTER COLUMN id SET DEFAULT nextval('public.wms_delivery_order_id_seq'::regclass);


--
-- Name: wms_delivery_order_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_delivery_order_item ALTER COLUMN id SET DEFAULT nextval('public.wms_delivery_order_item_id_seq'::regclass);


--
-- Name: wms_inventory id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_inventory ALTER COLUMN id SET DEFAULT nextval('public.wms_inventory_id_seq'::regclass);


--
-- Name: wms_inventory_record id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_inventory_record ALTER COLUMN id SET DEFAULT nextval('public.wms_inventory_record_id_seq'::regclass);


--
-- Name: wms_kanban_pull id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_kanban_pull ALTER COLUMN id SET DEFAULT nextval('public.wms_kanban_pull_id_seq'::regclass);


--
-- Name: wms_location id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_location ALTER COLUMN id SET DEFAULT nextval('public.wms_location_id_seq'::regclass);


--
-- Name: wms_production_complete id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_complete ALTER COLUMN id SET DEFAULT nextval('public.wms_production_complete_id_seq'::regclass);


--
-- Name: wms_production_complete_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_complete_item ALTER COLUMN id SET DEFAULT nextval('public.wms_production_complete_item_id_seq'::regclass);


--
-- Name: wms_production_issue id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_issue ALTER COLUMN id SET DEFAULT nextval('public.wms_production_issue_id_seq'::regclass);


--
-- Name: wms_production_issue_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_issue_item ALTER COLUMN id SET DEFAULT nextval('public.wms_production_issue_item_id_seq'::regclass);


--
-- Name: wms_production_return id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_return ALTER COLUMN id SET DEFAULT nextval('public.wms_production_return_id_seq'::regclass);


--
-- Name: wms_production_return_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_return_item ALTER COLUMN id SET DEFAULT nextval('public.wms_production_return_item_id_seq'::regclass);


--
-- Name: wms_production_stock_in id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_stock_in ALTER COLUMN id SET DEFAULT nextval('public.wms_production_stock_in_id_seq'::regclass);


--
-- Name: wms_production_stock_in_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_stock_in_item ALTER COLUMN id SET DEFAULT nextval('public.wms_production_stock_in_item_id_seq'::regclass);


--
-- Name: wms_purchase_return id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_purchase_return ALTER COLUMN id SET DEFAULT nextval('public.wms_purchase_return_id_seq'::regclass);


--
-- Name: wms_purchase_return_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_purchase_return_item ALTER COLUMN id SET DEFAULT nextval('public.wms_purchase_return_item_id_seq'::regclass);


--
-- Name: wms_receive_order id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_receive_order ALTER COLUMN id SET DEFAULT nextval('public.wms_receive_order_id_seq'::regclass);


--
-- Name: wms_receive_order_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_receive_order_item ALTER COLUMN id SET DEFAULT nextval('public.wms_receive_order_item_id_seq'::regclass);


--
-- Name: wms_sales_return id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_sales_return ALTER COLUMN id SET DEFAULT nextval('public.wms_sales_return_id_seq'::regclass);


--
-- Name: wms_sales_return_item id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_sales_return_item ALTER COLUMN id SET DEFAULT nextval('public.wms_sales_return_item_id_seq'::regclass);


--
-- Name: wms_side_location id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_side_location ALTER COLUMN id SET DEFAULT nextval('public.wms_side_location_id_seq'::regclass);


--
-- Name: wms_stock_check id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_stock_check ALTER COLUMN id SET DEFAULT nextval('public.wms_stock_check_id_seq'::regclass);


--
-- Name: wms_transfer_order id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_transfer_order ALTER COLUMN id SET DEFAULT nextval('public.wms_transfer_order_id_seq'::regclass);


--
-- Name: wms_warehouse id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_warehouse ALTER COLUMN id SET DEFAULT nextval('public.wms_warehouse_id_seq'::regclass);


--
-- Name: work_centers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.work_centers ALTER COLUMN id SET DEFAULT nextval('public.work_centers_id_seq'::regclass);


--
-- Name: workshops id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workshops ALTER COLUMN id SET DEFAULT nextval('public.workshops_id_seq'::regclass);


--
-- Name: workstations id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workstations ALTER COLUMN id SET DEFAULT nextval('public.workstations_id_seq'::regclass);


--
-- Data for Name: ai_chat_conversation; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ai_chat_conversation (id, tenant_id, user_id, session_id, title, created_at, updated_at, deleted_at) FROM stdin;
1	1	1	173206ce-a3a6-47ca-b8c5-a25ad438a9d9	\N	2026-04-08 16:13:22.645159+08	2026-04-08 16:13:22.645159+08	\N
2	1	1	c4dbcd7f-6967-4c84-8fb5-4da440bd4245	\N	2026-04-08 17:56:39.57677+08	2026-04-08 17:56:39.57677+08	\N
3	1	1	ea6c8408-7bfa-476e-8030-f487871c0580	\N	2026-04-08 17:56:55.422117+08	2026-04-08 17:56:55.422117+08	\N
4	1	1	4e8ef3fe-e868-4d05-84ab-89323165d45a	\N	2026-04-08 18:38:35.274165+08	2026-04-08 18:38:35.274165+08	\N
5	1	1	49e6a4aa-2e26-4651-93ea-4303391c2ad8	\N	2026-04-08 18:39:43.493033+08	2026-04-08 18:39:43.493033+08	\N
6	1	1	29faccf3-a8e9-4640-b0bf-7b19caa5a37c	\N	2026-04-08 18:40:04.048907+08	2026-04-08 18:40:04.048907+08	\N
7	1	1	de6ca72f-5201-4c27-9f40-73101df23bff	\N	2026-04-09 08:23:21.385435+08	2026-04-09 08:23:21.385435+08	\N
8	1	1	70db1561-1d04-43ab-96d8-4e38e7446d5f	\N	2026-04-09 08:26:38.400607+08	2026-04-09 08:26:38.400607+08	\N
9	1	1	a2de2e8d-3def-4ca6-b36d-b348757744ab	\N	2026-04-09 08:26:56.329931+08	2026-04-09 08:26:56.329931+08	\N
10	1	1	8e180d76-e23f-402a-8d0f-6f064be06404	\N	2026-04-09 08:27:55.599661+08	2026-04-09 08:27:55.599661+08	\N
11	1	1	9d5cbc89-875b-4786-b230-11387367c26b	\N	2026-04-09 08:30:08.827808+08	2026-04-09 08:30:08.827808+08	\N
12	1	1	37885dd8-a2c2-42fe-89a2-75d38c8390da	\N	2026-04-09 08:30:29.881742+08	2026-04-09 08:30:29.881742+08	\N
13	1	1	aca1b50d-550c-4773-9010-6aab363b4359	\N	2026-04-09 08:31:39.525628+08	2026-04-09 08:31:39.525628+08	\N
14	1	1	9f63ceb8-7c92-4b20-a701-509784e0d3f3	\N	2026-04-09 08:31:54.084672+08	2026-04-09 08:31:54.084672+08	\N
15	1	1	1f5d9bcb-cf31-4ccd-a0ff-2f647e41e878	\N	2026-04-09 08:33:37.44504+08	2026-04-09 08:33:37.44504+08	\N
16	1	1	b54b591d-72e3-4b20-a175-bae2b8635671	\N	2026-04-09 08:34:49.040447+08	2026-04-09 08:34:49.040447+08	\N
17	1	1	cc6fc24e-f0d1-49d7-bbdb-0ab85107ddb5	\N	2026-04-09 08:37:50.711281+08	2026-04-09 08:37:50.711281+08	\N
18	1	1	5c75dbcb-afc3-4910-a287-509dbd0dbabe	\N	2026-04-09 08:39:07.621034+08	2026-04-09 08:39:07.621034+08	\N
19	1	1	2c8bf6d9-9726-4b23-a633-6701bcd7e925	\N	2026-04-09 08:41:57.137234+08	2026-04-09 08:41:57.137234+08	\N
20	1	1	218904e9-f2e0-408d-a0a6-98a4878b5c56	\N	2026-04-09 08:44:21.331323+08	2026-04-09 08:44:21.331323+08	\N
21	1	1	444e25cc-07a3-44f3-931f-7920645665d5	\N	2026-04-09 08:49:27.826149+08	2026-04-09 08:49:27.826149+08	\N
22	1	1	10ea32c8-33d2-4f1f-b07a-488842161b74	\N	2026-04-09 08:54:48.505924+08	2026-04-09 08:54:48.505924+08	\N
23	1	1	8d8d5142-97c3-433d-985a-420f306b61a4	\N	2026-04-09 08:55:57.74313+08	2026-04-09 08:55:57.74313+08	\N
24	1	1	5fd8103d-3cce-4e60-9b8b-dd2b69911530	\N	2026-04-09 08:56:32.724272+08	2026-04-09 08:56:32.724272+08	\N
25	1	1	ee7223a9-f8fa-417e-997b-a91ed40960d8	\N	2026-04-09 08:56:50.160448+08	2026-04-09 08:56:50.160448+08	\N
26	1	1	81cb7ef9-7dc9-47ef-a665-3615b945cb01	\N	2026-04-09 08:58:16.790562+08	2026-04-09 08:58:16.790562+08	\N
27	1	1	964e1072-db97-446c-be22-0dd416eaccf5	\N	2026-04-09 09:00:55.390129+08	2026-04-09 09:00:55.390129+08	\N
28	1	1	d6720b15-3a58-411b-8f09-c2dfb5c4fe24	\N	2026-04-09 09:02:00.912904+08	2026-04-09 09:02:00.912904+08	\N
29	1	1	1f16fff9-a034-4a8f-b83e-f0c07b1a1436	\N	2026-04-09 09:04:03.234719+08	2026-04-09 09:04:03.234719+08	\N
30	1	1	ec33b25d-359c-4b62-b189-2f6597396202	\N	2026-04-09 09:04:51.211222+08	2026-04-09 09:04:51.211222+08	\N
31	1	1	fc59b208-5ba2-4805-9f44-c587e2fc13b7	\N	2026-04-09 09:05:21.347735+08	2026-04-09 09:05:21.347735+08	\N
32	1	1	bcc272be-0b4d-4316-9602-a2aade12d9b7	\N	2026-04-09 09:06:23.738449+08	2026-04-09 09:06:23.738449+08	\N
33	1	1	2cb10847-92fc-47dc-bbca-6fa2681a2be0	\N	2026-04-09 09:13:35.451253+08	2026-04-09 09:13:35.451253+08	\N
34	1	1	12d6574d-62da-4fdd-ab71-185143e9ef1b	\N	2026-04-09 09:14:35.561632+08	2026-04-09 09:14:35.561632+08	\N
35	1	1	455842f9-ba6c-4aa9-85d0-4377a93db231	\N	2026-04-09 09:17:46.477127+08	2026-04-09 09:17:46.477127+08	\N
36	1	1	8a990beb-a056-459a-86b4-8ca216013b82	\N	2026-04-09 14:15:34.699455+08	2026-04-09 14:15:34.699455+08	\N
37	1	1	edda4518-553f-41f9-8d61-be2179484760	\N	2026-04-09 15:08:10.436046+08	2026-04-09 15:08:10.436046+08	\N
38	1	1	0fe522ff-c911-4264-b824-475bfde12769	\N	2026-04-10 08:31:17.391496+08	2026-04-10 08:31:17.391496+08	\N
\.


--
-- Data for Name: ai_chat_message; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ai_chat_message (id, tenant_id, user_id, conversation_id, role, content, intent_json, operation_type, status, tool_result, created_at) FROM stdin;
1	1	1	1	user	查询用户	\N	\N	pending	\N	2026-04-08 16:13:22.654628+08
2	1	1	2	user	nihao	\N	\N	pending	\N	2026-04-08 17:56:39.580499+08
3	1	1	3	user	11	\N	\N	pending	\N	2026-04-08 17:56:55.425+08
4	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:03:05.470184+08
5	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:05:05.589587+08
6	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:07:54.70991+08
7	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:09:38.389338+08
8	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:10:26.647732+08
9	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:11:45.098694+08
10	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:14:15.928874+08
11	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:15:31.897494+08
12	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:16:36.777942+08
13	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:17:46.407341+08
14	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:18:35.686583+08
15	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:19:52.077155+08
16	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:21:13.76781+08
17	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:21:37.035047+08
18	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:23:06.339999+08
19	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:23:25.072054+08
20	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:25:40.898479+08
21	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:27:14.706024+08
22	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:27:33.260005+08
23	1	1	1	user	Hello	\N	\N	pending	\N	2026-04-08 18:28:46.334571+08
24	1	1	4	user	hello	\N	\N	pending	\N	2026-04-08 18:38:35.275679+08
25	1	1	5	user	hello	\N	\N	pending	\N	2026-04-08 18:39:43.496263+08
26	1	1	6	user	hello	\N	\N	pending	\N	2026-04-08 18:40:04.05043+08
27	1	1	7	user	hello	\N	\N	pending	\N	2026-04-09 08:23:21.388156+08
28	1	1	8	user	hello	\N	\N	pending	\N	2026-04-09 08:26:38.401686+08
29	1	1	9	user	hello	\N	\N	pending	\N	2026-04-09 08:26:56.331135+08
30	1	1	10	user	hello	\N	\N	pending	\N	2026-04-09 08:27:55.60021+08
31	1	1	11	user	hello	\N	\N	pending	\N	2026-04-09 08:30:08.829481+08
32	1	1	12	user	hello	\N	\N	pending	\N	2026-04-09 08:30:29.883035+08
33	1	1	13	user	hello	\N	\N	pending	\N	2026-04-09 08:31:39.528695+08
34	1	1	14	user	hello	\N	\N	pending	\N	2026-04-09 08:31:54.085403+08
35	1	1	15	user	hello	\N	\N	pending	\N	2026-04-09 08:33:37.446084+08
36	1	1	16	user	hello	\N	\N	pending	\N	2026-04-09 08:34:49.043326+08
37	1	1	17	user	hello	\N	\N	pending	\N	2026-04-09 08:37:50.712633+08
38	1	1	18	user	hello	\N	\N	pending	\N	2026-04-09 08:39:07.623018+08
39	1	1	19	user	hello	\N	\N	pending	\N	2026-04-09 08:41:57.138813+08
40	1	1	19	assistant	<think>\n用户发送了一个简单的问候"hello"。这是一个非结构化的自然语言输入，不包含任何明确的操作意图或请求。\n\n根据操作分类，这应该被归类为"unknown"操作，因为：\n1. 它不包含任何查询、写入、分析或导航意图\n2. 它只是一个简单的问候语\n\n我应该以JSON格式回复，说明这是一个未知的操作，并给出友好的回应。\n</think>\n\n```json\n{\n  "operation_type": "unknown",\n  "module": null,\n  "action": null,\n  "parameters": {},\n  "natural_desc": "用户打招呼，系统无法解析为具体操作",\n  "confidence": 0.0,\n  "message": "您好！我是MOM3.0智能助手，我可以帮您处理制造业执行系统相关的操作。请告诉我您需要什么帮助，例如：\\n- 查询库存、仓库、库位\\n- 创建或查询检验单（IQC/OQC/FQC/IPQC）\\n- 设备OEE分析\\n- 生产工单管理\\n- 主生产计划（MPS）查询\\n- 物料主数据管理\\n- 产能分析\\n\\n请问有什么可以帮您？"\n}\n```	{"operation_type":"unknown","module":"","action":"","parameters":{},"natural_desc":"用户打招呼，系统无法解析为具体操作","confidence":0}	unknown	pending	\N	2026-04-09 08:42:05.063365+08
41	1	1	20	user	1	\N	\N	pending	\N	2026-04-09 08:44:21.332538+08
42	1	1	20	assistant	<think>\n用户只输入了"1"，这是一个非常简短的输入，没有明确的意图。我需要将其归类为unknown操作，因为无法从中解析出有意义的操作请求。\n\n让我按照要求的JSON格式回复：\n- operation_type: unknown（无法识别）\n- module: 无\n- action: 无\n- parameters: 空对象\n- natural_desc: 简单描述输入的内容\n- confidence: 较低，因为信息不足\n</think>\n\n```json\n{\n  "operation_type": "unknown",\n  "module": null,\n  "action": null,\n  "parameters": {},\n  "natural_desc": "用户输入了字符'1'，无法识别有效操作意图。请提供更完整的请求描述。",\n  "confidence": 0.1\n}\n```\n\n---\n\n**💡 提示：** 我可以帮助您执行以下操作：\n\n| 模块 | 示例请求 |\n|------|----------|\n| **生产工单** | "查询今日生产工单"、"创建生产工单"、"开始生产" |\n| **质量管理** | "查询IQC检验列表"、"创建FQC检验单" |\n| **设备管理** | "查询设备列表"、"查看OEE分析" |\n| **物料管理** | "查询物料列表"、"查看物料详情" |\n| **仓储管理** | "查询库存"、"查看仓库列表" |\n| **APS排程** | "查看排程计划"、"产能分析" |\n\n请告诉我您需要执行的操作，例如：**"查询今天有哪些生产工单"**	{"operation_type":"unknown","module":"","action":"","parameters":{},"natural_desc":"用户输入了字符'1'，无法识别有效操作意图。请提供更完整的请求描述。","confidence":0.1}	unknown	pending	\N	2026-04-09 08:44:31.876994+08
43	1	1	20	user	你好，你是谁	\N	\N	pending	\N	2026-04-09 08:44:55.979542+08
66	1	1	31	user	hi	\N	\N	pending	\N	2026-04-09 09:05:21.348835+08
67	1	1	32	user	hi	\N	\N	pending	\N	2026-04-09 09:06:23.741348+08
44	1	1	20	assistant	<think>\n用户问"你好，你是谁"，这是一个简单的自我介绍问题，不是一个操作请求。我应该友好地回复并介绍自己。\n</think>\n\n```json\n{\n  "operation_type": "unknown",\n  "module": null,\n  "action": null,\n  "parameters": {},\n  "natural_desc": "用户询问身份，非系统操作请求",\n  "confidence": 1.0\n}\n```\n\n---\n\n您好！👋\n\n我是 **MOM3.0 智能助手**，一个专业的**制造业执行系统（Manufacturing Operations Management）** AI 助手。\n\n### 我能帮您做什么？\n\n| 模块 | 功能说明 |\n|------|----------|\n| 🏭 **生产管理** | 查询/创建/管理生产工单，报工统计 |\n| ✅ **质量管理** | IQC、IPQC、FQC、OQC 检验管理 |\n| 🔧 **设备管理** | 设备列表查询、OEE 分析 |\n| 📦 **物料管理** | 物料主数据查询与维护 |\n| 🏬 **仓储管理** | 仓库、库位、库存查询 |\n| 📅 **APS排程** | 主生产计划、排程结果、产能分析 |\n| 👥 **系统管理** | 部门、角色、用户管理 |\n\n### 使用方式\n\n您可以用自然语言告诉我您的需求，例如：\n\n- *"查询今天有哪些生产工单"*\n- *"帮我创建一个IQC检验单"*\n- *"查看设备OEE分析"*\n- *"产能利用率是多少"*\n\n---\n\n请问有什么可以帮您的？ 😊	{"operation_type":"unknown","module":"","action":"","parameters":{},"natural_desc":"用户询问身份，非系统操作请求","confidence":1}	unknown	pending	\N	2026-04-09 08:45:06.092503+08
45	1	1	20	user	汇报今天工单	\N	\N	pending	\N	2026-04-09 08:47:10.636924+08
46	1	1	20	assistant	<think>\n用户说"汇报今天工单"，这是一个生产管理相关的请求。\n\n分析意图：\n- "汇报今天工单" - 这可能是在请求查询今天的生产工单情况\n- 可能是想要获取工单列表或者工单统计信息\n\n相关操作：\n- list_orders: GET /api/v1/production/order/list - 查询生产工单列表\n- order_stats: GET /api/v1/production/order/stats - 工单统计\n\n"汇报"这个动作，我理解是想要获取工单的状态汇总信息，可能需要结合查询列表和统计两个操作。\n\n参数方面：\n- 需要查询今天的工单，日期过滤参数\n\n但用户没有明确说明具体的日期，我默认理解为查询今天的工单。\n\n这是一个 query 操作。\n</think>\n\n```json\n{\n  "operation_type": "query",\n  "module": "production",\n  "action": "list_orders",\n  "parameters": {\n    "date": "2024-01-15"\n  },\n  "natural_desc": "查询今日（2024-01-15）的生产工单列表和状态",\n  "confidence": 0.85\n}\n```\n\n---\n\n正在为您查询今日生产工单情况... 🔄\n\n请问您需要的是：\n\n1. **📋 工单列表** - 查看今天所有工单的基本信息\n2. **📊 工单统计** - 汇总工单数量、完成率、状态分布等\n3. **📋+📊 完整汇报** - 同时展示工单列表和统计汇总\n\n请告诉我您的偏好，或者直接告诉我具体想看什么内容，例如：*"查看正在生产的工单"* 或 *"统计各车间工单完成情况"*。	{"operation_type":"query","module":"production","action":"list_orders","parameters":{"date":"2024-01-15"},"natural_desc":"查询今日（2024-01-15）的生产工单列表和状态","confidence":0.85}	query	pending	\N	2026-04-09 08:47:21.605588+08
47	1	1	20	user	你好，你是什么模型？	\N	\N	pending	\N	2026-04-09 08:48:45.925712+08
48	1	1	20	assistant	<think>\n用户再次询问我的身份/模型信息，这不是系统操作请求，是一个闲聊/问答类型的问题。我应该简单回答，并引导用户回到实际工作场景。\n</think>\n\n```json\n{\n  "operation_type": "unknown",\n  "module": null,\n  "action": null,\n  "parameters": {},\n  "natural_desc": "用户询问AI模型身份，非系统操作",\n  "confidence": 1.0\n}\n```\n\n---\n\n您好！😊\n\n我是 **MOM3.0 智能助手**，专为制造业执行系统设计的 AI 助手。\n\n我的核心能力是帮助您快速完成日常工作：\n\n| 🚀 快捷操作 | 示例 |\n|------------|------|\n| 查询数据 | "查看今天的生产工单" |\n| 执行操作 | "帮我创建一个FQC检验单" |\n| 数据分析 | "分析本月设备OEE" |\n| 业务咨询 | "生产工单流程是什么" |\n\n---\n\n您可以随时用自然语言向我提问，或者告诉我您现在想处理什么工作，我来帮您快速完成！ 🎯	{"operation_type":"unknown","module":"","action":"","parameters":{},"natural_desc":"用户询问AI模型身份，非系统操作","confidence":1}	unknown	pending	\N	2026-04-09 08:48:52.224011+08
49	1	1	21	user	�鿴����2026-04-09����������	\N	\N	pending	\N	2026-04-09 08:49:27.827181+08
50	1	1	21	assistant	<think>\nThe user is asking to query production orders for a specific date (2026-04-09). Let me parse this request.\n\nThe request is in Chinese and says "查看2026-04-09的生产订单" which means "View production orders for 2026-04-09".\n\nThis is a query operation (GET request) for production orders. I need to use the list_orders function from the production module.\n\nThe parameters would be:\n- date: "2026-04-09" or similar parameter to filter by date\n\nLet me check the list_orders API:\n- GET /api/v1/production/order/list - 查询生产工单列表\n\nThe date parameter could be "order_date" or "production_date" or "start_date" etc. I'll need to make a reasonable assumption about the parameter name. Common parameter names for date filtering in order lists would be order_date, start_date, production_date, etc.\n\nSince I don't have the exact API documentation for the parameters, I'll use a common parameter name like "order_date" or "production_date". Let me use "order_date" as a reasonable guess.\n\nActually, I should be more careful. Without knowing the exact parameter names, I should note this uncertainty. But since the user is asking for a specific date, the most likely parameter would be something like "order_date" or "date" or "production_date".\n\nLet me provide the JSON response with reasonable assumptions.\n</think>\n\n```json\n{\n  "operation_type": "query",\n  "module": "production",\n  "action": "list_orders",\n  "parameters": {\n    "order_date": "2026-04-09"\n  },\n  "natural_desc": "查询2026-04-09的生产工单列表",\n  "confidence": 0.8\n}\n```	{"operation_type":"query","module":"production","action":"list_orders","parameters":{"order_date":"2026-04-09"},"natural_desc":"查询2026-04-09的生产工单列表","confidence":0.8}	query	pending	\N	2026-04-09 08:49:36.602206+08
51	1	1	22	user	���������治����	\N	\N	pending	\N	2026-04-09 08:54:48.507705+08
52	1	1	22	assistant	<think>\nThe user sent a message that appears to be garbled/corrupted text (likely encoding issues). It looks like Chinese characters that got corrupted during transmission. I need to interpret this as best as possible.\n\nThe garbled text looks like it could be: "帮我看看订单" (help me check the orders) or something similar. But it's really not clear due to encoding issues.\n\nLet me try to decode:\n- 帮 = help\n- 我 = me/I\n- 看 = look/watch\n- 订单 = order\n\nIt could be: "帮我看订单" or "帮我查订单" or similar.\n\nGiven the garbled nature, I should ask the user to clarify what they meant. Since I cannot confidently determine what they want, I should classify this as "chat" since I don't understand the request clearly enough to execute a system operation.\n\nActually, looking more carefully at the text pattern:\n"��������治����" - this is clearly mojibake (文字化け), which happens when Chinese text is decoded with the wrong encoding (like UTF-8 interpreted as Latin-1 or similar).\n\nI should respond by asking the user to re-send their message or clarify what they need.\n</think>\n\n```json\n{\n  "operation_type": "chat",\n  "module": null,\n  "action": null,\n  "parameters": {},\n  "natural_desc": "收到乱码信息，无法识别用户意图",\n  "confidence": 0.3\n}\n```\n\n---\n\n看起来您的消息出现了乱码（文字编码问题）。请您重新发送一下消息，我帮您处理。\n\n例如：\n- "帮我查询生产工单"\n- "查看库存"\n- "创建物料"	{"operation_type":"unknown","module":"","action":"","parameters":null,"natural_desc":"\\u003cthink\\u003e\\nThe user sent a message that appears to be garbled/corrupted text (likely encoding issues). It looks like Chinese characters that got corrupted during transmission. I need to interpret this as best as possible.\\n\\nThe garbled text looks like it could be: \\"帮我看看订单\\" (help me check the orders) or something similar. But it's really not clear due to encoding issues.\\n\\nLet me try to decode:\\n- 帮 = help\\n- 我 = me/I\\n- 看 = look/watch\\n- 订单 = order\\n\\nIt could be: \\"帮我看订单\\" or \\"帮我查订单\\" or similar.\\n\\nGiven the garbled nature, I should ask the user to clarify what they meant. Since I cannot confidently determine what they want, I should classify this as \\"chat\\" since I don't understand the request clearly enough to execute a system operation.\\n\\nActually, looking more carefully at the text pattern:\\n\\"��������治����\\" - this is clearly mojibake (文字化け), which happens when Chinese text is decoded with the wrong encoding (like UTF-8 interpreted as Latin-1 or similar).\\n\\nI should respond by asking the user to re-send their message or clarify what they need.\\n\\u003c/think\\u003e\\n\\n```json\\n{\\n  \\"operation_type\\": \\"chat\\",\\n  \\"module\\": null,\\n  \\"action\\": null,\\n  \\"parameters\\": {},\\n  \\"natural_desc\\": \\"收到乱码信息，无法识别用户意图\\",\\n  \\"confidence\\": 0.3\\n}\\n```\\n\\n---\\n\\n看起来您的消息出现了乱码（文字编码问题）。请您重新发送一下消息，我帮您处理。\\n\\n例如：\\n- \\"帮我查询生产工单\\"\\n- \\"查看库存\\"\\n- \\"创建物料\\"","confidence":0}	unknown	pending	\N	2026-04-09 08:54:59.742865+08
53	1	1	23	user	Hello, how are you today?	\N	\N	pending	\N	2026-04-09 08:55:57.744336+08
54	1	1	23	assistant	<think>\nThe user is just greeting me and asking how I am. This is a casual conversation without any specific intent to perform system operations. According to the rules, this should be classified as "chat" since they're just chatting/greeting without needing any MOM system operations.\n</think>\n\n```json\n{\n  "operation_type": "chat",\n  "module": null,\n  "action": null,\n  "parameters": {},\n  "natural_desc": "用户问候对话",\n  "confidence": 1.0\n}\n```	{"operation_type":"unknown","module":"","action":"","parameters":null,"natural_desc":"\\u003cthink\\u003e\\nThe user is just greeting me and asking how I am. This is a casual conversation without any specific intent to perform system operations. According to the rules, this should be classified as \\"chat\\" since they're just chatting/greeting without needing any MOM system operations.\\n\\u003c/think\\u003e\\n\\n```json\\n{\\n  \\"operation_type\\": \\"chat\\",\\n  \\"module\\": null,\\n  \\"action\\": null,\\n  \\"parameters\\": {},\\n  \\"natural_desc\\": \\"用户问候对话\\",\\n  \\"confidence\\": 1.0\\n}\\n```","confidence":0}	unknown	pending	\N	2026-04-09 08:56:00.905286+08
55	1	1	24	user	Hello, how are you today?	\N	\N	pending	\N	2026-04-09 08:56:32.725334+08
56	1	1	24	assistant	<think>\nThe user is greeting me and asking how I am. This is a friendly chat/greeting that doesn't require any system operations. According to the rules:\n\n- If the user is just chatting, greeting, or expressing emotions without needing any MOM system operations → chat\n\nSo this should be classified as "chat".\n</think>\n\n{\n  "operation_type": "chat",\n  "module": null,\n  "action": null,\n  "parameters": {},\n  "natural_desc": "用户打招呼问候",\n  "confidence": 1.0\n}	{"operation_type":"unknown","module":"","action":"","parameters":null,"natural_desc":"\\u003cthink\\u003e\\nThe user is greeting me and asking how I am. This is a friendly chat/greeting that doesn't require any system operations. According to the rules:\\n\\n- If the user is just chatting, greeting, or expressing emotions without needing any MOM system operations → chat\\n\\nSo this should be classified as \\"chat\\".\\n\\u003c/think\\u003e\\n\\n{\\n  \\"operation_type\\": \\"chat\\",\\n  \\"module\\": null,\\n  \\"action\\": null,\\n  \\"parameters\\": {},\\n  \\"natural_desc\\": \\"用户打招呼问候\\",\\n  \\"confidence\\": 1.0\\n}","confidence":0}	unknown	pending	\N	2026-04-09 08:56:36.204513+08
57	1	1	25	user	Hello, how are you today?	\N	\N	pending	\N	2026-04-09 08:56:50.162741+08
58	1	1	25	assistant	<think>\n用户只是在问候和聊天，没有要求任何MOM系统操作。根据判断规则：\n- 如果用户只是在聊天、问问题、问候、或表达情绪，但不需要执行任何MOM系统操作 → chat\n\n这明显是一个chat类型的操作。\n</think>\n\n```json\n{\n  "operation_type": "chat",\n  "module": null,\n  "action": null,\n  "parameters": {},\n  "natural_desc": "用户问候对话",\n  "confidence": 1.0\n}\n```	{"operation_type":"unknown","module":"","action":"","parameters":null,"natural_desc":"\\u003cthink\\u003e\\n用户只是在问候和聊天，没有要求任何MOM系统操作。根据判断规则：\\n- 如果用户只是在聊天、问问题、问候、或表达情绪，但不需要执行任何MOM系统操作 → chat\\n\\n这明显是一个chat类型的操作。\\n\\u003c/think\\u003e\\n\\n```json\\n{\\n  \\"operation_type\\": \\"chat\\",\\n  \\"module\\": null,\\n  \\"action\\": null,\\n  \\"parameters\\": {},\\n  \\"natural_desc\\": \\"用户问候对话\\",\\n  \\"confidence\\": 1.0\\n}\\n```","confidence":0}	unknown	pending	\N	2026-04-09 08:56:53.32422+08
59	1	1	26	user	Hello, how are you today?	\N	\N	pending	\N	2026-04-09 08:58:16.792493+08
60	1	1	26	assistant	<think>\nThe user is greeting me with a casual "Hello, how are you today?" This is a simple greeting/chat interaction that doesn't require any system operations. According to the rules:\n\n"If the user is just chatting, asking questions, greeting, or expressing emotions, but doesn't need to perform any MOM system operations → chat"\n\nThis is clearly a chat interaction, not a query or write operation.\n</think>\n\n{\n  "operation_type": "chat",\n  "module": null,\n  "action": null,\n  "parameters": {},\n  "natural_desc": "用户问候，询问我的状态",\n  "confidence": 1.0\n}	{"operation_type":"unknown","module":"","action":"","parameters":null,"natural_desc":"\\u003cthink\\u003e\\nThe user is greeting me with a casual \\"Hello, how are you today?\\" This is a simple greeting/chat interaction that doesn't require any system operations. According to the rules:\\n\\n\\"If the user is just chatting, asking questions, greeting, or expressing emotions, but doesn't need to perform any MOM system operations → chat\\"\\n\\nThis is clearly a chat interaction, not a query or write operation.\\n\\u003c/think\\u003e\\n\\n{\\n  \\"operation_type\\": \\"chat\\",\\n  \\"module\\": null,\\n  \\"action\\": null,\\n  \\"parameters\\": {},\\n  \\"natural_desc\\": \\"用户问候，询问我的状态\\",\\n  \\"confidence\\": 1.0\\n}","confidence":0}	unknown	pending	\N	2026-04-09 08:58:20.964919+08
61	1	1	27	user	Hello, how are you today?	\N	\N	pending	\N	2026-04-09 09:00:55.391974+08
62	1	1	28	user	Hello	\N	\N	pending	\N	2026-04-09 09:02:00.916917+08
63	1	1	28	assistant	<think>\n用户只是说"Hello"，这是一个简单的问候语，不需要执行任何MOM系统操作。根据判断规则，这属于chat类型。\n</think>\n\n{\n  "operation_type": "chat",\n  "module": null,\n  "action": null,\n  "parameters": {},\n  "natural_desc": "用户打招呼问候",\n  "confidence": 0.99\n}	{"operation_type":"unknown","module":"","action":"","parameters":null,"natural_desc":"\\u003cthink\\u003e\\n用户只是说\\"Hello\\"，这是一个简单的问候语，不需要执行任何MOM系统操作。根据判断规则，这属于chat类型。\\n\\u003c/think\\u003e\\n\\n{\\n  \\"operation_type\\": \\"chat\\",\\n  \\"module\\": null,\\n  \\"action\\": null,\\n  \\"parameters\\": {},\\n  \\"natural_desc\\": \\"用户打招呼问候\\",\\n  \\"confidence\\": 0.99\\n}","confidence":0}	unknown	pending	\N	2026-04-09 09:02:07.018283+08
64	1	1	29	user	Hello	\N	\N	pending	\N	2026-04-09 09:04:03.236403+08
65	1	1	30	user	hi	\N	\N	pending	\N	2026-04-09 09:04:51.21383+08
68	1	1	32	assistant	<think>\n用户只是说了"hi"，这是一个简单的问候语。用户没有表达任何需要执行MOM系统操作的意图，只是打招呼或开始对话。\n\n根据判断规则：\n- 如果用户只是在聊天、问问题、问候、或表达情绪，但不需要执行任何MOM系统操作 → chat\n\n所以这应该被归类为chat操作。\n</think>\n\n{\n  "operation_type": "chat",\n  "module": null,\n  "action": null,\n  "parameters": {},\n  "natural_desc": "用户打招呼问候",\n  "confidence": 1.0\n}	{"operation_type":"unknown","module":"","action":"","parameters":null,"natural_desc":"\\u003cthink\\u003e\\n用户只是说了\\"hi\\"，这是一个简单的问候语。用户没有表达任何需要执行MOM系统操作的意图，只是打招呼或开始对话。\\n\\n根据判断规则：\\n- 如果用户只是在聊天、问问题、问候、或表达情绪，但不需要执行任何MOM系统操作 → chat\\n\\n所以这应该被归类为chat操作。\\n\\u003c/think\\u003e\\n\\n{\\n  \\"operation_type\\": \\"chat\\",\\n  \\"module\\": null,\\n  \\"action\\": null,\\n  \\"parameters\\": {},\\n  \\"natural_desc\\": \\"用户打招呼问候\\",\\n  \\"confidence\\": 1.0\\n}","confidence":0}	unknown	pending	\N	2026-04-09 09:06:28.505226+08
69	1	1	33	user	你现在能修改这个项目的代码吗？	\N	\N	pending	\N	2026-04-09 09:13:35.456224+08
70	1	1	34	user	你现在能修改项目的代码吗？	\N	\N	pending	\N	2026-04-09 09:14:35.56268+08
71	1	1	35	user	Uncaught (in promise) Error: 错误: 关系 "mdm_customer" 不存在 (SQLSTATE 42P01)\n    at request.ts:54:27\n    at async loadData (CustomerList.vue:90:72)\n    at async loadData (CustomerList.vue:90:72)	\N	\N	pending	\N	2026-04-09 09:17:46.480422+08
72	1	1	35	assistant	<think>\n用户分享了一个错误信息，这是一个前端错误，来自 CustomerList.vue 文件。这个错误表明在数据库中找不到 "mdm_customer" 表（关系）。\n\n这是一个技术错误/调试问题，不是MOM系统的业务操作请求。用户似乎在报告一个bug或寻求帮助解决这个错误。\n\n根据判断规则：\n- 如果用户只是在聊天、问问题、问候、或表达情绪，但不需要执行任何MOM系统操作 → chat\n- 这个错误是技术问题，用户可能在寻求帮助\n\n我应该询问用户需要什么帮助，这是一个chat类型的交互。\n</think>\n\n{\n  "operation_type": "chat",\n  "action": null,\n  "parameters": {},\n  "natural_desc": "这是一个数据库表不存在的错误。系统找不到 'mdm_customer' 表。这可能是以下原因之一：1) 数据库表未创建；2) 表名拼写错误；3) 数据库迁移未执行；4) 配置文件中的表名设置有误。请问您需要我帮您排查这个问题吗？或者您需要执行其他MOM系统操作？",\n  "confidence": 1.0\n}	{"operation_type":"unknown","module":"","action":"","parameters":null,"natural_desc":"\\u003cthink\\u003e\\n用户分享了一个错误信息，这是一个前端错误，来自 CustomerList.vue 文件。这个错误表明在数据库中找不到 \\"mdm_customer\\" 表（关系）。\\n\\n这是一个技术错误/调试问题，不是MOM系统的业务操作请求。用户似乎在报告一个bug或寻求帮助解决这个错误。\\n\\n根据判断规则：\\n- 如果用户只是在聊天、问问题、问候、或表达情绪，但不需要执行任何MOM系统操作 → chat\\n- 这个错误是技术问题，用户可能在寻求帮助\\n\\n我应该询问用户需要什么帮助，这是一个chat类型的交互。\\n\\u003c/think\\u003e\\n\\n{\\n  \\"operation_type\\": \\"chat\\",\\n  \\"action\\": null,\\n  \\"parameters\\": {},\\n  \\"natural_desc\\": \\"这是一个数据库表不存在的错误。系统找不到 'mdm_customer' 表。这可能是以下原因之一：1) 数据库表未创建；2) 表名拼写错误；3) 数据库迁移未执行；4) 配置文件中的表名设置有误。请问您需要我帮您排查这个问题吗？或者您需要执行其他MOM系统操作？\\",\\n  \\"confidence\\": 1.0\\n}","confidence":0}	unknown	pending	\N	2026-04-09 09:17:57.584034+08
73	1	1	35	user	排查并修复	\N	\N	pending	\N	2026-04-09 09:18:40.970547+08
74	1	1	35	assistant	<think>\n用户要求排查并修复这个数据库表不存在的问题。这不是一个标准的MOM系统API操作，而是一个技术问题排查。\n\n让我分析一下：\n1. 错误信息显示 `mdm_customer` 表不存在\n2. 错误发生在 `CustomerList.vue` 第90行调用 `loadData` 时\n3. 这是一个前端Vue组件调用后端API时遇到的问题\n\n我需要：\n1. 首先了解上下文 - 用户在使用客户管理模块\n2. 可能需要检查数据库表是否存在\n3. 可能需要检查后端API路由配置\n4. 可能需要检查数据库迁移\n\n但从系统功能角度，我应该：\n- 先查询一下相关的模块数据，看看是否能获取信息\n- 提供排查建议\n\n让我先尝试查询客户相关信息，看看是否能触发API调用来进一步诊断。\n\n不过，更实际的做法是，我应该提供一份排查步骤清单，让用户或运维人员去执行：\n\n1. 检查数据库中是否真的没有这个表\n2. 检查后端代码中的表名配置\n3. 检查数据库迁移文件\n4. 如果表确实不存在，需要创建它\n\n让我提供详细的排查和修复方案。\n</think>\n\n{\n  "operation_type": "unknown",\n  "action": null,\n  "parameters": {},\n  "natural_desc": "好的，这是一个数据库表不存在的错误。以下是排查和修复步骤：\\n\\n## 🔍 问题原因\\n表 `mdm_customer` 在数据库中不存在。\\n\\n## 🛠️ 排查步骤\\n\\n**步骤1：确认数据库表状态**\\n```sql\\n-- 连接数据库后执行：\\nSELECT table_name FROM information_schema.tables \\nWHERE table_schema = 'public' AND table_name = 'mdm_customer';\\n```\\n\\n**步骤2：检查后端表名配置**\\n查看后端代码中 `Customer` 实体/模型的表名配置，确认是否配置为 `mdm_customer`。\\n\\n**步骤3：检查数据库迁移文件**\\n查找项目中是否有创建该表的迁移文件。\\n\\n## 🛠️ 修复方案\\n\\n**方案A：如果表完全不存在**\\n需要执行建表SQL（需联系DBA或查看原项目迁移脚本）：\\n\\n```sql\\nCREATE TABLE mdm_customer (\\n    id SERIAL PRIMARY KEY,\\n    code VARCHAR(50) NOT NULL UNIQUE,\\n    name VARCHAR(200) NOT NULL,\\n    -- 根据业务需求添加其他字段\\n    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,\\n    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP\\n);\\n```\\n\\n**方案B：如果表名配置错误**\\n修改后端代码中的表名映射配置。\\n\\n**方案C：运行数据库迁移**\\n如果项目有迁移脚本，执行：\\n```bash\\nnpm run migrate\\n# 或\\nyarn migrate\\n```\\n\\n---\\n\\n请问：\\n1. 您是否有该项目的数据库结构文档或迁移脚本？\\n2. 需要我帮您生成建表SQL语句吗？（需要提供表的完整字段定义）\\n3. 您是否知道这个 `mdm_customer` 表应该包含哪些字段？",\n  "confidence": 0.95\n}	{"operation_type":"unknown","module":"","action":"","parameters":{},"natural_desc":"好的，这是一个数据库表不存在的错误。以下是排查和修复步骤：\\n\\n## 🔍 问题原因\\n表 `mdm_customer` 在数据库中不存在。\\n\\n## 🛠️ 排查步骤\\n\\n**步骤1：确认数据库表状态**\\n```sql\\n-- 连接数据库后执行：\\nSELECT table_name FROM information_schema.tables \\nWHERE table_schema = 'public' AND table_name = 'mdm_customer';\\n```\\n\\n**步骤2：检查后端表名配置**\\n查看后端代码中 `Customer` 实体/模型的表名配置，确认是否配置为 `mdm_customer`。\\n\\n**步骤3：检查数据库迁移文件**\\n查找项目中是否有创建该表的迁移文件。\\n\\n## 🛠️ 修复方案\\n\\n**方案A：如果表完全不存在**\\n需要执行建表SQL（需联系DBA或查看原项目迁移脚本）：\\n\\n```sql\\nCREATE TABLE mdm_customer (\\n    id SERIAL PRIMARY KEY,\\n    code VARCHAR(50) NOT NULL UNIQUE,\\n    name VARCHAR(200) NOT NULL,\\n    -- 根据业务需求添加其他字段\\n    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,\\n    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP\\n);\\n```\\n\\n**方案B：如果表名配置错误**\\n修改后端代码中的表名映射配置。\\n\\n**方案C：运行数据库迁移**\\n如果项目有迁移脚本，执行：\\n```bash\\nnpm run migrate\\n# 或\\nyarn migrate\\n```\\n\\n---\\n\\n请问：\\n1. 您是否有该项目的数据库结构文档或迁移脚本？\\n2. 需要我帮您生成建表SQL语句吗？（需要提供表的完整字段定义）\\n3. 您是否知道这个 `mdm_customer` 表应该包含哪些字段？","confidence":0.95}	unknown	pending	\N	2026-04-09 09:19:04.466332+08
75	1	1	36	user	客户管理没有数据，新增20条测试数据	\N	\N	pending	\N	2026-04-09 14:15:34.702576+08
76	1	1	37	user	菜单管理的图标 menu没有正确加载是怎么回事	\N	\N	pending	\N	2026-04-09 15:08:10.88666+08
77	1	1	38	user	你好	\N	\N	pending	\N	2026-04-10 08:31:17.394452+08
78	1	1	38	assistant	<think>\n用户只是说"你好"，这是一个简单的问候，不需要执行任何MOM系统操作。这应该被判断为chat类型。\n</think>\n\n{\n  "operation_type": "chat",\n  "parameters": {},\n  "natural_desc": "用户问候，闲聊模式",\n  "confidence": 1.0\n}	{"operation_type":"unknown","module":"","action":"","parameters":null,"natural_desc":"\\u003cthink\\u003e\\n用户只是说\\"你好\\"，这是一个简单的问候，不需要执行任何MOM系统操作。这应该被判断为chat类型。\\n\\u003c/think\\u003e\\n\\n{\\n  \\"operation_type\\": \\"chat\\",\\n  \\"parameters\\": {},\\n  \\"natural_desc\\": \\"用户问候，闲聊模式\\",\\n  \\"confidence\\": 1.0\\n}","confidence":0}	unknown	pending	\N	2026-04-10 08:31:19.770084+08
79	1	1	38	user	汇报库存余额	\N	\N	pending	\N	2026-04-10 08:32:55.248179+08
80	1	1	38	assistant	<think>\n用户想要查看库存余额，这是一个查询操作。\n\n根据系统提供的模块，我应该使用 wms (仓储管理) 模块下的 list_inventory 来查询库存列表。\n\n这是一个 query 操作，参数可能包括分页、筛选条件等。由于用户没有指定具体仓库或物料，我先查询全部库存列表。\n</think>\n\n{\n  "operation_type": "query",\n  "module": "wms",\n  "action": "list_inventory",\n  "parameters": {},\n  "natural_desc": "查询库存列表，汇报库存余额",\n  "confidence": 0.95\n}	{"operation_type":"query","module":"wms","action":"list_inventory","parameters":{},"natural_desc":"查询库存列表，汇报库存余额","confidence":0.95}	query	pending	\N	2026-04-10 08:33:01.565263+08
\.


--
-- Data for Name: ai_config; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ai_config (id, tenant_id, config_name, provider, endpoint, api_version, model_name, api_key, temperature, max_tokens, timeout, enable, created_at, updated_at) FROM stdin;
1	1	minimax	openai	https://api.minimaxi.com/v1		MiniMax-M2.7	sk-cp-41VkVKlJ2rXxGiDd5qBeTAXKgr1FOqodHA7ePb23UGYiUCq35t2yph-lveezUDpbeB6--RBHuxIgQiBLhojz6mfxx_p15gYwiJbZiuGqV8J30Gz9-eprHds	0.70	2048	60	t	0001-01-01 08:05:43+08:05:43	2026-04-09 08:45:36.831951+08
\.


--
-- Data for Name: alert_escalation_rule; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.alert_escalation_rule (id, created_at, rule_code, rule_name, alert_type, severity_level, escalation_levels, is_enabled, tenant_id, updated_at) FROM stdin;
\.


--
-- Data for Name: alert_notification_log; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.alert_notification_log (id, created_at, alert_id, alert_no, channel, receiver_type, receiver_id, receiver_name, receiver_value, notification_status, sent_time, read_time, error_code, error_msg, retry_count, tenant_id) FROM stdin;
\.


--
-- Data for Name: alert_record; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.alert_record (id, created_at, alert_no, rule_id, rule_code, rule_name, alert_type, severity_level, title, content, trigger_time, source_module, source_id, source_no, source_data, status, urgency_level, acknowledged_by, acknowledged_by_name, acknowledged_time, acknowledged_remark, resolved_by, resolved_by_name, resolved_time, resolution_remark, closed_time, escalation_count, last_escalation_time, tenant_id, updated_at) FROM stdin;
\.


--
-- Data for Name: alert_rule; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.alert_rule (id, created_at, rule_code, rule_name, alert_type, biz_module, condition_expression, condition_params, severity_level, notification_channels, notify_templates, escalation_rule_id, is_enabled, check_interval, max_trigger_count, trigger_count, last_trigger_time, tenant_id, created_by, updated_by, updated_at) FROM stdin;
\.


--
-- Data for Name: andon_call; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.andon_call (id, created_at, updated_at, deleted_at, tenant_id, call_no, workshop_id, workshop_name, production_line_id, production_line_name, workstation_id, workstation_name, andon_type, andon_type_name, call_level, priority, description, photos, audio_url, call_by, call_time, response_by, response_time, handle_by, handle_time, handle_result, handle_remark, related_order_id, related_ncr_id, related_repair_id, status, is_escalated, escalated_at, escalation_count, response_duration, handle_duration) FROM stdin;
\.


--
-- Data for Name: andon_calls; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.andon_calls (id, created_at, updated_at, deleted_at, tenant_id, call_no, line_id, line_name, station_id, station_name, call_type, call_level, call_desc, call_user_id, call_user_name, call_time, response_user_id, response_time, resolve_time, status, remark) FROM stdin;
\.


--
-- Data for Name: andon_escalation_log; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.andon_escalation_log (id, created_at, updated_at, deleted_at, call_id, from_level, to_level, escalation_type, trigger_user, trigger_reason, notify_result) FROM stdin;
\.


--
-- Data for Name: andon_escalation_rule; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.andon_escalation_rule (id, created_at, updated_at, deleted_at, rule_code, rule_name, andon_type, workshop_id, priority_range, is_default, level1_timeout, level1_notify_type, level1_notify_json, level2_timeout, level2_notify_type, level2_notify_json, level3_timeout, level3_notify_type, level3_notify_json, level4_timeout, level4_notify_type, level4_notify_json, escalation_mode, max_escalation_level, audio_enabled, audio_message_template, audio_repeat_times, is_enabled, sort_order, remark, tenant_id, created_by) FROM stdin;
\.


--
-- Data for Name: andon_notification_log; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.andon_notification_log (id, created_at, updated_at, deleted_at, call_id, channel, receiver_type, receiver_id, receiver_name, title, content, priority, send_time, send_result, error_msg, tenant_id) FROM stdin;
\.


--
-- Data for Name: andon_report; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.andon_report (id, created_at, updated_at, deleted_at, tenant_id, report_date, workshop_id, workshop_name, line_id, line_name, station_id, station_name, total_call_count, material_call_count, quality_call_count, equipment_call_count, other_call_count, avg_response_time, avg_resolve_time, unresolved_count, remark) FROM stdin;
\.


--
-- Data for Name: aps_capacity_analysis; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aps_capacity_analysis (id, created_at, updated_at, deleted_at, tenant_id, workshop_id, workshop_name, line_id, line_name, work_date, shift_id, shift_name, plan_capacity, actual_capacity, utilization, output_qty, target_qty, reject_qty, reject_rate, uptime, downtime, status) FROM stdin;
\.


--
-- Data for Name: aps_changeover_matrix; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aps_changeover_matrix (id, created_at, updated_at, deleted_at, tenant_id, from_product_id, from_product_code, from_product_name, to_product_id, to_product_code, to_product_name, changeover_time, setup_time, clean_time, is_optimized, priority, status, remark) FROM stdin;
\.


--
-- Data for Name: aps_delivery_rate; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aps_delivery_rate (id, created_at, updated_at, deleted_at, tenant_id, order_no, customer_id, customer_name, material_id, material_code, material_name, plan_delivery_date, actual_delivery_date, plan_qty, delivery_qty, on_time_qty, on_time_rate, fulfillment_rate, status) FROM stdin;
\.


--
-- Data for Name: aps_jit_demand; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aps_jit_demand (id, created_at, updated_at, deleted_at, tenant_id, demand_no, demand_type, material_id, material_code, material_name, customer_id, customer_name, demand_qty, demand_time, priority, frequency, lead_time, kanban_qty, status) FROM stdin;
\.


--
-- Data for Name: aps_mps; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aps_mps (id, created_at, updated_at, deleted_at, tenant_id, mps_no, plan_month, material_id, material_code, material_name, quantity, status) FROM stdin;
\.


--
-- Data for Name: aps_mrp; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aps_mrp (id, created_at, updated_at, deleted_at, tenant_id, mrp_no, mrp_type, plan_date, status, remark) FROM stdin;
\.


--
-- Data for Name: aps_mrp_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aps_mrp_item (id, created_at, updated_at, deleted_at, mrp_id, material_id, material_code, material_name, quantity, stock_qty, allocated_qty, net_qty, source_type, source_no) FROM stdin;
\.


--
-- Data for Name: aps_resource; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aps_resource (id, created_at, updated_at, deleted_at, tenant_id, resource_code, resource_name, resource_type, workshop_id, capacity, unit, efficiency, status) FROM stdin;
\.


--
-- Data for Name: aps_rolling_schedule; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aps_rolling_schedule (id, created_at, updated_at, deleted_at, tenant_id, schedule_no, schedule_type, start_date, end_date, workshop_id, workshop_name, line_id, line_name, plan_qty, completed_qty, start_time, end_time, horizon_days, status, execute_time) FROM stdin;
\.


--
-- Data for Name: aps_schedule_plan; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aps_schedule_plan (id, created_at, updated_at, deleted_at, tenant_id, plan_no, plan_type, start_date, end_date, algorithm, status, remark) FROM stdin;
\.


--
-- Data for Name: aps_schedule_result; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aps_schedule_result (id, created_at, updated_at, deleted_at, plan_id, order_id, order_no, sequence, line_id, line_name, station_id, station_name, plan_start_time, plan_end_time, actual_start_time, actual_end_time, status) FROM stdin;
\.


--
-- Data for Name: aps_work_center; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aps_work_center (id, created_at, updated_at, deleted_at, tenant_id, work_center_code, work_center_name, workshop_id, capacity, status, work_center_type, capacity_unit, standard_capacity, max_capacity, efficiency_factor, utilization_target, setup_time, description) FROM stdin;
\.


--
-- Data for Name: aps_working_calendar; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.aps_working_calendar (id, created_at, updated_at, deleted_at, tenant_id, workshop_id, calendar_name, work_days, shifts, holiday_dates, special_work_dates, effective_from, effective_to, status) FROM stdin;
\.


--
-- Data for Name: bom_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bom_items (id, created_at, updated_at, deleted_at, bom_id, material_id, material_code, material_name, quantity, unit, scrap_rate, level, parent_id, sort) FROM stdin;
\.


--
-- Data for Name: boms; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.boms (id, created_at, updated_at, deleted_at, tenant_id, bom_code, product_id, product_code, product_name, version, bom_type, status, eff_date, exp_date) FROM stdin;
\.


--
-- Data for Name: bpm_approval_record; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bpm_approval_record (id, created_at, task_id, instance_id, approver_id, approver_name, approver_dept, action, comment, action_time, assignee_id, assignee_name, is_rollback, rollback_target, tenant_id) FROM stdin;
\.


--
-- Data for Name: bpm_biz_mapping; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bpm_biz_mapping (id, created_at, biz_type, model_id, trigger_event, is_active, tenant_id) FROM stdin;
\.


--
-- Data for Name: bpm_delegate_record; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bpm_delegate_record (id, created_at, delegate_id, delegate_name, delegatee_id, delegatee_name, start_date, end_date, biz_types, is_active, tenant_id) FROM stdin;
\.


--
-- Data for Name: bpm_form_definition; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bpm_form_definition (id, created_at, form_code, form_name, form_type, version, category, is_published, published_at, tenant_id, created_by, updated_at) FROM stdin;
\.


--
-- Data for Name: bpm_form_field; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bpm_form_field (id, form_id, field_code, field_name, field_type, field_config, default_value, validation_rules, is_required, is_readonly, is_hidden, sort_order, tenant_id, created_at) FROM stdin;
\.


--
-- Data for Name: bpm_form_version; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bpm_form_version (id, created_at, form_id, version, change_description, fields_snapshot, published_by, published_at, is_current, tenant_id) FROM stdin;
\.


--
-- Data for Name: bpm_model_version; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bpm_model_version (id, created_at, model_id, version, change_description, published_by, published_at, is_current, tenant_id) FROM stdin;
\.


--
-- Data for Name: bpm_node_definition; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bpm_node_definition (id, model_id, node_id, node_name, node_type, position_x, position_y, width, height, node_config, sort_order, tenant_id, created_at) FROM stdin;
\.


--
-- Data for Name: bpm_node_timeout; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bpm_node_timeout (id, model_id, node_id, timeout_minutes, timeout_action, timeout_notice_roles, escalation_node_id, tenant_id, created_at) FROM stdin;
\.


--
-- Data for Name: bpm_order_day_workorder_map; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bpm_order_day_workorder_map (id, created_at, updated_at, deleted_at, day_plan_id, day_plan_item_id, work_order_id, work_order_no, tenant_id) FROM stdin;
\.


--
-- Data for Name: bpm_process_instance; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bpm_process_instance (id, created_at, instance_no, model_id, model_code, model_name, version, biz_type, biz_id, biz_no, title, initiator_id, initiator_name, current_node_id, current_node_name, status, priority, due_date, completed_at, cancel_reason, form_data, business_data, tenant_id, updated_at) FROM stdin;
\.


--
-- Data for Name: bpm_process_model; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bpm_process_model (id, created_at, model_code, model_name, model_type, version, category, description, form_type, form_definition_id, form_url, is_published, is_active, published_at, published_by, config, tenant_id, created_by, updated_at) FROM stdin;
\.


--
-- Data for Name: bpm_reminder_record; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bpm_reminder_record (id, created_at, instance_id, task_id, remind_type, remind_from, remind_to, remind_time, remind_count, tenant_id) FROM stdin;
\.


--
-- Data for Name: bpm_sequence_flow; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bpm_sequence_flow (id, model_id, flow_id, source_node_id, target_node_id, flow_name, condition_type, condition_expression, is_default, flow_config, sort_order, tenant_id, created_at) FROM stdin;
\.


--
-- Data for Name: bpm_task_assignment; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bpm_task_assignment (id, model_id, node_id, rule_type, rule_value, priority, allow_assign, allow_transfer, tenant_id, created_at) FROM stdin;
\.


--
-- Data for Name: bpm_task_instance; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.bpm_task_instance (id, task_no, instance_id, node_id, node_name, node_type, task_type, assign_type, assign_value, assignee_id, assignee_name, assignee_list, action_result, action_comment, action_time, sign_type, required_approvers, current_approvers, status, is_current, created_at, completed_at) FROM stdin;
\.


--
-- Data for Name: container_master; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.container_master (id, tenant_id, container_code, container_name, container_type, standard_qty, applicable_products, status, location_type, current_location, customer_id, barcode, total_trips, last_clean_date, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: container_movement; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.container_movement (id, tenant_id, container_id, container_code, movement_type, from_location, to_location, qty, related_order_no, operator_id, operator_name, movement_time, remark, created_at) FROM stdin;
\.


--
-- Data for Name: data_collections; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.data_collections (id, created_at, updated_at, deleted_at, tenant_id, equipment_id, equipment_code, station_id, data_type, data_key, data_value, unit, collect_time) FROM stdin;
\.


--
-- Data for Name: dc_collect_record; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.dc_collect_record (id, tenant_id, point_id, device_id, value_raw, value_numeric, value_text, quality, collect_time, created_at) FROM stdin;
\.


--
-- Data for Name: dc_data_point; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.dc_data_point (id, tenant_id, point_code, point_name, device_id, data_type, protocol, address, unit, scan_rate, deadband, store_policy, alarm_enabled, alarm_high, alarm_low, map_to_field, status, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: dc_scan_log; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.dc_scan_log (id, tenant_id, workshop_id, scan_type, scan_code, scan_device, workstation_id, scan_user_id, scan_time, parse_result, business_type, related_id, status, fail_reason, created_at) FROM stdin;
\.


--
-- Data for Name: defect_codes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.defect_codes (id, created_at, updated_at, deleted_at, tenant_id, defect_code, defect_name, defect_type, severity, status) FROM stdin;
\.


--
-- Data for Name: defect_records; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.defect_records (id, created_at, updated_at, deleted_at, tenant_id, record_no, order_id, order_no, process_id, process_name, defect_code_id, defect_code, defect_name, quantity, handle_method, handle_user_id, handle_date, status) FROM stdin;
\.


--
-- Data for Name: delivery_order_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.delivery_order_items (id, created_at, updated_at, deleted_at, delivery_id, material_id, material_code, material_name, quantity, shipped_qty, unit, batch_no) FROM stdin;
\.


--
-- Data for Name: delivery_orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.delivery_orders (id, created_at, updated_at, deleted_at, tenant_id, delivery_no, customer_id, customer_name, warehouse_id, delivery_date, delivery_user_id, status, remark) FROM stdin;
\.


--
-- Data for Name: delivery_report; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.delivery_report (id, created_at, updated_at, deleted_at, tenant_id, report_month, customer_id, customer_name, order_count, total_order_qty, delivered_qty, on_time_deliver_qty, delivery_rate, on_time_rate, late_deliver_qty, remark) FROM stdin;
\.


--
-- Data for Name: dispatches; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.dispatches (id, created_at, updated_at, deleted_at, tenant_id, order_id, order_no, process_id, process_name, station_id, station_name, assign_user_id, assign_user_name, quantity, status) FROM stdin;
\.


--
-- Data for Name: eam_equipment_downtime; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.eam_equipment_downtime (id, tenant_id, equipment_id, equipment_code, equipment_name, downtime_type, downtime_reason, start_time, end_time, duration, lost_production, work_order_id, work_order_code, shift_id, operator_id, operator_name, maintainer_id, maintainer_name, status, remark, created_by, updated_by, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: eam_repair_flow; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.eam_repair_flow (id, created_at, updated_at, deleted_at, tenant_id, flow_code, flow_name, flow_steps, status, created_by, updated_by) FROM stdin;
\.


--
-- Data for Name: eam_repair_job; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.eam_repair_job (id, created_at, updated_at, deleted_at, tenant_id, job_code, equipment_id, equipment_code, fault_type, fault_reason, fault_desc, level, status, reporter_id, reporter_name, assignee_id, assignee_name, plan_start_time, plan_end_time, actual_start_time, actual_end_time, result, evaluation, created_by, updated_by) FROM stdin;
\.


--
-- Data for Name: eam_repair_std; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.eam_repair_std (id, created_at, updated_at, deleted_at, tenant_id, std_code, std_name, fault_type, repair_steps, tools_required, materials_required, standard_hours, status, created_by, updated_by) FROM stdin;
\.


--
-- Data for Name: ene_energy_record; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ene_energy_record (id, created_at, updated_at, deleted_at, tenant_id, energy_type, workshop_id, workshop_name, equipment_id, equipment_name, meter_no, quantity, unit, unit_price, amount, record_date, remark) FROM stdin;
\.


--
-- Data for Name: energy_records; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.energy_records (id, created_at, updated_at, deleted_at, tenant_id, energy_type, workshop_id, workshop_name, equipment_id, equipment_name, meter_no, quantity, unit, unit_price, amount, record_date, remark) FROM stdin;
\.


--
-- Data for Name: equ_equipment; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.equ_equipment (id, created_at, updated_at, deleted_at, tenant_id, equipment_code, equipment_name, equipment_type, brand, model, serial_number, workshop_id, workshop_name, line_id, line_name, station_id, station_name, supplier, purchase_date, purchase_price, warranty_end_date, status) FROM stdin;
1	2026-03-31 10:48:01.60823	2026-03-31 10:48:01.60823	\N	1	EQ001	装配机器人A1	production	ABB	IRB120	\N	1	\N	1	\N	1	\N	\N	\N	\N	\N	1
2	2026-03-31 10:48:01.60823	2026-03-31 10:48:01.60823	\N	1	EQ002	检测设备C1	inspection	KEYENCE	CV-X300	\N	3	\N	\N	\N	3	\N	\N	\N	\N	\N	1
3	2026-03-31 10:48:01.60823	2026-03-31 10:48:01.60823	\N	1	EQ003	加工中心M1	production	DMG	CMX50U	\N	2	\N	3	\N	\N	\N	\N	\N	\N	\N	1
5	2026-04-04 01:34:18.206068	2026-04-04 01:34:18.206068	\N	1	EQ004	装配机械手1号	生产设备	ABB	IRB 6700	\N	2	二车间	6	二车间生产线1	6	二车间产线1工位1	精工轴承制造有限公司	2024-02-10 00:00:00	420000.00	\N	1
6	2026-04-04 01:34:18.206068	2026-04-04 01:34:18.206068	\N	1	EQ005	三坐标测量仪1号	检测设备	蔡司	CONTURA	\N	2	二车间	7	二车间生产线2	8	二车间产线2工位1	精工轴承制造有限公司	2023-11-05 00:00:00	680000.00	\N	1
\.


--
-- Data for Name: equ_equipment_check; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.equ_equipment_check (id, created_at, updated_at, deleted_at, tenant_id, equipment_id, equipment_code, equipment_name, check_plan_id, check_user_id, check_user_name, check_date, check_result, status, remark) FROM stdin;
\.


--
-- Data for Name: equ_equipment_document; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.equ_equipment_document (id, created_at, updated_at, deleted_at, tenant_id, equipment_id, equipment_code, equipment_name, doc_type, doc_name, doc_code, file_name, file_path, file_size, file_type, file_url, version, effective_date, expiry_date, description, uploaded_by, uploaded_by_name, upload_time, status) FROM stdin;
\.


--
-- Data for Name: equ_equipment_maintenance; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.equ_equipment_maintenance (id, created_at, updated_at, deleted_at, tenant_id, equipment_id, equipment_code, equipment_name, maint_type, maint_plan_id, maint_user_id, maint_user_name, maint_date, start_time, end_time, duration, content, cost, status) FROM stdin;
\.


--
-- Data for Name: equ_equipment_part; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.equ_equipment_part (id, created_at, updated_at, deleted_at, tenant_id, equipment_id, equipment_code, equipment_name, part_code, part_name, spec, unit, qty, supplier, unit_price, total_price, replacement_freq, max_stock, min_stock, current_stock, status) FROM stdin;
\.


--
-- Data for Name: equ_equipment_repair; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.equ_equipment_repair (id, created_at, updated_at, deleted_at, tenant_id, equipment_id, equipment_code, equipment_name, fault_desc, fault_time, report_user_id, report_user_name, repair_user_id, repair_user_name, start_time, end_time, duration, repair_content, cost, status) FROM stdin;
\.


--
-- Data for Name: equ_gauge; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.equ_gauge (id, created_at, updated_at, deleted_at, tenant_id, gauge_code, gauge_name, gauge_type, spec, "precision", measure_range, cal_cycle, last_cal_date, next_cal_date, workshop_id, workshop_name, location_id, location_name, status, remark) FROM stdin;
\.


--
-- Data for Name: equ_gauge_calibration; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.equ_gauge_calibration (id, created_at, updated_at, deleted_at, tenant_id, gauge_id, gauge_code, gauge_name, cal_date, cal_type, cal_result, standard, actual_value, error_value, cal_user_id, user_name, agency, certificate, cost, remark) FROM stdin;
\.


--
-- Data for Name: equ_mold; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.equ_mold (id, created_at, updated_at, deleted_at, tenant_id, mold_code, mold_name, mold_type, product_code, product_name, cavity_count, lifecycle, used_count, workshop_id, workshop_name, location_id, location_name, status, remark) FROM stdin;
\.


--
-- Data for Name: equ_mold_maintenance; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.equ_mold_maintenance (id, created_at, updated_at, deleted_at, tenant_id, mold_id, mold_code, mold_name, maint_type, maint_date, maint_user_id, user_name, content, duration, cost, status) FROM stdin;
\.


--
-- Data for Name: equ_mold_repair; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.equ_mold_repair (id, created_at, updated_at, deleted_at, tenant_id, mold_id, mold_code, mold_name, repair_type, repair_date, repair_user_id, user_name, reason, content, parts_used, duration, cost, status) FROM stdin;
\.


--
-- Data for Name: equ_oee; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.equ_oee (id, created_at, updated_at, deleted_at, tenant_id, equipment_id, equipment_code, equipment_name, workshop_id, record_date, plan_time, run_time, down_time, idle_time, plan_stop_time, output_qty, qualified_qty, availability, performance, quality, oee) FROM stdin;
\.


--
-- Data for Name: equ_oee_record; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.equ_oee_record (id, created_at, updated_at, deleted_at, tenant_id, oee_id, event_type, start_time, end_time, duration, reason) FROM stdin;
\.


--
-- Data for Name: equ_spare_part; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.equ_spare_part (id, created_at, updated_at, deleted_at, tenant_id, spare_part_code, spare_part_name, spec, unit, quantity, min_quantity, price, supplier, status) FROM stdin;
\.


--
-- Data for Name: equ_teep_data; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.equ_teep_data (id, created_at, updated_at, deleted_at, tenant_id, equipment_id, equipment_code, equipment_name, report_date, plan_time, down_time, actual_time, availability, ideal_cycle_time, actual_output, ideal_output, performance, pass_output, fail_output, quality, teep, oee, remark) FROM stdin;
\.


--
-- Data for Name: fin_payment_request; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.fin_payment_request (id, created_at, updated_at, deleted_at, request_no, request_type, supplier_customer_id, supplier_customer_name, request_amount, amount_in_words, purpose, bank_name, bank_account, settlement_ids, attachment_urls, status, approval_status, approved_by, approved_time, approver_comment, paid_by, paid_time, payment_status, tenant_id, created_by) FROM stdin;
\.


--
-- Data for Name: fin_purchase_advance; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.fin_purchase_advance (id, created_at, updated_at, deleted_at, advance_no, supplier_id, supplier_name, advance_amount, paid_date, payment_method, payment_account, bank_flow_no, status, used_amount, settlement_ids, remark, tenant_id, created_by) FROM stdin;
\.


--
-- Data for Name: fin_purchase_settlement; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.fin_purchase_settlement (id, created_at, updated_at, deleted_at, settlement_no, settlement_type, related_type, related_id, related_no, supplier_id, supplier_code, supplier_name, invoice_no, invoice_date, goods_amount, tax_amount, total_amount, paid_amount, discount_amount, currency, exchange_rate, payment_terms, payment_due_date, payment_method, status, approved_by, approved_time, settlement_date, remark, tenant_id, created_by) FROM stdin;
\.


--
-- Data for Name: fin_purchase_settlement_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.fin_purchase_settlement_item (id, created_at, updated_at, deleted_at, settlement_id, line_no, material_id, material_code, material_name, specification, unit, invoice_qty, received_qty, settled_qty, this_settle_qty, unit_price, tax_rate, goods_amount, tax_amount, line_amount, batch_no, warehouse_id, warehouse_name, remark, tenant_id) FROM stdin;
\.


--
-- Data for Name: fin_sales_receipt; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.fin_sales_receipt (id, created_at, updated_at, deleted_at, receipt_no, customer_id, customer_name, receipt_amount, receipt_date, receipt_method, receipt_account, bank_flow_no, settlement_ids, used_amount, status, remark, tenant_id, created_by) FROM stdin;
\.


--
-- Data for Name: fin_sales_settlement; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.fin_sales_settlement (id, created_at, updated_at, deleted_at, settlement_no, settlement_type, related_type, related_id, related_no, customer_id, customer_code, customer_name, invoice_no, invoice_date, goods_amount, tax_amount, total_amount, received_amount, discount_amount, currency, payment_terms, payment_due_date, payment_method, status, approved_by, approved_time, settlement_date, remark, tenant_id, created_by) FROM stdin;
\.


--
-- Data for Name: fin_sales_settlement_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.fin_sales_settlement_item (id, created_at, updated_at, deleted_at, settlement_id, line_no, material_id, material_code, material_name, specification, unit, invoice_qty, shipped_qty, settled_qty, this_settle_qty, unit_price, tax_rate, goods_amount, tax_amount, line_amount, batch_no, remark, tenant_id) FROM stdin;
\.


--
-- Data for Name: fin_statement_detail; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.fin_statement_detail (id, created_at, updated_at, deleted_at, statement_id, detail_type, related_no, biz_date, amount, remark, tenant_id) FROM stdin;
\.


--
-- Data for Name: fin_supplier_statement; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.fin_supplier_statement (id, created_at, updated_at, deleted_at, statement_no, supplier_id, supplier_name, statement_period, start_date, end_date, beginning_amount, purchase_amount, other_amount, payment_amount, discount_amount, adjustment_amount, ending_amount, status, confirmed_by, confirmed_time, dispute_reason, tenant_id, created_by) FROM stdin;
\.


--
-- Data for Name: fqcs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.fqcs (id, created_at, updated_at, deleted_at, tenant_id, fqc_no, order_id, order_no, quantity, sample_size, qualified_qty, rejected_qty, check_user_id, check_user_name, check_date, result, remark) FROM stdin;
\.


--
-- Data for Name: integration_erp_mappings; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.integration_erp_mappings (id, erp_table_name, erp_field_name, mes_table_name, mes_field_name, transform_rule, tenant_id, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: integration_erp_sync_logs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.integration_erp_sync_logs (id, sync_type, direction, erp_bill_no, mes_bill_no, request_body, response_body, status, error_msg, retry_count, tenant_id, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: inventories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.inventories (id, created_at, updated_at, deleted_at, tenant_id, material_id, material_code, material_name, warehouse_id, location_id, quantity, available_qty, allocated_qty, locked_qty, batch_no) FROM stdin;
\.


--
-- Data for Name: inventory_records; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.inventory_records (id, created_at, updated_at, deleted_at, tenant_id, record_no, record_type, material_id, material_code, material_name, warehouse_id, location_id, quantity, batch_no, source_no, operator_id, operator_name, operate_time, remark) FROM stdin;
\.


--
-- Data for Name: ip_qcs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ip_qcs (id, created_at, updated_at, deleted_at, tenant_id, ip_qc_no, order_id, order_no, process_id, process_name, quantity, sample_size, check_user_id, check_user_name, check_date, result, remark) FROM stdin;
\.


--
-- Data for Name: iqc_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.iqc_items (id, created_at, updated_at, deleted_at, iqc_id, check_item_id, check_item, check_standard, check_method, result, remark) FROM stdin;
\.


--
-- Data for Name: iqcs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.iqcs (id, created_at, updated_at, deleted_at, tenant_id, iqc_no, supplier_id, supplier_name, material_id, material_code, material_name, quantity, unit, check_user_id, check_user_name, check_date, result, remark) FROM stdin;
\.


--
-- Data for Name: lab_calibration; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.lab_calibration (id, tenant_id, instrument_id, calibration_date, calibration_result, calibrated_by, certificate_no, next_calibration_date, calibration_items, attachment_url, remark, created_by, updated_by, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: lab_instrument; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.lab_instrument (id, tenant_id, instrument_code, instrument_name, instrument_type, manufacturer, model, serial_number, calibration_cycle, last_calibration_date, next_calibration_date, calibration_status, lab_location, status, remark, created_by, updated_by, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: lab_report; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.lab_report (id, tenant_id, report_no, sample_id, conclusion, remarks, attachments, inspector_id, inspector_name, approved_by, report_date, created_by, updated_by, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: lab_sample; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.lab_sample (id, tenant_id, sample_code, sample_name, inspection_type, source_type, source_id, source_no, quantity, sample_qty, received_by, received_at, due_date, status, remark, created_by, updated_by, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: lab_test_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.lab_test_item (id, tenant_id, sample_id, item_name, test_method, standard, upper_limit, lower_limit, unit, result, result_value, tester_id, tester_name, tested_at, remark, created_at) FROM stdin;
\.


--
-- Data for Name: locations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.locations (id, created_at, updated_at, deleted_at, tenant_id, location_code, location_name, warehouse_id, zone_code, "row", col, layer, location_type, capacity, status) FROM stdin;
\.


--
-- Data for Name: login_logs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.login_logs (id, created_at, updated_at, deleted_at, tenant_id, username, ip, login_location, browser, os, status, msg, login_time) FROM stdin;
\.


--
-- Data for Name: material_categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.material_categories (id, created_at, updated_at, deleted_at, tenant_id, parent_id, category_name, category_code, sort, status) FROM stdin;
\.


--
-- Data for Name: materials; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.materials (id, created_at, updated_at, deleted_at, tenant_id, material_code, material_name, material_type, spec, unit, unit_name, weight, length, width, height, category_id, status) FROM stdin;
\.


--
-- Data for Name: mdm_bom; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mdm_bom (id, created_at, updated_at, deleted_at, tenant_id, bom_code, bom_name, material_id, material_code, material_name, version, status, eff_date, exp_date, remark, product_id, product_code, product_name, bom_type) FROM stdin;
1	2026-04-04 01:36:12.081969+08	2026-04-04 01:36:12.081969+08	\N	1	BOM-001	电机产品BOM	27	MAT-027	控制柜A	V1.0	ACTIVE	2026-04-01	\N	\N	27	MAT-027	控制柜A	standard
2	2026-04-04 01:36:12.081969+08	2026-04-04 01:36:12.081969+08	\N	1	BOM-002	电机产品BOM	28	MAT-028	控制柜B	V1.0	ACTIVE	2026-04-01	\N	\N	28	MAT-028	控制柜B	standard
\.


--
-- Data for Name: mdm_bom_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mdm_bom_item (id, created_at, updated_at, deleted_at, tenant_id, bom_id, line_no, material_id, material_code, material_name, quantity, unit, scrap_rate, substitute_group, is_alternative) FROM stdin;
1	2026-04-04 01:36:31.811581+08	2026-04-04 01:36:31.811581+08	\N	1	1	1	1	MAT-001	钢板A3	2.5000	张	0.0200	\N	0
2	2026-04-04 01:36:31.811581+08	2026-04-04 01:36:31.811581+08	\N	1	1	2	9	MAT-009	轴承6205	4.0000	个	0.0100	\N	0
3	2026-04-04 01:36:31.811581+08	2026-04-04 01:36:31.811581+08	\N	1	1	3	11	MAT-011	电机0.75KW	1.0000	台	0.0050	\N	0
4	2026-04-04 01:36:31.811581+08	2026-04-04 01:36:31.811581+08	\N	1	1	4	13	MAT-013	接触器LC1D	2.0000	个	0.0100	\N	0
5	2026-04-04 01:36:31.811581+08	2026-04-04 01:36:31.811581+08	\N	1	1	5	14	MAT-014	路基NSX	2.0000	个	0.0100	\N	0
6	2026-04-04 01:36:31.811581+08	2026-04-04 01:36:31.811581+08	\N	1	1	6	15	MAT-015	开关电源24V	1.0000	个	0.0100	\N	0
7	2026-04-04 01:36:31.811581+08	2026-04-04 01:36:31.811581+08	\N	1	2	1	2	MAT-002	钢板A4	3.0000	张	0.0200	\N	0
8	2026-04-04 01:36:31.811581+08	2026-04-04 01:36:31.811581+08	\N	1	2	2	10	MAT-010	轴承6208	4.0000	个	0.0100	\N	0
9	2026-04-04 01:36:31.811581+08	2026-04-04 01:36:31.811581+08	\N	1	2	3	12	MAT-012	电机1.5KW	1.0000	台	0.0050	\N	0
10	2026-04-04 01:36:31.811581+08	2026-04-04 01:36:31.811581+08	\N	1	2	4	13	MAT-013	接触器LC1D	3.0000	个	0.0100	\N	0
11	2026-04-04 01:36:31.811581+08	2026-04-04 01:36:31.811581+08	\N	1	2	5	16	MAT-016	PLC模块	1.0000	个	0.0100	\N	0
\.


--
-- Data for Name: mdm_customer; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mdm_customer (id, created_at, updated_at, deleted_at, tenant_id, code, name, type, contact, phone, email, address, status) FROM stdin;
1	2026-04-09 09:38:03.266514	2026-04-09 09:38:03.266514	\N	1	CUS001	杭州汽车配件有限公司	end_customer	赵经理	13900010001	zhao@hzauto.com	杭州市萧山区	1
2	2026-04-09 09:38:03.266514	2026-04-09 09:38:03.266514	\N	1	CUS002	上海机电设备有限公司	dealer	钱经理	13900010002	qian@shjd.com	上海市闵行区	1
3	2026-04-09 09:38:03.266514	2026-04-09 09:38:03.266514	\N	1	CUS003	宁波电子科技有限公司	end_customer	孙经理	13900010003	sun@nbec.com	宁波市鄞州区	1
4	2026-04-09 09:38:03.266514	2026-04-09 09:38:03.266514	\N	1	CUS004	温州电器股份有限公司	agent	李经理	13900010004	li@wzelectric.com	温州市乐清市	1
\.


--
-- Data for Name: mdm_material; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mdm_material (id, created_at, updated_at, deleted_at, tenant_id, material_code, material_name, material_type, spec, unit, unit_name, weight, length, width, height, category_id, status) FROM stdin;
1	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-001	钢板A3	raw	1200*2400mm	PCS	张	25.5000	1200.0000	2400.0000	2.0000	\N	1
2	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-002	钢板A4	raw	1200*1800mm	PCS	张	19.5000	1200.0000	1800.0000	2.0000	\N	1
3	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-003	铝合金板	raw	1000*2000mm	PCS	张	15.0000	1000.0000	2000.0000	2.0000	\N	1
4	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-004	不锈钢板	raw	1500*3000mm	PCS	张	45.0000	1500.0000	3000.0000	3.0000	\N	1
5	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-005	塑料粒子PP	raw	通用级	KG	千克	1000.0000	\N	\N	\N	\N	1
6	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-006	塑料粒子ABS	raw	通用级	KG	千克	1000.0000	\N	\N	\N	\N	1
7	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-007	铜线	raw	2.0mm	KG	千克	50.0000	\N	\N	\N	\N	1
8	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-008	铝线	raw	1.5mm	KG	千克	30.0000	\N	\N	\N	\N	1
9	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-009	轴承6205	parts	内径25mm	PCS	个	0.3000	\N	\N	\N	\N	1
10	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-010	轴承6208	parts	内径40mm	PCS	个	0.5000	\N	\N	\N	\N	1
11	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-011	电机0.75KW	parts	0.75KW/380V	PCS	台	8.5000	\N	\N	\N	\N	1
12	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-012	电机1.5KW	parts	1.5KW/380V	PCS	台	12.0000	\N	\N	\N	\N	1
13	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-013	接触器LC1D	parts	AC-3 40A	PCS	个	0.4000	\N	\N	\N	\N	1
14	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-014	断路器NSX	parts	63A	PCS	个	0.5000	\N	\N	\N	\N	1
15	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-015	开关电源24V	parts	DC24V 10A	PCS	个	0.8000	\N	\N	\N	\N	1
16	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-016	PLC控制器	parts	S7-200Smart	PCS	个	0.5000	\N	\N	\N	\N	1
17	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-017	触摸屏7寸	parts	7寸彩色	PCS	个	1.2000	\N	\N	\N	\N	1
18	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-018	传感器光电	parts	NPN	PCS	个	0.1000	\N	\N	\N	\N	1
19	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-019	传感器接近	parts	电感式	PCS	个	0.1500	\N	\N	\N	\N	1
20	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-020	气缸标准	parts	ISO15552	PCS	个	2.5000	\N	\N	\N	\N	1
21	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-021	电磁阀	parts	5V二位五通	PCS	个	0.3000	\N	\N	\N	\N	1
22	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-022	液压油	aux	ISO VG46	L	升	0.9000	\N	\N	\N	\N	1
23	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-023	润滑油	aux	ISO VG32	L	升	0.8500	\N	\N	\N	\N	1
24	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-024	清洁剂	aux	工业级	L	升	1.0000	\N	\N	\N	\N	1
25	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-025	包装箱	package	600*400*300	PCS	个	0.5000	600.0000	400.0000	300.0000	\N	1
26	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-026	说明书	package	A5	PCS	本	0.1000	\N	\N	\N	\N	1
27	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-027	电机组件A	semi		SET	套	5.0000	\N	\N	\N	\N	1
28	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-028	电机组件B	semi		SET	套	6.0000	\N	\N	\N	\N	1
29	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-029	控制柜	semi		SET	套	25.0000	\N	\N	\N	\N	1
30	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MAT-030	配电箱	semi		SET	套	15.0000	\N	\N	\N	\N	1
35	2026-04-14 14:05:27.294281+08	2026-04-14 14:05:27.294281+08	\N	1	MAT001	��������		\N	PCS	\N	\N	\N	\N	\N	\N	1
\.


--
-- Data for Name: mdm_material_category; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mdm_material_category (id, created_at, updated_at, deleted_at, tenant_id, parent_id, category_name, category_code, sort, status) FROM stdin;
\.


--
-- Data for Name: mdm_operation; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mdm_operation (id, created_at, updated_at, deleted_at, tenant_id, operation_code, operation_name, workcenter_id, workcenter_name, standard_worktime, quality_std, is_key_process, is_qc_point, sequence, remark) FROM stdin;
1	2026-04-04 01:35:49.090387+08	2026-04-04 01:35:49.090387+08	\N	1	OP001	原材料检验	1	质检工段	30	符合原材料检验标准	1	1	1	\N
2	2026-04-04 01:35:49.090387+08	2026-04-04 01:35:49.090387+08	\N	1	OP002	钣金加工	1	加工工段	120	尺寸公差±0.5mm	1	0	2	\N
3	2026-04-04 01:35:49.090387+08	2026-04-04 01:35:49.090387+08	\N	1	OP003	电机装配	1	装配工段	90	装配牢固,转动灵活	1	1	3	\N
4	2026-04-04 01:35:49.090387+08	2026-04-04 01:35:49.090387+08	\N	1	OP004	电气接线	1	装配工段	60	接线正确,绝缘合格	1	0	4	\N
5	2026-04-04 01:35:49.090387+08	2026-04-04 01:35:49.090387+08	\N	1	OP005	成品检验	1	质检工段	45	符合出厂检验标准	1	1	5	\N
6	2026-04-04 01:35:49.090387+08	2026-04-04 01:35:49.090387+08	\N	1	OP006	包装入库	1	包装工段	30	包装完整,标识清晰	0	0	6	\N
\.


--
-- Data for Name: mdm_production_line; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mdm_production_line (id, created_at, updated_at, deleted_at, tenant_id, line_code, line_name, workshop_id, line_type, status) FROM stdin;
1	2026-03-31 10:48:01.603525+08	2026-03-31 10:48:01.603525+08	\N	1	LINE001	组装线A	1	automation	1
2	2026-03-31 10:48:01.603525+08	2026-03-31 10:48:01.603525+08	\N	1	LINE002	组装线B	1	automation	1
3	2026-03-31 10:48:01.603525+08	2026-03-31 10:48:01.603525+08	\N	1	LINE003	加工线C	2	semiautomation	1
4	2026-04-04 01:32:53.304798+08	2026-04-04 01:32:53.304798+08	\N	1	PL001	一车间生产线1	1	主生产线	1
5	2026-04-04 01:32:53.304798+08	2026-04-04 01:32:53.304798+08	\N	1	PL002	一车间生产线2	1	辅生产线	1
6	2026-04-04 01:32:53.304798+08	2026-04-04 01:32:53.304798+08	\N	1	PL003	二车间生产线1	2	主生产线	1
7	2026-04-04 01:32:53.304798+08	2026-04-04 01:32:53.304798+08	\N	1	PL004	二车间生产线2	2	辅生产线	1
8	2026-04-04 01:32:53.304798+08	2026-04-04 01:32:53.304798+08	\N	1	PL005	三车间生产线1	3	主生产线	1
\.


--
-- Data for Name: mdm_shift; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mdm_shift (id, created_at, updated_at, deleted_at, tenant_id, shift_code, shift_name, start_time, end_time, break_start, break_end, status, work_hours, is_night, remark) FROM stdin;
1	2026-04-04 01:32:09.742479+08	2026-04-04 01:32:09.742479+08	\N	1	SHIFT_DAY	早班	08:00	17:00	12:00	13:00	1	8.00	0	\N
2	2026-04-04 01:32:09.742479+08	2026-04-04 01:32:09.742479+08	\N	1	SHIFT_MID	中班	16:00	01:00	20:00	21:00	1	8.00	0	\N
3	2026-04-04 01:32:09.742479+08	2026-04-04 01:32:09.742479+08	\N	1	SHIFT_NIGHT	晚班	00:00	08:00	04:00	05:00	1	8.00	1	\N
\.


--
-- Data for Name: mdm_supplier; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mdm_supplier (id, created_at, updated_at, deleted_at, tenant_id, code, name, type, contact, phone, email, address, category, level, status, remark) FROM stdin;
20	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:03.867722+08	1	SUP-020	镇江五金工具	\N	蒋经理	13800010020	jiang@zjwj.com	镇江市京口区	\N	1	1	\N
19	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:03.935816+08	1	SUP-019	扬州环保设备	\N	卫经理	13800010019	wei@yzhb.com	扬州市广陵区	\N	1	1	\N
18	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:03.999521+08	1	SUP-018	南通船舶配件	\N	冯经理	13800010018	feng@ntcp.com	南通市崇川区	\N	1	1	\N
17	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:04.066303+08	1	SUP-017	常州自动化	\N	孙经理	13800010017	sun@czauto.com	常州市新北区	\N	1	1	\N
15	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:04.199688+08	1	SUP-015	苏州电子科技	\N	谢经理	13800010015	xie@szec.com	苏州市工业园区	\N	1	1	\N
14	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:04.264514+08	1	SUP-014	南京钢材集团	\N	戚经理	13800010014	qi@njsteel.com	南京市栖霞区	\N	1	1	\N
13	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:04.329877+08	1	SUP-013	义乌商贸有限公司	\N	褚经理	13800010013	chu@ywsm.com	义乌市	\N	1	1	\N
12	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:04.39466+08	1	SUP-012	舟山港务有限公司	\N	冯经理	13800010012	feng@zsgw.com	舟山市定海区	\N	1	1	\N
11	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:04.45814+08	1	SUP-011	丽水电机有限公司	\N	朱经理	13800010011	zhu@lsmotor.com	丽水市莲都区	\N	1	1	\N
10	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:04.521273+08	1	SUP-010	台州机械有限公司	\N	马经理	13800010010	ma@tzmachine.com	台州市椒江区	\N	1	1	\N
9	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:04.583034+08	1	SUP-009	衢州矿业有限公司	\N	孙经理	13800010009	sun@qzmining.com	衢州市衢江区	\N	1	1	\N
8	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:04.64476+08	1	SUP-008	金华化工有限公司	\N	徐经理	13800010008	xu@jhchemical.com	金华市婺城区	\N	1	1	\N
7	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:04.709097+08	1	SUP-007	绍兴纺织有限公司	\N	吴经理	13800010007	wu@sxtextile.com	绍兴市柯桥区	\N	1	1	\N
6	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:04.769922+08	1	SUP-006	湖州木材有限公司	\N	周经理	13800010006	zhou@hzwood.com	湖州市南浔区	\N	1	1	\N
5	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:04.828955+08	1	SUP-005	嘉兴包装有限公司	\N	刘经理	13800010005	liu@jxpackage.com	嘉兴市秀洲区	\N	1	1	\N
4	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:04.897843+08	1	SUP-004	温州五金有限公司	\N	陈经理	13800010004	chen@wzhardware.com	温州市瓯海区	\N	1	1	\N
3	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:04.963213+08	1	SUP-003	宁波电子有限公司	\N	王经理	13800010003	wang@nbelectronic.com	宁波市高新区	\N	1	1	\N
1	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:05.093719+08	1	SUP-001	上海钢材有限公司	\N	张经理	13800010001	zhang@shsteel.com	上海市宝山区	\N	1	1	\N
16	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:04.133306+08	1	SUP-016	无锡液压设备	\N	戚经理	13800010016	xu@wxhyd.com	无锡市锡山区	\N	1	1	\N
2	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	2026-04-14 14:56:05.028375+08	1	SUP-002	杭州塑料有限公司	\N	李经理	13800010002	li@hzplastic.com	杭州市余杭区	\N	1	1	\N
61	2026-04-14 14:58:15.270965+08	2026-04-14 14:58:15.270965+08	\N	1	SUP-TEST01	华东精密机械有限公司	原材料	张经理	13812345601	\N	\N	\N	1	1	\N
62	2026-04-14 14:58:15.275037+08	2026-04-14 14:58:15.275037+08	\N	1	SUP-TEST02	深圳创新电子科技有限公司	原材料	李总监	13812345602	\N	\N	\N	1	1	\N
63	2026-04-14 14:58:15.277331+08	2026-04-14 14:58:15.277331+08	\N	1	SUP-TEST03	苏州工业园区华鑫塑胶制品厂	辅料	王厂长	13812345603	\N	\N	\N	1	1	\N
64	2026-04-14 14:58:15.279537+08	2026-04-14 14:58:15.279537+08	\N	1	SUP-TEST04	杭州中达金属表面处理有限公司	服务	陈工	13812345604	\N	\N	\N	1	1	\N
65	2026-04-14 14:58:15.282316+08	2026-04-14 14:58:15.282316+08	\N	1	SUP-TEST05	广州岭南包装材料有限公司	辅料	刘经理	13812345605	\N	\N	\N	1	1	\N
66	2026-04-14 14:58:15.284925+08	2026-04-14 14:58:15.284925+08	\N	1	SUP-TEST06	北京中关村软件股份有限公司	设备	赵总	13812345606	\N	\N	\N	1	1	\N
67	2026-04-14 14:58:15.287841+08	2026-04-14 14:58:15.287841+08	\N	1	SUP-TEST07	天津港保税区润通物流有限公司	服务	孙经理	13812345607	\N	\N	\N	1	1	\N
68	2026-04-14 14:58:15.289981+08	2026-04-14 14:58:15.289981+08	\N	1	SUP-TEST08	成都西南航空材料有限公司	原材料	周经理	13812345608	\N	\N	\N	1	1	\N
69	2026-04-14 14:58:15.292041+08	2026-04-14 14:58:15.292041+08	\N	1	SUP-TEST09	武汉光谷激光设备有限公司	设备	吴工	13812345609	\N	\N	\N	1	1	\N
70	2026-04-14 14:58:15.293597+08	2026-04-14 14:58:15.293597+08	\N	1	SUP-TEST10	南京金陵化工有限公司	原材料	郑经理	13812345610	\N	\N	\N	1	1	\N
71	2026-04-14 14:58:15.294713+08	2026-04-14 14:58:15.294713+08	\N	1	SUP-TEST11	西安古都标准件有限公司	辅料	冯经理	13812345611	\N	\N	\N	1	1	\N
72	2026-04-14 14:58:15.296486+08	2026-04-14 14:58:15.296486+08	\N	1	SUP-TEST12	青岛海尔智造装备有限公司	设备	卫总	13812345612	\N	\N	\N	1	1	\N
73	2026-04-14 14:58:15.298642+08	2026-04-14 14:58:15.298642+08	\N	1	SUP-TEST13	厦门特区进出口贸易有限公司	原材料	蒋经理	13812345613	\N	\N	\N	1	1	\N
74	2026-04-14 14:58:15.300735+08	2026-04-14 14:58:15.300735+08	\N	1	SUP-TEST14	长沙中沙电缆桥架有限公司	辅料	沈经理	13812345614	\N	\N	\N	1	1	\N
75	2026-04-14 14:58:15.301926+08	2026-04-14 14:58:15.301926+08	\N	1	SUP-TEST15	东莞松山湖智能装备有限公司	设备	楚经理	13812345615	\N	\N	\N	1	1	\N
\.


--
-- Data for Name: mdm_supplier_material; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mdm_supplier_material (id, created_at, updated_at, supplier_id, material_id, material_code, material_name, supplier_part_no, price, currency, min_order_qty, lead_time, is_preferred, status, tenant_id) FROM stdin;
\.


--
-- Data for Name: mdm_workshop; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mdm_workshop (id, created_at, updated_at, deleted_at, tenant_id, workshop_code, workshop_name, workshop_type, manager, phone, address, status) FROM stdin;
1	2026-03-31 10:48:01.602127+08	2026-03-31 10:48:01.602127+08	\N	1	WS001	一车间	assembly	张三	13800138001	\N	1
2	2026-03-31 10:48:01.602127+08	2026-03-31 10:48:01.602127+08	\N	1	WS002	二车间	machining	李四	13800138002	\N	1
3	2026-03-31 10:48:01.602127+08	2026-03-31 10:48:01.602127+08	\N	1	WS003	三车间	inspection	王五	13800138003	\N	1
\.


--
-- Data for Name: mdm_workshop_config; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mdm_workshop_config (id, created_at, updated_at, deleted_at, tenant_id, workshop_id, erp_plant_code, max_devices, max_workers, max_capacity_per_day, time_zone, is_default) FROM stdin;
\.


--
-- Data for Name: mdm_workstation; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mdm_workstation (id, created_at, updated_at, deleted_at, tenant_id, station_code, station_name, line_id, station_type, status) FROM stdin;
1	2026-03-31 10:48:01.605288+08	2026-03-31 10:48:01.605288+08	\N	1	ST001	组装工位1	1	assembly	1
2	2026-03-31 10:48:01.605288+08	2026-03-31 10:48:01.605288+08	\N	1	ST002	组装工位2	1	assembly	1
3	2026-03-31 10:48:01.605288+08	2026-03-31 10:48:01.605288+08	\N	1	ST003	检测工位	1	inspection	1
4	2026-04-04 01:33:50.548366+08	2026-04-04 01:33:50.548366+08	\N	1	WS001-01	一车间产线1工位1	4	加工	1
5	2026-04-04 01:33:50.548366+08	2026-04-04 01:33:50.548366+08	\N	1	WS001-02	一车间产线1工位2	4	加工	1
6	2026-04-04 01:33:50.548366+08	2026-04-04 01:33:50.548366+08	\N	1	WS001-03	一车间产线1工位3	4	装配	1
7	2026-04-04 01:33:50.548366+08	2026-04-04 01:33:50.548366+08	\N	1	WS002-01	一车间产线2工位1	5	加工	1
8	2026-04-04 01:33:50.548366+08	2026-04-04 01:33:50.548366+08	\N	1	WS002-02	一车间产线2工位2	5	检测	1
9	2026-04-04 01:33:50.548366+08	2026-04-04 01:33:50.548366+08	\N	1	WS003-01	二车间产线1工位1	6	加工	1
10	2026-04-04 01:33:50.548366+08	2026-04-04 01:33:50.548366+08	\N	1	WS003-02	二车间产线1工位2	6	装配	1
11	2026-04-04 01:33:50.548366+08	2026-04-04 01:33:50.548366+08	\N	1	WS004-01	二车间产线2工位1	7	加工	1
12	2026-04-04 01:33:50.548366+08	2026-04-04 01:33:50.548366+08	\N	1	WS005-01	三车间产线1工位1	8	加工	1
13	2026-04-04 01:33:50.548366+08	2026-04-04 01:33:50.548366+08	\N	1	WS005-02	三车间产线1工位2	8	装配	1
\.


--
-- Data for Name: mes_code_rule; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_code_rule (id, created_at, updated_at, deleted_at, tenant_id, rule_code, rule_name, entity_type, prefix, date_format, seq_length, seq_start, seq_current, mid_fix, suffix, reset_type, last_gen_date, example, status, remark) FROM stdin;
\.


--
-- Data for Name: mes_code_rule_record; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_code_rule_record (id, created_at, updated_at, deleted_at, tenant_id, rule_id, rule_code, entity_type, gen_date, seq_value, gen_code) FROM stdin;
\.


--
-- Data for Name: mes_container_lifecycle; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_container_lifecycle (id, created_at, updated_at, tenant_id, container_id, container_code, event_type, event_date, operator_id, operator_name, location_id, location_name, status, remark) FROM stdin;
\.


--
-- Data for Name: mes_container_maintenance; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_container_maintenance (id, created_at, updated_at, tenant_id, container_id, container_code, maintenance_type, maintenance_date, completed_date, technician_id, technician_name, fault_description, maintenance_content, spare_parts_used, cost, status, remark) FROM stdin;
\.


--
-- Data for Name: mes_electronic_sop; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_electronic_sop (id, created_at, updated_at, deleted_at, tenant_id, sop_no, sop_name, material_id, material_code, material_name, version, process_id, process_name, content_type, content_url, thumbnail_url, steps, workstation_id, workstation_name, workshop_id, workshop_name, status, eff_date, exp_date, approved_by, approved_at, remark) FROM stdin;
\.


--
-- Data for Name: mes_first_last_inspect; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_first_last_inspect (id, tenant_id, workshop_id, inspect_no, inspect_type, production_order_id, process_id, workstation_id, shift_id, product_id, serial_no, inspect_items, overall_result, inspector_id, inspector_name, inspect_time, bluetooth_device_id, remark, status, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: mes_flow_card; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_flow_card (id, created_at, updated_at, deleted_at, tenant_id, card_no, order_id, order_no, material_id, material_code, material_name, workshop_id, workshop_name, line_id, line_name, process_id, process_name, station_id, station_name, plan_qty, completed_qty, status, priority, plan_start_time, plan_end_time, actual_start_time, actual_end_time, remark) FROM stdin;
\.


--
-- Data for Name: mes_flow_card_detail; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_flow_card_detail (id, created_at, updated_at, deleted_at, card_id, step_no, process_id, process_name, station_id, station_name, work_content, std_cycle_time, seq_qty) FROM stdin;
\.


--
-- Data for Name: mes_job_report_log; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_job_report_log (id, tenant_id, work_order_id, work_order_code, process_code, process_name, report_type, quantity, report_time, reporter_id, reporter_name, remark, status, created_at) FROM stdin;
\.


--
-- Data for Name: mes_order_day; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_order_day (id, created_at, updated_at, deleted_at, tenant_id, day_plan_no, plan_date, month_plan_id, month_plan_no, workshop_id, workshop_name, production_line_id, line_name, total_product_count, total_plan_qty, total_completed_qty, kit_status, kit_check_time, kit_check_by, status, shift_type, published_at, published_by, remark, created_by) FROM stdin;
\.


--
-- Data for Name: mes_order_day_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_order_day_item (id, created_at, updated_at, deleted_at, day_plan_id, line_no, product_id, product_code, product_name, specification, unit, plan_qty, completed_qty, bom_id, bom_version, process_route_id, route_version, production_mode, kit_status, kit_check_remark, month_plan_item_id, item_status, priority, remark, tenant_id) FROM stdin;
\.


--
-- Data for Name: mes_order_month; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_order_month (id, created_at, updated_at, deleted_at, tenant_id, month_plan_no, plan_month, title, source_type, source_no, workshop_id, workshop_name, total_product_count, total_plan_qty, total_completed_qty, approval_status, submitted_by, submitted_at, approved_by, approved_at, released_by, released_at, remark, created_by) FROM stdin;
\.


--
-- Data for Name: mes_order_month_audit; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_order_month_audit (id, created_at, updated_at, deleted_at, month_plan_id, approval_status, approver_id, approver_name, approval_time, comment, tenant_id) FROM stdin;
\.


--
-- Data for Name: mes_order_month_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_order_month_item (id, created_at, updated_at, deleted_at, month_plan_id, line_no, product_id, product_code, product_name, specification, unit, plan_qty, completed_qty, released_qty, delivery_date, priority, remark, tenant_id) FROM stdin;
\.


--
-- Data for Name: mes_package; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_package (id, tenant_id, workshop_id, package_no, package_type, production_order_id, product_id, product_code, qty, serial_nos, status, seal_time, seal_by, ship_time, customer_id, container_id, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: mes_person_skill; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_person_skill (id, created_at, updated_at, deleted_at, tenant_id, person_id, person_code, person_name, workshop_id, workstation_id, skill_level, certified_date, expiry_date, status, created_by) FROM stdin;
\.


--
-- Data for Name: mes_person_skill_score; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_person_skill_score (id, created_at, updated_at, deleted_at, tenant_id, person_skill_id, skill_type, score, evaluated_by, evaluator_name, evaluated_at, remark) FROM stdin;
\.


--
-- Data for Name: mes_process; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_process (id, created_at, updated_at, deleted_at, tenant_id, process_code, process_name, material_id, material_code, material_name, version, status, eff_date, exp_date, remark, is_current) FROM stdin;
\.


--
-- Data for Name: mes_process_operation; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_process_operation (id, created_at, updated_at, deleted_at, tenant_id, process_id, operation_id, operation_code, operation_name, line_no, standard_worktime, workcenter_id, workcenter_name, required_capacity, min_workers, max_workers, is_key_process, is_qc_point, quality_std, status, remark) FROM stdin;
\.


--
-- Data for Name: mes_team; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_team (id, created_at, updated_at, deleted_at, tenant_id, team_code, team_name, workshop_id, shift_id, leader_id, phone, remark, status) FROM stdin;
\.


--
-- Data for Name: mes_team_member; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_team_member (id, created_at, updated_at, deleted_at, tenant_id, team_id, user_id, user_name, role, join_date, remark, status) FROM stdin;
\.


--
-- Data for Name: mes_team_shift; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_team_shift (id, created_at, updated_at, deleted_at, tenant_id, team_id, shift_id, shift_date, start_time, end_time, leader_id, remark) FROM stdin;
\.


--
-- Data for Name: mes_work_scheduling; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_work_scheduling (id, created_at, updated_at, deleted_at, tenant_id, plan_no_day, scheduling_code, product_code, product_name, status, quantity, finished_qty, work_mode, task_mode, plan_date, workshop_code, line_code, created_by, updated_by) FROM stdin;
\.


--
-- Data for Name: mes_work_scheduling_detail; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mes_work_scheduling_detail (id, created_at, updated_at, deleted_at, tenant_id, scheduling_id, working_node, working_name, status, equipment_id, equipment_code, workstation_id, workstation_name, worker_id, worker_name, start_time, end_time, plan_qty, finished_qty, work_minutes, created_by, updated_by) FROM stdin;
\.


--
-- Data for Name: mps; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mps (id, created_at, updated_at, deleted_at, tenant_id, mps_no, plan_month, material_id, material_code, material_name, quantity, status) FROM stdin;
\.


--
-- Data for Name: mrp_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mrp_items (id, created_at, updated_at, deleted_at, mrp_id, material_id, material_code, material_name, quantity, stock_qty, allocated_qty, net_qty, source_type, source_no) FROM stdin;
\.


--
-- Data for Name: mrps; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.mrps (id, created_at, updated_at, deleted_at, tenant_id, mrp_no, mrp_type, plan_date, status, remark) FROM stdin;
\.


--
-- Data for Name: ncrs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.ncrs (id, created_at, updated_at, deleted_at, tenant_id, ncr_no, defect_id, source_type, issue_desc, root_cause, corrective_action, preventive_action, verify_result, verify_user_id, verify_date, status) FROM stdin;
\.


--
-- Data for Name: notification_channel; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.notification_channel (id, created_at, channel_code, channel_name, channel_type, config, is_enabled, priority, tenant_id, created_by, updated_by, updated_at) FROM stdin;
\.


--
-- Data for Name: oee_report; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.oee_report (id, created_at, updated_at, deleted_at, tenant_id, report_date, workshop_id, workshop_name, line_id, line_name, availability, performance, quality, oee, planned_production_time, actual_production_time, down_time, speed_loss, defect_loss, remark) FROM stdin;
\.


--
-- Data for Name: oper_logs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.oper_logs (id, created_at, updated_at, deleted_at, tenant_id, title, business_type, method, request_method, operator_type, oper_name, dept_name, oper_url, oper_ip, oper_location, oper_param, json_result, status, error_msg, oper_time) FROM stdin;
\.


--
-- Data for Name: oqcs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.oqcs (id, created_at, updated_at, deleted_at, tenant_id, oqc_no, shipping_no, customer_id, customer_name, quantity, check_user_id, check_user_name, check_date, result, remark) FROM stdin;
\.


--
-- Data for Name: pro_dispatch; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.pro_dispatch (id, created_at, updated_at, deleted_at, tenant_id, order_id, order_no, process_id, process_name, station_id, station_name, assign_user_id, assign_user_name, quantity, status) FROM stdin;
\.


--
-- Data for Name: pro_production_order; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.pro_production_order (id, created_at, updated_at, deleted_at, tenant_id, order_no, sales_order_no, material_id, material_code, material_name, material_spec, unit, quantity, completed_qty, rejected_qty, workshop_id, workshop_name, line_id, line_name, route_id, bom_id, plan_start_date, plan_end_date, actual_start_date, actual_end_date, priority, status, remark) FROM stdin;
\.


--
-- Data for Name: pro_production_report; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.pro_production_report (id, created_at, updated_at, deleted_at, tenant_id, order_id, order_no, process_id, process_name, station_id, station_name, report_user_id, report_user_name, report_date, quantity, qualified_qty, rejected_qty, work_time, remark) FROM stdin;
\.


--
-- Data for Name: pro_sales_order; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.pro_sales_order (id, created_at, updated_at, deleted_at, tenant_id, order_no, customer_id, customer_name, order_date, delivery_date, order_type, priority, status, remark, order_id, material_id, material_code, material_name, quantity, unit, price, amount, shipped_qty) FROM stdin;
\.


--
-- Data for Name: pro_sales_order_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.pro_sales_order_item (id, created_at, updated_at, deleted_at, order_id, material_id, material_code, material_name, quantity, unit, price, amount, shipped_qty) FROM stdin;
\.


--
-- Data for Name: processes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.processes (id, created_at, updated_at, deleted_at, tenant_id, process_code, process_name, process_type, sequence, station_id, status) FROM stdin;
1	\N	\N	\N	1	PROC001	来料检验	检验	1	3	1
2	\N	\N	\N	1	PROC002	钣金冲压	加工	2	1	1
3	\N	\N	\N	1	PROC003	电机绕线	加工	3	4	1
4	\N	\N	\N	1	PROC004	总装	装配	4	2	1
5	\N	\N	\N	1	PROC005	电气测试	测试	5	5	1
6	\N	\N	\N	1	PROC006	成品检验	检验	6	3	1
7	\N	\N	\N	1	PROC007	包装	包装	7	6	1
\.


--
-- Data for Name: production_daily_report; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.production_daily_report (id, created_at, updated_at, deleted_at, tenant_id, report_date, workshop_id, workshop_name, production_order_count, completed_order_count, total_output_qty, qualified_qty, defect_qty, pass_rate, first_pass_rate, oee, output_per_hour, worker_count, working_hours, remark) FROM stdin;
\.


--
-- Data for Name: production_lines; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.production_lines (id, created_at, updated_at, deleted_at, tenant_id, line_code, line_name, workshop_id, line_type, status) FROM stdin;
\.


--
-- Data for Name: production_offline; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.production_offline (id, tenant_id, offline_code, work_order_id, work_order_code, product_id, product_code, product_name, offline_type, offline_reason, offline_qty, process_route_id, current_op_id, current_op_name, handle_method, handle_qty, handle_result, rework_order_id, scrap_qty, downgrade_qty, status, operator_id, operator_name, remark, created_by, updated_by, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: production_offline_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.production_offline_item (id, tenant_id, offline_id, serial_no, batch_no, offline_qty, handle_method, handle_qty, handle_result, remark, created_at) FROM stdin;
\.


--
-- Data for Name: production_orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.production_orders (id, created_at, updated_at, deleted_at, tenant_id, order_no, sales_order_no, material_id, material_code, material_name, material_spec, unit, quantity, completed_qty, rejected_qty, workshop_id, workshop_name, line_id, line_name, route_id, bom_id, plan_start_date, plan_end_date, actual_start_date, actual_end_date, priority, status, remark) FROM stdin;
\.


--
-- Data for Name: production_reports; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.production_reports (id, created_at, updated_at, deleted_at, tenant_id, order_id, order_no, process_id, process_name, station_id, station_name, report_user_id, report_user_name, report_date, quantity, qualified_qty, rejected_qty, work_time, remark) FROM stdin;
\.


--
-- Data for Name: qc_aql_level; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.qc_aql_level (id, created_at, updated_at, deleted_at, tenant_id, level, name, type, "order", status, remark) FROM stdin;
\.


--
-- Data for Name: qc_aql_table_row; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.qc_aql_table_row (id, created_at, updated_at, deleted_at, tenant_id, aql_level_id, aql_value, batch_min, batch_max, sample_size, ac, re) FROM stdin;
\.


--
-- Data for Name: qc_defect_code; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.qc_defect_code (id, created_at, updated_at, deleted_at, tenant_id, defect_code, defect_name, defect_type, severity, status) FROM stdin;
\.


--
-- Data for Name: qc_defect_record; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.qc_defect_record (id, created_at, updated_at, deleted_at, tenant_id, record_no, order_id, order_no, process_id, process_name, defect_code_id, defect_code, defect_name, quantity, handle_method, handle_user_id, handle_date, status) FROM stdin;
\.


--
-- Data for Name: qc_fqc; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.qc_fqc (id, created_at, updated_at, deleted_at, tenant_id, fqc_no, order_id, order_no, quantity, sample_size, qualified_qty, rejected_qty, check_user_id, check_user_name, check_date, result, remark) FROM stdin;
\.


--
-- Data for Name: qc_inspection_characteristic; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.qc_inspection_characteristic (id, code, name, type, spec_lower, spec_upper, usl, lsl, target, unit, aql, inspection_method, tenant_id, status) FROM stdin;
\.


--
-- Data for Name: qc_inspection_feature; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.qc_inspection_feature (id, created_at, updated_at, tenant_id, feature_code, feature_name, product_id, product_code, product_name, inspection_type, feature_type, technical_spec, lower_limit, upper_limit, unit, sample_size, gauges_method, aql_level, status, remark, created_by, updated_by) FROM stdin;
\.


--
-- Data for Name: qc_ipqc; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.qc_ipqc (id, created_at, updated_at, deleted_at, tenant_id, ipqc_no, order_id, order_no, process_id, process_name, quantity, sample_size, check_user_id, check_user_name, check_date, result, remark, ip_qc_no) FROM stdin;
\.


--
-- Data for Name: qc_iqc; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.qc_iqc (id, created_at, updated_at, deleted_at, tenant_id, iqc_no, supplier_id, supplier_name, material_id, material_code, material_name, quantity, unit, check_user_id, check_user_name, check_date, result, remark) FROM stdin;
\.


--
-- Data for Name: qc_iqc_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.qc_iqc_item (id, created_at, updated_at, deleted_at, iqc_id, check_item_id, check_item, check_standard, check_method, result, remark) FROM stdin;
\.


--
-- Data for Name: qc_ncr; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.qc_ncr (id, created_at, updated_at, deleted_at, tenant_id, ncr_no, defect_id, source_type, issue_desc, root_cause, corrective_action, preventive_action, verify_result, verify_user_id, verify_date, status) FROM stdin;
\.


--
-- Data for Name: qc_oqc; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.qc_oqc (id, created_at, updated_at, deleted_at, tenant_id, oqc_no, shipping_no, customer_id, customer_name, quantity, check_user_id, check_user_name, check_date, result, remark) FROM stdin;
\.


--
-- Data for Name: qc_sampling_plan; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.qc_sampling_plan (id, created_at, updated_at, deleted_at, tenant_id, code, name, inspection_type, aql_level_id, default_aql, min_batch_size, max_batch_size, status, remark) FROM stdin;
\.


--
-- Data for Name: qc_spc_data; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.qc_spc_data (id, created_at, updated_at, deleted_at, tenant_id, equipment_id, station_id, process_id, process_name, check_item, check_value, usl, lsl, cl, ucl, lcl, check_time) FROM stdin;
\.


--
-- Data for Name: qms_sampling_plan; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.qms_sampling_plan (id, created_at, updated_at, deleted_at, tenant_id, plan_code, plan_name, inspection_level, sample_type, aql, status, remark) FROM stdin;
\.


--
-- Data for Name: qms_sampling_record; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.qms_sampling_record (id, created_at, updated_at, deleted_at, tenant_id, plan_id, plan_code, inspection_id, batch_qty, sample_size, defect_count, ac_result, inspector, inspect_time) FROM stdin;
\.


--
-- Data for Name: qms_sampling_rule; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.qms_sampling_rule (id, created_at, updated_at, deleted_at, plan_id, batch_qty_from, batch_qty_to, sample_size, ac_accept, re_reject) FROM stdin;
\.


--
-- Data for Name: quality_weekly_report; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.quality_weekly_report (id, created_at, updated_at, deleted_at, tenant_id, report_year, report_week, start_date, end_date, workshop_id, workshop_name, total_inspection_qty, qualified_qty, defect_qty, pass_rate, iqc_insp_qty, iqc_qualified_qty, iqc_defect_qty, ip_qc_insp_qty, ip_qc_qualified_qty, ip_qc_defect_qty, fqc_insp_qty, fqc_qualified_qty, fqc_defect_qty, oqc_insp_qty, oqc_qualified_qty, oqc_defect_qty, ncr_count, customer_complaint_count, remark) FROM stdin;
\.


--
-- Data for Name: receive_order_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.receive_order_items (id, created_at, updated_at, deleted_at, receive_id, material_id, material_code, material_name, quantity, received_qty, unit, batch_no) FROM stdin;
\.


--
-- Data for Name: receive_orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.receive_orders (id, created_at, updated_at, deleted_at, tenant_id, receive_no, supplier_id, supplier_name, warehouse_id, receive_date, receive_user_id, status, remark) FROM stdin;
\.


--
-- Data for Name: resources; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.resources (id, created_at, updated_at, deleted_at, tenant_id, resource_code, resource_name, resource_type, workshop_id, capacity, unit, efficiency, status) FROM stdin;
\.


--
-- Data for Name: route_operations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.route_operations (id, created_at, updated_at, deleted_at, route_id, process_id, process_name, sequence, station_id, station_name, standard_time, queue_time, transport_time) FROM stdin;
1	\N	\N	\N	1	1	来料检验	1	3	检测工位	30	10	5
2	\N	\N	\N	1	2	钣金冲压	2	1	装配工位1	120	15	10
3	\N	\N	\N	1	3	电机绕线	3	4	一车间产线1工位1	90	10	8
4	\N	\N	\N	1	4	总装	4	2	装配工位2	60	10	5
5	\N	\N	\N	1	5	电气测试	5	5	一车间产线2工位2	45	10	5
6	\N	\N	\N	1	6	成品检验	6	3	检测工位	30	5	3
7	\N	\N	\N	1	7	包装	7	6	包装工位	20	5	3
8	\N	\N	\N	2	1	来料检验	1	3	检测工位	30	10	5
9	\N	\N	\N	2	2	钣金冲压	2	1	装配工位1	150	15	10
10	\N	\N	\N	2	3	电机绕线	3	4	一车间产线1工位1	120	10	8
11	\N	\N	\N	2	4	总装	4	2	装配工位2	80	10	5
12	\N	\N	\N	2	5	电气测试	5	5	一车间产线2工位2	60	10	5
13	\N	\N	\N	2	6	成品检验	6	3	检测工位	40	5	3
14	\N	\N	\N	2	7	包装	7	6	包装工位	25	5	3
\.


--
-- Data for Name: routes; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.routes (id, created_at, updated_at, deleted_at, tenant_id, route_code, route_name, material_id, version, status) FROM stdin;
1	\N	\N	\N	1	ROUTE-001	控制柜A生产工艺	27	V1.0	1
2	\N	\N	\N	1	ROUTE-002	控制柜B生产工艺	28	V1.0	1
\.


--
-- Data for Name: sales_order_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sales_order_items (id, created_at, updated_at, deleted_at, order_id, material_id, material_code, material_name, quantity, unit, price, amount, shipped_qty) FROM stdin;
\.


--
-- Data for Name: sales_orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sales_orders (id, created_at, updated_at, deleted_at, tenant_id, order_no, customer_id, customer_name, order_date, delivery_date, order_type, priority, status, remark) FROM stdin;
\.


--
-- Data for Name: schedule_plans; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schedule_plans (id, created_at, updated_at, deleted_at, tenant_id, plan_no, plan_type, start_date, end_date, algorithm, status, remark) FROM stdin;
1	2026-03-31 16:54:14.460546+08	2026-04-01 17:29:51.529267+08	\N	0		粗排	2026-03-31 08:00:00+08	2026-04-30 08:00:00+08	遗传	3	\N
\.


--
-- Data for Name: schedule_results; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schedule_results (id, created_at, updated_at, deleted_at, plan_id, order_id, order_no, sequence, line_id, line_name, station_id, station_name, plan_start_time, plan_end_time, actual_start_time, actual_end_time, status) FROM stdin;
\.


--
-- Data for Name: scp_customer_inquiry; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_customer_inquiry (id, created_at, updated_at, tenant_id, inquiry_no, inquiry_date, customer_id, customer_name, contact_person, contact_phone, contact_email, expected_date, valid_until, currency, status, quoted_amount, winner_supplier_id, remark, created_by) FROM stdin;
\.


--
-- Data for Name: scp_inquiry_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_inquiry_item (id, inquiry_id, line_no, material_id, material_code, material_name, specification, unit, required_qty, target_price, quoted_price, quoted_supplier_id, lead_time_days, remark, created_at) FROM stdin;
\.


--
-- Data for Name: scp_mrs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_mrs (id, created_at, updated_at, tenant_id, mrs_no, plan_month, status, source_type, source_no, total_items, total_qty, published_by, published_at, remark) FROM stdin;
\.


--
-- Data for Name: scp_mrs_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_mrs_item (id, created_at, updated_at, mrs_id, mrs_no, material_id, material_code, material_name, spec, unit, req_qty, on_hand_qty, short_qty, supplier_id, supplier_name, promise_date, status) FROM stdin;
\.


--
-- Data for Name: scp_po_change_log; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_po_change_log (id, po_id, change_type, change_field, old_value, new_value, changed_by, changed_by_name, change_time, reason, tenant_id, created_at) FROM stdin;
\.


--
-- Data for Name: scp_purchase_order; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_purchase_order (id, created_at, updated_at, tenant_id, po_no, po_type, supplier_id, supplier_code, supplier_name, contact_person, contact_phone, contact_email, order_date, promised_date, currency, payment_terms, tax_rate, total_amount, total_qty, approved_by, approved_time, approval_status, status, source_type, source_no, remark, close_reason) FROM stdin;
1	2026-04-14 16:43:02.652079+08	2026-04-14 16:43:02.652079+08	1	PO-2026-04001	STANDARD	1	SUP-TEST01	华东精密机械有限公司	张经理	13812345601	\N	2026-04-01	2026-04-15	CNY	月结30天	13.00	125500.00	100.000	\N	\N	APPROVED	ISSUED	MANUAL	\N	\N	\N
2	2026-04-14 16:43:02.652079+08	2026-04-14 16:43:02.652079+08	1	PO-2026-04002	STANDARD	2	SUP-TEST02	深圳创新电子科技有限公司	李总监	13812345602	\N	2026-04-02	2026-04-16	CNY	月结30天	13.00	89000.00	200.000	\N	\N	APPROVED	ISSUED	MANUAL	\N	\N	\N
3	2026-04-14 16:43:02.652079+08	2026-04-14 16:43:02.652079+08	1	PO-2026-04003	URGENT	3	SUP-TEST03	苏州工业园区华鑫塑胶制品厂	王厂长	13812345603	\N	2026-04-03	2026-04-12	CNY	月结30天	13.00	45600.00	500.000	\N	\N	APPROVED	PARTIAL	MANUAL	\N	\N	\N
4	2026-04-14 16:43:02.652079+08	2026-04-14 16:43:02.652079+08	1	PO-2026-04004	STANDARD	4	SUP-TEST04	杭州中达金属表面处理有限公司	陈工	13812345604	\N	2026-04-04	2026-04-18	CNY	月结30天	13.00	32000.00	50.000	\N	\N	APPROVED	ISSUED	MANUAL	\N	\N	\N
6	2026-04-14 16:43:02.652079+08	2026-04-14 16:43:02.652079+08	1	PO-2026-04006	LONG_TERM	6	SUP-TEST06	北京中关村软件股份有限公司	赵总	13812345606	\N	2026-04-06	2026-05-06	CNY	月结60天	13.00	150000.00	10.000	\N	\N	APPROVED	ISSUED	MANUAL	\N	\N	\N
7	2026-04-14 16:43:02.652079+08	2026-04-14 16:43:02.652079+08	1	PO-2026-04007	STANDARD	7	SUP-TEST07	天津港保税区润通物流有限公司	孙经理	13812345607	\N	2026-04-07	2026-04-22	CNY	月结30天	13.00	18000.00	20.000	\N	\N	APPROVED	ISSUED	MANUAL	\N	\N	\N
9	2026-04-14 16:43:02.652079+08	2026-04-14 16:43:02.652079+08	1	PO-2026-04009	URGENT	9	SUP-TEST09	武汉光谷激光设备有限公司	吴工	13812345609	\N	2026-04-09	2026-04-14	CNY	月结30天	13.00	98000.00	5.000	\N	\N	APPROVED	PARTIAL	MANUAL	\N	\N	\N
8	2026-04-14 16:43:02.652079+08	2026-04-14 17:26:33.4466+08	1	PO-2026-04008	STANDARD	8	SUP-TEST08	成都西南航空材料有限公司	周经理	13812345608	\N	2026-04-08	2026-04-23	CNY	月结30天	13.00	67500.00	80.000	1	2026-04-14 17:26:33.444917+08	APPROVED	APPROVED	MANUAL	\N	\N	\N
10	2026-04-14 16:43:02.652079+08	2026-04-16 14:52:21.396033+08	1	PO-2026-04010	STANDARD	10	SUP-TEST10	南京金陵化工有限公司	郑经理	13812345610	\N	2026-04-10	2026-04-25	CNY	月结30天	13.00	54200.00	150.000	\N	\N	APPROVED	RECEIVED	MANUAL	\N	\N	\N
5	2026-04-14 16:43:02.652079+08	2026-04-16 14:53:19.177125+08	1	PO-2026-04005	STANDARD	5	SUP-TEST05	广州岭南包装材料有限公司	刘经理	13812345605	\N	2026-04-05	2026-04-20	CNY	月结45天	13.00	28500.00	300.000	1	2026-04-16 14:35:25.248338+08	APPROVED	RECEIVED	MANUAL	\N	\N	\N
\.


--
-- Data for Name: scp_purchase_order_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_purchase_order_item (id, po_id, line_no, material_id, material_code, material_name, specification, unit, unit_price, order_qty, delivered_qty, received_qty, tax_amount, line_amount, promised_date, actual_delivery_date, batch_no, quality_require, package_require, is_gifted, status, remark, created_at) FROM stdin;
1	1	1	\N	MAT-RAW-001	钢板A3	1200*2400mm	PCS	850.0000	100.000	100.000	100.000	11050.00	96050.00	2026-04-15	\N	\N	\N	\N	0	COMPLETED	\N	\N
2	1	2	\N	MAT-RAW-002	钢板A4	1200*1800mm	PCS	295.0000	100.000	0.000	0.000	3835.00	29500.00	2026-04-15	\N	\N	\N	\N	0	PENDING	\N	\N
3	2	1	\N	MAT-RAW-003	铝合金板	1000*2000mm	PCS	380.0000	150.000	150.000	150.000	7410.00	57000.00	2026-04-16	\N	\N	\N	\N	0	COMPLETED	\N	\N
4	2	2	\N	MAT-RAW-004	不锈钢板	1500*3000mm	PCS	320.0000	100.000	50.000	50.000	4160.00	32000.00	2026-04-16	\N	\N	\N	\N	0	PARTIAL	\N	\N
5	3	1	\N	MAT-RAW-005	塑料粒子PP	通用级	KG	45.0000	500.000	300.000	300.000	1755.00	13500.00	2026-04-12	\N	\N	\N	\N	0	PARTIAL	\N	\N
6	3	2	\N	MAT-RAW-005	塑料粒子PP	增强级	KG	52.0000	400.000	0.000	0.000	2704.00	20800.00	2026-04-12	\N	\N	\N	\N	0	PENDING	\N	\N
7	4	1	\N	MAT-RAW-001	钢板A3	1200*2400mm	PCS	860.0000	50.000	50.000	50.000	5590.00	43000.00	2026-04-18	\N	\N	\N	\N	0	COMPLETED	\N	\N
8	5	1	\N	MAT-RAW-002	钢板A4	1200*1800mm	PCS	285.0000	100.000	0.000	0.000	3705.00	28500.00	2026-04-20	\N	\N	\N	\N	0	PENDING	\N	\N
9	6	1	\N	MAT-RAW-003	铝合金板	1000*2000mm	PCS	15000.0000	10.000	10.000	0.000	19500.00	150000.00	2026-05-06	\N	\N	\N	\N	0	PARTIAL	\N	\N
10	7	1	\N	MAT-RAW-001	钢板A3	1200*2400mm	PCS	900.0000	20.000	20.000	20.000	2340.00	18000.00	2026-04-22	\N	\N	\N	\N	0	COMPLETED	\N	\N
11	1	1	\N	MAT-RAW-001	钢板A3	1200*2400mm	PCS	850.0000	100.000	100.000	100.000	11050.00	96050.00	2026-04-15	\N	\N	\N	\N	0	COMPLETED	\N	\N
12	1	2	\N	MAT-RAW-002	钢板A4	1200*1800mm	PCS	295.0000	100.000	0.000	0.000	3835.00	29500.00	2026-04-15	\N	\N	\N	\N	0	PENDING	\N	\N
13	2	1	\N	MAT-RAW-003	铝合金板	1000*2000mm	PCS	380.0000	150.000	150.000	150.000	7410.00	57000.00	2026-04-16	\N	\N	\N	\N	0	COMPLETED	\N	\N
14	2	2	\N	MAT-RAW-004	不锈钢板	1500*3000mm	PCS	320.0000	100.000	50.000	50.000	4160.00	32000.00	2026-04-16	\N	\N	\N	\N	0	PARTIAL	\N	\N
15	3	1	\N	MAT-RAW-005	塑料粒子PP	通用级	KG	45.0000	500.000	300.000	300.000	1755.00	13500.00	2026-04-12	\N	\N	\N	\N	0	PARTIAL	\N	\N
16	3	2	\N	MAT-RAW-005	塑料粒子PP	增强级	KG	52.0000	400.000	0.000	0.000	2704.00	20800.00	2026-04-12	\N	\N	\N	\N	0	PENDING	\N	\N
17	4	1	\N	MAT-RAW-001	钢板A3	1200*2400mm	PCS	860.0000	50.000	50.000	50.000	5590.00	43000.00	2026-04-18	\N	\N	\N	\N	0	COMPLETED	\N	\N
18	5	1	\N	MAT-RAW-002	钢板A4	1200*1800mm	PCS	285.0000	100.000	0.000	0.000	3705.00	28500.00	2026-04-20	\N	\N	\N	\N	0	PENDING	\N	\N
19	6	1	\N	MAT-RAW-003	铝合金板	1000*2000mm	PCS	15000.0000	10.000	10.000	0.000	19500.00	150000.00	2026-05-06	\N	\N	\N	\N	0	PARTIAL	\N	\N
20	7	1	\N	MAT-RAW-001	钢板A3	1200*2400mm	PCS	900.0000	20.000	20.000	20.000	2340.00	18000.00	2026-04-22	\N	\N	\N	\N	0	COMPLETED	\N	\N
\.


--
-- Data for Name: scp_purchase_plan; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_purchase_plan (id, created_at, updated_at, tenant_id, plan_no, title, plan_type, plan_year, plan_month, quarter, status, total_items, total_amount, currency, department, submitter_id, submitted_at, confirmed_by, confirmed_at, published_by, published_at, closed_by, closed_at, close_reason, remark) FROM stdin;
\.


--
-- Data for Name: scp_purchase_plan_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_purchase_plan_item (id, created_at, updated_at, plan_id, plan_no, line_no, material_id, material_code, material_name, spec, unit, req_qty, unit_price, line_amount, supplier_id, supplier_code, supplier_name, req_delivery_date, promise_date, mrs_no, mrs_line_no, status, remark) FROM stdin;
\.


--
-- Data for Name: scp_qad_sync_log; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_qad_sync_log (id, tenant_id, sync_type, sync_direction, qad_doc_no, mom_doc_no, status, request_content, response_content, error_msg, sync_time, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: scp_quote_comparison; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_quote_comparison (id, created_at, updated_at, tenant_id, rfq_id, comparison_no, compared_by, compared_at, summary_data, recommendation, decision_remark) FROM stdin;
\.


--
-- Data for Name: scp_quote_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_quote_item (id, quote_id, rfq_line_id, material_id, material_code, material_name, unit, quoted_qty, unit_price, line_amount, delivery_date, lead_time_days, remark, created_at) FROM stdin;
1	1	1	\N	MAT-RAW-001	钢板A3	PCS	500.000	820.0000	410000.00	2026-05-01	15	\N	\N
2	1	2	\N	MAT-RAW-002	钢板A4	PCS	300.000	50.0000	15000.00	2026-05-01	15	\N	\N
3	2	1	\N	MAT-RAW-001	钢板A3	PCS	500.000	850.0000	425000.00	2026-05-01	20	\N	\N
4	2	2	\N	MAT-RAW-002	钢板A4	PCS	300.000	43.3300	13000.00	2026-05-01	20	\N	\N
5	3	3	\N	MAT-RAW-003	铝合金板	PCS	100.000	380.0000	38000.00	2026-04-15	10	\N	\N
6	4	4	\N	MAT-RAW-005	塑料粒子PP	KG	2000.000	43.0000	86000.00	2026-05-01	7	\N	\N
7	4	5	\N	MAT-RAW-005	塑料粒子PP	KG	1000.000	50.0000	5000.00	2026-05-01	7	\N	\N
8	5	7	\N	MAT-RAW-005	包装材料	PCS	5000.000	5.5000	27500.00	2026-05-01	5	\N	\N
9	6	9	\N	MAT-SVC-001	物流服务	次	12.000	5500.0000	66000.00	2026-05-01	3	\N	\N
10	7	10	\N	MAT-RAW-004	不锈钢板	PCS	100.000	285.0000	28500.00	2026-05-10	12	\N	\N
11	1	1	\N	MAT-RAW-001	钢板A3	PCS	500.000	820.0000	410000.00	2026-05-01	15	\N	\N
12	1	2	\N	MAT-RAW-002	钢板A4	PCS	300.000	50.0000	15000.00	2026-05-01	15	\N	\N
13	2	1	\N	MAT-RAW-001	钢板A3	PCS	500.000	850.0000	425000.00	2026-05-01	20	\N	\N
14	2	2	\N	MAT-RAW-002	钢板A4	PCS	300.000	43.3300	13000.00	2026-05-01	20	\N	\N
15	3	3	\N	MAT-RAW-003	铝合金板	PCS	100.000	380.0000	38000.00	2026-04-15	10	\N	\N
16	4	4	\N	MAT-RAW-005	塑料粒子PP	KG	2000.000	43.0000	86000.00	2026-05-01	7	\N	\N
17	4	5	\N	MAT-RAW-005	塑料粒子PP	KG	1000.000	50.0000	5000.00	2026-05-01	7	\N	\N
18	5	7	\N	MAT-RAW-005	包装材料	PCS	5000.000	5.5000	27500.00	2026-05-01	5	\N	\N
19	6	9	\N	MAT-SVC-001	物流服务	次	12.000	5500.0000	66000.00	2026-05-01	3	\N	\N
20	7	10	\N	MAT-RAW-004	不锈钢板	PCS	100.000	285.0000	28500.00	2026-05-10	12	\N	\N
\.


--
-- Data for Name: scp_rfq; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_rfq (id, created_at, updated_at, tenant_id, rfq_no, rfq_name, rfq_type, inquiry_date, deadline_date, currency, payment_terms, delivery_terms, quality_standard, status, awarded_supplier_id, awarded_total_amount, total_bids, is_evaluated, evaluation_by, evaluation_time, remark, created_by) FROM stdin;
1	2026-04-14 16:43:02.660487+08	2026-04-14 16:43:02.660487+08	1	RFQ-2026-04001	钢板年度采购询价	ANNUAL	2026-03-25	2026-04-10	CNY	月结30天	FOB上海港	国标GB/T3274-2017	CLOSED	\N	\N	5	1	\N	\N	\N	admin
2	2026-04-14 16:43:02.660487+08	2026-04-14 16:43:02.660487+08	1	RFQ-2026-04002	电子元器件紧急询价	QUICK	2026-04-01	2026-04-08	CNY	月结30天	EXW深圳	行业标准	AWARDED	\N	\N	3	1	\N	\N	\N	admin
3	2026-04-14 16:43:02.660487+08	2026-04-14 16:43:02.660487+08	1	RFQ-2026-04003	塑料粒子季度询价	STANDARD	2026-04-02	2026-04-18	CNY	月结30天	CIF广州	ISO标准	PUBLISHED	\N	\N	2	0	\N	\N	\N	admin
4	2026-04-14 16:43:02.660487+08	2026-04-14 16:43:02.660487+08	1	RFQ-2026-04004	表面处理服务询价	STANDARD	2026-04-03	2026-04-20	CNY	月结30天	DDP杭州	工艺标准	DRAFT	\N	\N	0	0	\N	\N	\N	admin
5	2026-04-14 16:43:02.660487+08	2026-04-14 16:43:02.660487+08	1	RFQ-2026-04005	包装材料询价	STANDARD	2026-04-04	2026-04-22	CNY	月结45天	FOB上海港	包装标准	PUBLISHED	\N	\N	1	0	\N	\N	\N	admin
6	2026-04-14 16:43:02.660487+08	2026-04-14 16:43:02.660487+08	1	RFQ-2026-04006	软件系统采购询价	STANDARD	2026-04-05	2026-04-25	CNY	年结	SAAS服务	甲方要求	DRAFT	\N	\N	0	0	\N	\N	\N	admin
7	2026-04-14 16:43:02.660487+08	2026-04-14 16:43:02.660487+08	1	RFQ-2026-04007	物流运输服务询价	STANDARD	2026-04-06	2026-04-23	CNY	月结30天	门到门	物流标准	PUBLISHED	\N	\N	2	0	\N	\N	\N	admin
8	2026-04-14 16:43:02.660487+08	2026-04-14 16:43:02.660487+08	1	RFQ-2026-04008	复合材料采购询价	ANNUAL	2026-04-07	2026-04-24	CNY	月结30天	FOB成都	国标	CLOSED	\N	\N	4	1	\N	\N	\N	admin
9	2026-04-14 16:43:02.660487+08	2026-04-14 16:43:02.660487+08	1	RFQ-2026-04009	激光设备询价	STANDARD	2026-04-08	2026-04-26	CNY	月结60天	CIF武汉	设备标准	PUBLISHED	\N	\N	1	0	\N	\N	\N	admin
10	2026-04-14 16:43:02.660487+08	2026-04-14 16:43:02.660487+08	1	RFQ-2026-04010	化工原料询价	STANDARD	2026-04-09	2026-04-27	CNY	月结30天	FOB南京	危化品标准	DRAFT	\N	\N	0	0	\N	\N	\N	admin
\.


--
-- Data for Name: scp_rfq_invite; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_rfq_invite (id, rfq_id, supplier_id, supplier_code, supplier_name, contact_person, contact_email, invite_date, response_status, quoted_date, tenant_id, created_at) FROM stdin;
2	1	1	SUP-TEST01	华东精密机械有限公司	张经理	zhang@example.com	2026-03-25	QUOTED	\N	1	\N
3	1	2	SUP-TEST02	深圳创新电子科技有限公司	李总监	li@example.com	2026-03-25	QUOTED	\N	1	\N
4	1	8	SUP-TEST08	成都西南航空材料有限公司	周经理	zhou@example.com	2026-03-25	DECLINED	\N	1	\N
5	2	2	SUP-TEST02	深圳创新电子科技有限公司	李总监	li@example.com	2026-04-01	QUOTED	\N	1	\N
6	2	9	SUP-TEST09	武汉光谷激光设备有限公司	吴工	wu@example.com	2026-04-01	QUOTED	\N	1	\N
7	3	3	SUP-TEST03	苏州工业园区华鑫塑胶制品厂	王厂长	wang@example.com	2026-04-02	QUOTED	\N	1	\N
8	3	5	SUP-TEST05	广州岭南包装材料有限公司	刘经理	liu@example.com	2026-04-02	QUOTED	\N	1	\N
9	4	4	SUP-TEST04	杭州中达金属表面处理有限公司	陈工	chen@example.com	2026-04-03	PENDING	\N	1	\N
10	5	5	SUP-TEST05	广州岭南包装材料有限公司	刘经理	liu@example.com	2026-04-04	QUOTED	\N	1	\N
11	7	7	SUP-TEST07	天津港保税区润通物流有限公司	孙经理	sun@example.com	2026-04-06	QUOTED	\N	1	\N
12	8	8	SUP-TEST08	成都西南航空材料有限公司	周经理	zhou@example.com	2026-04-07	QUOTED	\N	1	\N
13	8	1	SUP-TEST01	华东精密机械有限公司	张经理	zhang@example.com	2026-04-07	QUOTED	\N	1	\N
14	9	9	SUP-TEST09	武汉光谷激光设备有限公司	吴工	wu@example.com	2026-04-08	QUOTED	\N	1	\N
\.


--
-- Data for Name: scp_rfq_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_rfq_item (id, rfq_id, line_no, material_id, material_code, material_name, specification, unit, required_qty, target_price, market_price, requested_date, quality_require, remark, created_at) FROM stdin;
1	1	1	\N	MAT-RAW-001	钢板A3	1200*2400mm	PCS	500.000	800.0000	850.0000	2026-05-01	\N	\N	\N
2	1	2	\N	MAT-RAW-002	钢板A4	1200*1800mm	PCS	300.000	280.0000	295.0000	2026-05-01	\N	\N	\N
3	2	1	\N	MAT-RAW-003	铝合金板	1000*2000mm	PCS	100.000	350.0000	380.0000	2026-04-15	\N	\N	\N
4	3	1	\N	MAT-RAW-005	塑料粒子PP	通用级	KG	2000.000	42.0000	45.0000	2026-05-01	\N	\N	\N
5	3	2	\N	MAT-RAW-005	塑料粒子PP	增强级	KG	1000.000	50.0000	52.0000	2026-05-01	\N	\N	\N
6	4	1	\N	MAT-RAW-001	钢板A3	1200*2400mm	PCS	200.000	850.0000	\N	2026-05-15	\N	\N	\N
7	5	1	\N	MAT-RAW-005	包装材料	标准包装	PCS	5000.000	5.0000	5.5000	2026-05-01	\N	\N	\N
8	6	1	\N	MAT-SW-001	ERP软件	标准版	套	1.000	80000.0000	95000.0000	2026-06-01	\N	\N	\N
9	7	1	\N	MAT-SVC-001	物流服务	国内整车	次	12.000	5000.0000	5500.0000	2026-05-01	\N	\N	\N
10	8	1	\N	MAT-RAW-004	不锈钢板	1500*3000mm	PCS	100.000	280.0000	320.0000	2026-05-10	\N	\N	\N
11	1	1	\N	MAT-RAW-001	钢板A3	1200*2400mm	PCS	500.000	800.0000	850.0000	2026-05-01	\N	\N	\N
12	1	2	\N	MAT-RAW-002	钢板A4	1200*1800mm	PCS	300.000	280.0000	295.0000	2026-05-01	\N	\N	\N
13	2	1	\N	MAT-RAW-003	铝合金板	1000*2000mm	PCS	100.000	350.0000	380.0000	2026-04-15	\N	\N	\N
14	3	1	\N	MAT-RAW-005	塑料粒子PP	通用级	KG	2000.000	42.0000	45.0000	2026-05-01	\N	\N	\N
15	3	2	\N	MAT-RAW-005	塑料粒子PP	增强级	KG	1000.000	50.0000	52.0000	2026-05-01	\N	\N	\N
16	4	1	\N	MAT-RAW-001	钢板A3	1200*2400mm	PCS	200.000	850.0000	\N	2026-05-15	\N	\N	\N
17	5	1	\N	MAT-RAW-005	包装材料	标准包装	PCS	5000.000	5.0000	5.5000	2026-05-01	\N	\N	\N
18	6	1	\N	MAT-SW-001	ERP软件	标准版	套	1.000	80000.0000	95000.0000	2026-06-01	\N	\N	\N
19	7	1	\N	MAT-SVC-001	物流服务	国内整车	次	12.000	5000.0000	5500.0000	2026-05-01	\N	\N	\N
20	8	1	\N	MAT-RAW-004	不锈钢板	1500*3000mm	PCS	100.000	280.0000	320.0000	2026-05-10	\N	\N	\N
\.


--
-- Data for Name: scp_sales_order; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_sales_order (id, created_at, updated_at, tenant_id, so_no, so_type, customer_id, customer_code, customer_name, contact_person, contact_phone, contact_email, sales_person_id, sales_person_name, order_date, promised_date, currency, payment_terms, tax_rate, total_amount, total_qty, delivered_amount, delivered_qty, approved_by, approved_time, approval_status, status, source_type, source_no, delivery_address, delivery_warehouse_id, remark) FROM stdin;
1	2026-04-14 16:43:02.672739+08	2026-04-14 16:43:02.672739+08	1	SO-2026-04001	STANDARD	1	CUS001	杭州汽车配件有限公司	赵经理	13900010001	zhao@hzauto.com	\N	张三	2026-04-01	2026-04-20	CNY	月结30天	13.00	226000.00	200.000	0.00	0.000	\N	\N	APPROVED	CONFIRMED	MANUAL	\N	杭州市萧山区	\N	\N
2	2026-04-14 16:43:02.672739+08	2026-04-14 16:43:02.672739+08	1	SO-2026-04002	STANDARD	2	CUS002	上海机电设备有限公司	钱经理	13900010002	qian@shjd.com	\N	张三	2026-04-02	2026-04-22	CNY	月结30天	13.00	169000.00	150.000	0.00	0.000	\N	\N	APPROVED	CONFIRMED	MANUAL	\N	上海市闵行区	\N	\N
3	2026-04-14 16:43:02.672739+08	2026-04-14 16:43:02.672739+08	1	SO-2026-04003	URGENT	3	CUS003	宁波电子科技有限公司	孙经理	13900010003	sun@nbec.com	\N	李四	2026-04-03	2026-04-12	CNY	预付30%	13.00	85000.00	100.000	85000.00	100.000	\N	\N	APPROVED	SHIPPED	MANUAL	\N	宁波市鄞州区	\N	\N
4	2026-04-14 16:43:02.672739+08	2026-04-14 16:43:02.672739+08	1	SO-2026-04004	STANDARD	4	CUS004	温州电器股份有限公司	李经理	13900010004	li@wzelectric.com	\N	李四	2026-04-04	2026-04-25	CNY	月结45天	13.00	135000.00	120.000	45000.00	40.000	\N	\N	APPROVED	PARTIAL	MANUAL	\N	温州市乐清市	\N	\N
5	2026-04-14 16:43:02.672739+08	2026-04-14 16:43:02.672739+08	1	SO-2026-04005	DISTRIBUTION	1	CUS001	杭州汽车配件有限公司	赵经理	13900010001	zhao@hzauto.com	\N	王五	2026-04-05	2026-04-28	CNY	月结30天	13.00	95000.00	500.000	0.00	0.000	\N	\N	PENDING	PENDING	MANUAL	\N	杭州市萧山区	\N	\N
6	2026-04-14 16:43:02.672739+08	2026-04-14 16:43:02.672739+08	1	SO-2026-04006	STANDARD	2	CUS002	上海机电设备有限公司	钱经理	13900010002	qian@shjd.com	\N	王五	2026-04-06	2026-05-01	CNY	月结30天	13.00	280000.00	80.000	0.00	0.000	\N	\N	APPROVED	CONFIRMED	MANUAL	\N	上海市闵行区	\N	\N
7	2026-04-14 16:43:02.672739+08	2026-04-14 16:43:02.672739+08	1	SO-2026-04007	STANDARD	3	CUS003	宁波电子科技有限公司	孙经理	13900010003	sun@nbec.com	\N	张三	2026-04-07	2026-05-05	CNY	月结30天	13.00	78000.00	60.000	0.00	0.000	\N	\N	PENDING	PENDING	MANUAL	\N	宁波市鄞州区	\N	\N
8	2026-04-14 16:43:02.672739+08	2026-04-14 16:43:02.672739+08	1	SO-2026-04008	URGENT	4	CUS004	温州电器股份有限公司	李经理	13900010004	li@wzelectric.com	\N	李四	2026-04-08	2026-04-18	CNY	预付50%	13.00	156000.00	45.000	156000.00	45.000	\N	\N	APPROVED	SHIPPED	MANUAL	\N	温州市乐清市	\N	\N
9	2026-04-14 16:43:02.672739+08	2026-04-14 16:43:02.672739+08	1	SO-2026-04009	STANDARD	1	CUS001	杭州汽车配件有限公司	赵经理	13900010001	zhao@hzauto.com	\N	王五	2026-04-09	2026-05-10	CNY	月结30天	13.00	198000.00	180.000	0.00	0.000	\N	\N	APPROVED	CONFIRMED	MANUAL	\N	杭州市萧山区	\N	\N
10	2026-04-14 16:43:02.672739+08	2026-04-14 16:43:02.672739+08	1	SO-2026-04010	STANDARD	2	CUS002	上海机电设备有限公司	钱经理	13900010002	qian@shjd.com	\N	张三	2026-04-10	2026-05-15	CNY	月结30天	13.00	245000.00	220.000	0.00	0.000	\N	\N	PENDING	PENDING	MANUAL	\N	上海市闵行区	\N	\N
\.


--
-- Data for Name: scp_sales_order_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_sales_order_item (id, so_id, line_no, material_id, material_code, material_name, specification, unit, unit_price, order_qty, delivered_qty, shipped_qty, tax_amount, line_amount, promised_date, actual_delivery_date, production_order_id, status, remark, created_at) FROM stdin;
1	1	1	\N	MAT-RAW-001	钢板A3	1200*2400mm	PCS	980.0000	150.000	0.000	0.000	19110.00	147000.00	2026-04-20	\N	\N	PENDING	\N	\N
2	1	2	\N	MAT-RAW-002	钢板A4	1200*1800mm	PCS	580.0000	50.000	0.000	0.000	3770.00	29000.00	2026-04-20	\N	\N	PENDING	\N	\N
3	1	3	\N	MAT-RAW-003	铝合金板	1000*2000mm	PCS	1000.0000	50.000	0.000	0.000	6500.00	50000.00	2026-04-20	\N	\N	PENDING	\N	\N
4	2	1	\N	MAT-RAW-001	钢板A3	1200*2400mm	PCS	990.0000	100.000	0.000	0.000	12870.00	99000.00	2026-04-22	\N	\N	PENDING	\N	\N
5	2	2	\N	MAT-RAW-004	不锈钢板	1500*3000mm	PCS	700.0000	100.000	0.000	0.000	9100.00	70000.00	2026-04-22	\N	\N	PENDING	\N	\N
6	3	1	\N	MAT-RAW-003	铝合金板	1000*2000mm	PCS	850.0000	100.000	100.000	100.000	11050.00	85000.00	2026-04-12	\N	\N	COMPLETED	\N	\N
7	4	1	\N	MAT-RAW-001	钢板A3	1200*2400mm	PCS	1000.0000	80.000	40.000	40.000	10400.00	80000.00	2026-04-25	\N	\N	PARTIAL	\N	\N
8	4	2	\N	MAT-RAW-002	钢板A4	1200*1800mm	PCS	550.0000	40.000	0.000	0.000	2860.00	22000.00	2026-04-25	\N	\N	PENDING	\N	\N
9	4	3	\N	MAT-RAW-003	铝合金板	1000*2000mm	PCS	825.0000	33.000	0.000	0.000	3543.75	27272.73	2026-04-25	\N	\N	PENDING	\N	\N
10	5	1	\N	MAT-RAW-005	塑料粒子PP	通用级	KG	190.0000	500.000	0.000	0.000	12350.00	95000.00	2026-04-28	\N	\N	PENDING	\N	\N
11	1	1	\N	MAT-RAW-001	钢板A3	1200*2400mm	PCS	980.0000	150.000	0.000	0.000	19110.00	147000.00	2026-04-20	\N	\N	PENDING	\N	\N
12	1	2	\N	MAT-RAW-002	钢板A4	1200*1800mm	PCS	580.0000	50.000	0.000	0.000	3770.00	29000.00	2026-04-20	\N	\N	PENDING	\N	\N
13	1	3	\N	MAT-RAW-003	铝合金板	1000*2000mm	PCS	1000.0000	50.000	0.000	0.000	6500.00	50000.00	2026-04-20	\N	\N	PENDING	\N	\N
14	2	1	\N	MAT-RAW-001	钢板A3	1200*2400mm	PCS	990.0000	100.000	0.000	0.000	12870.00	99000.00	2026-04-22	\N	\N	PENDING	\N	\N
15	2	2	\N	MAT-RAW-004	不锈钢板	1500*3000mm	PCS	700.0000	100.000	0.000	0.000	9100.00	70000.00	2026-04-22	\N	\N	PENDING	\N	\N
16	3	1	\N	MAT-RAW-003	铝合金板	1000*2000mm	PCS	850.0000	100.000	100.000	100.000	11050.00	85000.00	2026-04-12	\N	\N	COMPLETED	\N	\N
17	4	1	\N	MAT-RAW-001	钢板A3	1200*2400mm	PCS	1000.0000	80.000	40.000	40.000	10400.00	80000.00	2026-04-25	\N	\N	PARTIAL	\N	\N
18	4	2	\N	MAT-RAW-002	钢板A4	1200*1800mm	PCS	550.0000	40.000	0.000	0.000	2860.00	22000.00	2026-04-25	\N	\N	PENDING	\N	\N
19	4	3	\N	MAT-RAW-003	铝合金板	1000*2000mm	PCS	825.0000	33.000	0.000	0.000	3543.75	27272.73	2026-04-25	\N	\N	PENDING	\N	\N
20	5	1	\N	MAT-RAW-005	塑料粒子PP	通用级	KG	190.0000	500.000	0.000	0.000	12350.00	95000.00	2026-04-28	\N	\N	PENDING	\N	\N
\.


--
-- Data for Name: scp_so_change_log; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_so_change_log (id, so_id, change_type, change_field, old_value, new_value, changed_by, changed_by_name, change_time, reason, tenant_id, created_at) FROM stdin;
\.


--
-- Data for Name: scp_supplier_bank; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_supplier_bank (id, created_at, updated_at, tenant_id, supplier_id, supplier_code, supplier_name, bank_name, bank_code, branch_name, branch_code, account_type, account_no, account_name, currency, is_primary, is_active, status, remark) FROM stdin;
\.


--
-- Data for Name: scp_supplier_contact; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_supplier_contact (id, created_at, updated_at, tenant_id, supplier_id, supplier_code, supplier_name, name, gender, department, "position", phone, mobile, email, wechat, qq, is_primary, is_active, remark) FROM stdin;
\.


--
-- Data for Name: scp_supplier_delivery_record; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_supplier_delivery_record (id, supplier_id, po_no, po_line_id, promised_date, actual_delivery_date, delivery_status, delay_days, is_penalty_applied, penalty_amount, remark, tenant_id, created_at) FROM stdin;
\.


--
-- Data for Name: scp_supplier_grade_standard; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_supplier_grade_standard (id, grade, grade_name, min_score, max_score, on_time_rate_threshold, quality_rate_threshold, is_active, tenant_id, created_at) FROM stdin;
\.


--
-- Data for Name: scp_supplier_kpi; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_supplier_kpi (id, created_at, updated_at, tenant_id, supplier_id, supplier_code, supplier_name, evaluation_month, evaluation_date, evaluated_by, evaluated_by_name, on_time_delivery_rate, total_delivery_orders, on_time_delivery_count, avg_delay_days, quality_pass_rate, total_iqc_orders, passed_iqc_orders, defect_parts_count, defect_rate, price_competitiveness, last_purchase_price, market_avg_price, total_score, grade, rank_position, evaluation_remark) FROM stdin;
1	\N	\N	1	1	SUP-TEST01	华东精密机械有限公司	2026-03	2026-04-05	1	系统管理员	95.50	20	19	0.50	98.00	15	15	5	0.0025	92.00	850.0000	870.0000	95.20	A	1	\N
2	\N	\N	1	2	SUP-TEST02	深圳创新电子科技有限公司	2026-03	2026-04-05	1	系统管理员	88.00	15	13	1.20	95.50	10	10	8	0.0060	88.00	380.0000	395.0000	89.80	B	3	\N
3	\N	\N	1	3	SUP-TEST03	苏州工业园区华鑫塑胶制品厂	2026-03	2026-04-05	1	系统管理员	92.00	12	11	0.80	96.00	8	8	3	0.0030	90.00	45.0000	47.0000	92.50	A	2	\N
4	\N	\N	1	4	SUP-TEST04	杭州中达金属表面处理有限公司	2026-03	2026-04-05	1	系统管理员	85.00	8	7	1.50	93.00	6	6	4	0.0055	85.00	\N	\N	86.30	B	5	\N
5	\N	\N	1	5	SUP-TEST05	广州岭南包装材料有限公司	2026-03	2026-04-05	1	系统管理员	90.00	10	9	1.00	94.00	7	7	4	0.0045	88.00	5.5000	5.8000	90.50	B	4	\N
6	\N	\N	1	6	SUP-TEST06	北京中关村软件股份有限公司	2026-03	2026-04-05	1	系统管理员	100.00	3	3	0.00	100.00	2	2	0	0.0000	95.00	95000.0000	98000.0000	98.00	A	1	\N
7	\N	\N	1	7	SUP-TEST07	天津港保税区润通物流有限公司	2026-03	2026-04-05	1	系统管理员	93.00	6	6	0.30	\N	\N	\N	\N	\N	90.00	5500.0000	5800.0000	91.50	B	2	\N
8	\N	\N	1	8	SUP-TEST08	成都西南航空材料有限公司	2026-03	2026-04-05	1	系统管理员	87.00	9	8	1.80	91.00	5	5	6	0.0080	82.00	320.0000	350.0000	85.20	C	6	\N
9	\N	\N	1	9	SUP-TEST09	武汉光谷激光设备有限公司	2026-03	2026-04-05	1	系统管理员	80.00	5	4	2.50	88.00	4	4	5	0.0100	78.00	98000.0000	105000.0000	81.00	C	7	\N
10	\N	\N	1	10	SUP-TEST10	南京金陵化工有限公司	2026-03	2026-04-05	1	系统管理员	91.00	11	10	0.70	95.00	9	9	5	0.0045	89.00	\N	\N	91.00	B	3	\N
11	\N	\N	1	1	SUP-TEST01	华东精密机械有限公司	2026-03	2026-04-05	1	系统管理员	95.50	20	19	0.50	98.00	15	15	5	0.0025	92.00	850.0000	870.0000	95.20	A	1	\N
12	\N	\N	1	2	SUP-TEST02	深圳创新电子科技有限公司	2026-03	2026-04-05	1	系统管理员	88.00	15	13	1.20	95.50	10	10	8	0.0060	88.00	380.0000	395.0000	89.80	B	3	\N
13	\N	\N	1	3	SUP-TEST03	苏州工业园区华鑫塑胶制品厂	2026-03	2026-04-05	1	系统管理员	92.00	12	11	0.80	96.00	8	8	3	0.0030	90.00	45.0000	47.0000	92.50	A	2	\N
14	\N	\N	1	4	SUP-TEST04	杭州中达金属表面处理有限公司	2026-03	2026-04-05	1	系统管理员	85.00	8	7	1.50	93.00	6	6	4	0.0055	85.00	\N	\N	86.30	B	5	\N
15	\N	\N	1	5	SUP-TEST05	广州岭南包装材料有限公司	2026-03	2026-04-05	1	系统管理员	90.00	10	9	1.00	94.00	7	7	4	0.0045	88.00	5.5000	5.8000	90.50	B	4	\N
16	\N	\N	1	6	SUP-TEST06	北京中关村软件股份有限公司	2026-03	2026-04-05	1	系统管理员	100.00	3	3	0.00	100.00	2	2	0	0.0000	95.00	95000.0000	98000.0000	98.00	A	1	\N
17	\N	\N	1	7	SUP-TEST07	天津港保税区润通物流有限公司	2026-03	2026-04-05	1	系统管理员	93.00	6	6	0.30	\N	\N	\N	\N	\N	90.00	5500.0000	5800.0000	91.50	B	2	\N
18	\N	\N	1	8	SUP-TEST08	成都西南航空材料有限公司	2026-03	2026-04-05	1	系统管理员	87.00	9	8	1.80	91.00	5	5	6	0.0080	82.00	320.0000	350.0000	85.20	C	6	\N
19	\N	\N	1	9	SUP-TEST09	武汉光谷激光设备有限公司	2026-03	2026-04-05	1	系统管理员	80.00	5	4	2.50	88.00	4	4	5	0.0100	78.00	98000.0000	105000.0000	81.00	C	7	\N
20	\N	\N	1	10	SUP-TEST10	南京金陵化工有限公司	2026-03	2026-04-05	1	系统管理员	91.00	11	10	0.70	95.00	9	9	5	0.0045	89.00	\N	\N	91.00	B	3	\N
\.


--
-- Data for Name: scp_supplier_purchase_info; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_supplier_purchase_info (id, supplier_id, payment_terms, credit_limit, tax_rate, min_order_amount, lead_time_days, supplier_grade, is_preferred, is_blacklist, blacklist_reason, cooperation_start_date, cooperation_end_date, total_cooperation_amount, tenant_id, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: scp_supplier_quality_record; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_supplier_quality_record (id, supplier_id, material_code, iqc_record_id, iqc_no, inspect_date, inspect_qty, qualified_qty, defect_qty, defect_rate, defect_types, is_chargeback, chargeback_amount, remark, tenant_id, created_at) FROM stdin;
\.


--
-- Data for Name: scp_supplier_quote; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.scp_supplier_quote (id, created_at, updated_at, tenant_id, rfq_id, rfq_no, supplier_id, supplier_code, supplier_name, quote_no, quote_date, valid_until, currency, payment_terms, delivery_days, total_amount, is_accepted, is_lowest, rank_position, quote_status, evaluation_score, evaluation_result, evaluator_id, evaluation_time, evaluation_remark, remark) FROM stdin;
1	2026-04-14 16:43:02.667345+08	2026-04-14 16:43:02.667345+08	1	1	RFQ-2026-04001	1	SUP-TEST01	华东精密机械有限公司	QUO-2026-04001	2026-04-08	2026-05-08	CNY	月结30天	15	425000.00	1	1	1	SUBMITTED	\N	\N	\N	\N	\N	\N
2	2026-04-14 16:43:02.667345+08	2026-04-14 16:43:02.667345+08	1	1	RFQ-2026-04001	2	SUP-TEST02	深圳创新电子科技有限公司	QUO-2026-04002	2026-04-08	2026-05-08	CNY	月结30天	20	438000.00	0	0	2	SUBMITTED	\N	\N	\N	\N	\N	\N
3	2026-04-14 16:43:02.667345+08	2026-04-14 16:43:02.667345+08	1	2	RFQ-2026-04002	2	SUP-TEST02	深圳创新电子科技有限公司	QUO-2026-04003	2026-04-06	2026-05-06	CNY	月结30天	10	38000.00	1	1	1	SUBMITTED	\N	\N	\N	\N	\N	\N
4	2026-04-14 16:43:02.667345+08	2026-04-14 16:43:02.667345+08	1	3	RFQ-2026-04003	3	SUP-TEST03	苏州工业园区华鑫塑胶制品厂	QUO-2026-04004	2026-04-15	2026-05-15	CNY	月结30天	7	94000.00	0	1	1	SUBMITTED	\N	\N	\N	\N	\N	\N
5	2026-04-14 16:43:02.667345+08	2026-04-14 16:43:02.667345+08	1	5	RFQ-2026-04005	5	SUP-TEST05	广州岭南包装材料有限公司	QUO-2026-04005	2026-04-18	2026-05-18	CNY	月结45天	5	27500.00	1	1	1	SUBMITTED	\N	\N	\N	\N	\N	\N
6	2026-04-14 16:43:02.667345+08	2026-04-14 16:43:02.667345+08	1	7	RFQ-2026-04007	7	SUP-TEST07	天津港保税区润通物流有限公司	QUO-2026-04006	2026-04-20	2026-05-20	CNY	月结30天	3	66000.00	0	1	1	SUBMITTED	\N	\N	\N	\N	\N	\N
7	2026-04-14 16:43:02.667345+08	2026-04-14 16:43:02.667345+08	1	8	RFQ-2026-04008	8	SUP-TEST08	成都西南航空材料有限公司	QUO-2026-04007	2026-04-20	2026-05-20	CNY	月结30天	12	28500.00	1	1	1	SUBMITTED	\N	\N	\N	\N	\N	\N
8	2026-04-14 16:43:02.667345+08	2026-04-14 16:43:02.667345+08	1	8	RFQ-2026-04008	1	SUP-TEST01	华东精密机械有限公司	QUO-2026-04008	2026-04-20	2026-05-20	CNY	月结30天	15	29500.00	0	0	2	SUBMITTED	\N	\N	\N	\N	\N	\N
9	2026-04-14 16:43:02.667345+08	2026-04-14 16:43:02.667345+08	1	9	RFQ-2026-04009	9	SUP-TEST09	武汉光谷激光设备有限公司	QUO-2026-04009	2026-04-22	2026-05-22	CNY	月结60天	30	98000.00	1	1	1	SUBMITTED	\N	\N	\N	\N	\N	\N
10	2026-04-14 16:43:02.667345+08	2026-04-14 16:43:02.667345+08	1	3	RFQ-2026-04003	5	SUP-TEST05	广州岭南包装材料有限公司	QUO-2026-04010	2026-04-16	2026-05-16	CNY	月结45天	5	98000.00	0	0	2	REVISED	\N	\N	\N	\N	\N	\N
\.


--
-- Data for Name: serial_numbers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.serial_numbers (id, created_at, updated_at, deleted_at, tenant_id, serial_number, material_id, material_code, material_name, batch_no, line_id, line_name, order_id, order_no, production_date, status) FROM stdin;
\.


--
-- Data for Name: shifts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.shifts (id, created_at, updated_at, deleted_at, tenant_id, shift_code, shift_name, start_time, end_time, break_start, break_end, status) FROM stdin;
\.


--
-- Data for Name: spc_data; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.spc_data (id, created_at, updated_at, deleted_at, tenant_id, equipment_id, station_id, process_id, process_name, check_item, check_value, usl, lsl, cl, ucl, lcl, check_time) FROM stdin;
\.


--
-- Data for Name: stock_checks; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.stock_checks (id, created_at, updated_at, deleted_at, tenant_id, check_no, warehouse_id, check_date, check_user_id, status, remark) FROM stdin;
\.


--
-- Data for Name: supplier_asn_items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.supplier_asn_items (id, asn_id, line_no, material_code, material_name, spec, unit, batch_no, qty, qualified_qty, price, amount, packing_qty, packing_unit, received_qty, tenant_id, created_at) FROM stdin;
\.


--
-- Data for Name: supplier_asns; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.supplier_asns (id, asn_no, supplier_id, supplier_code, supplier_name, delivery_type, delivery_date, delivery_time_start, delivery_time_end, warehouse_code, contact_person, contact_phone, status, total_qty, total_amount, remark, tenant_id, created_by, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: sys_dept; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_dept (id, created_at, updated_at, deleted_at, tenant_id, parent_id, dept_name, dept_code, dept_sort, leader, phone, email, status) FROM stdin;
1	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	0	总公司	HQ	1	张总	13800000001	zong@fengmei.com	1
2	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	1	信息中心	IT	1	李经理	13800000002	it@fengmei.com	1
3	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	1	生产部	PROD	2	王经理	13800000003	prod@fengmei.com	1
4	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	1	质量部	QA	3	陈经理	13800000004	qa@fengmei.com	1
5	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	1	采购部	PUR	4	刘经理	13800000005	pur@fengmei.com	1
6	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	1	销售部	SALE	5	周经理	13800000006	sale@fengmei.com	1
7	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	1	财务部	FIN	6	吴经理	13800000007	fin@fengmei.com	1
8	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	1	人力资源部	HR	7	郑经理	13800000008	hr@fengmei.com	1
9	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	3	一车间	PROD-01	1	赵主任	13800000009	prod01@fengmei.com	1
10	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	3	二车间	PROD-02	2	钱主任	13800000010	prod02@fengmei.com	1
11	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	3	三车间	PROD-03	3	孙主任	13800000011	prod03@fengmei.com	1
12	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	3	四车间	PROD-04	4	李主任	13800000012	prod04@fengmei.com	1
13	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	4	来料检验组	QA-IQC	1	周检验	13800000013	iqc@fengmei.com	1
14	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	4	过程检验组	QA-IPQC	2	吴检验	13800000014	ipqc@fengmei.com	1
15	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	4	成品检验组	QA-FQC	3	郑检验	13800000015	fqc@fengmei.com	1
16	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	4	出货检验组	QA-OQC	4	冯检验	13800000016	oqc@fengmei.com	1
17	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	2	开发组	IT-DEV	1	褚开发	13800000017	dev@fengmei.com	1
18	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	2	运维组	IT-OPS	2	戚运维	13800000018	ops@fengmei.com	1
19	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	2	网络组	IT-NET	3	谢网络	13800000019	net@fengmei.com	1
20	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	5	采购一组	PUR-01	1	葛采购	13800000020	pur01@fengmei.com	1
42	2026-04-04 00:22:35.173841+08	2026-04-04 00:22:35.173841+08	\N	1	1	实施部	20001	0			\N	1
43	2026-04-04 00:23:09.198042+08	2026-04-04 00:23:09.198042+08	\N	1	0	运营部	1001	0			\N	1
41	2026-04-03 22:48:48.013547+08	2026-04-03 22:48:48.013547+08	2026-04-04 00:52:02.015296	1	1	实施部	20001	0			\N	1
\.


--
-- Data for Name: sys_dict_data; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_dict_data (id, created_at, updated_at, deleted_at, dict_sort, dict_label, dict_value, dict_type, dict_key, css_class, list_class, is_default, status, remark) FROM stdin;
1	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	1	正常	1	sys_user_status	\N	\N	\N	0	1	\N
2	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	2	停用	0	sys_user_status	\N	\N	\N	0	1	\N
3	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	1	正常	1	sys_role_status	\N	\N	\N	0	1	\N
4	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	2	停用	0	sys_role_status	\N	\N	\N	0	1	\N
5	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	1	运行	1	equ_status	\N	\N	\N	0	1	\N
6	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	2	待机	2	equ_status	\N	\N	\N	0	1	\N
7	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	3	故障	3	equ_status	\N	\N	\N	0	1	\N
8	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	4	维修	4	equ_status	\N	\N	\N	0	1	\N
9	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	5	报废	5	equ_status	\N	\N	\N	0	1	\N
10	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	1	生产设备	production	equ_type	\N	\N	\N	0	1	\N
11	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	2	检测设备	inspection	equ_type	\N	\N	\N	0	1	\N
12	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	3	辅助设备	auxiliary	equ_type	\N	\N	\N	0	1	\N
13	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	1	正常	1	check_result	\N	\N	\N	0	1	\N
14	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	2	异常	2	check_result	\N	\N	\N	0	1	\N
15	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	1	日常保养	daily	maint_type	\N	\N	\N	0	1	\N
16	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	2	一级保养	level1	maint_type	\N	\N	\N	0	1	\N
17	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	3	二级保养	level2	maint_type	\N	\N	\N	0	1	\N
18	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	1	待维修	1	repair_status	\N	\N	\N	0	1	\N
19	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	2	维修中	2	repair_status	\N	\N	\N	0	1	\N
20	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	3	已完成	3	repair_status	\N	\N	\N	0	1	\N
21	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	1	待生产	1	pro_order_status	\N	\N	\N	0	1	\N
22	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	2	生产中	2	pro_order_status	\N	\N	\N	0	1	\N
23	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	3	已完成	3	pro_order_status	\N	\N	\N	0	1	\N
24	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	4	已取消	4	pro_order_status	\N	\N	\N	0	1	\N
25	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	1	待确认	1	sales_order_status	\N	\N	\N	0	1	\N
26	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	2	已确认	2	sales_order_status	\N	\N	\N	0	1	\N
27	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	3	生产中	3	sales_order_status	\N	\N	\N	0	1	\N
28	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	4	已完成	4	sales_order_status	\N	\N	\N	0	1	\N
29	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	5	已关闭	5	sales_order_status	\N	\N	\N	0	1	\N
30	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	1	待开始	1	dispatch_status	\N	\N	\N	0	1	\N
31	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	2	进行中	2	dispatch_status	\N	\N	\N	0	1	\N
32	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	3	已完成	3	dispatch_status	\N	\N	\N	0	1	\N
33	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	1	正常	1	warehouse_status	\N	\N	\N	0	1	\N
34	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	2	停用	0	warehouse_status	\N	\N	\N	0	1	\N
35	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	1	原料仓	raw	warehouse_type	\N	\N	\N	0	1	\N
36	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	2	成品仓	finished	warehouse_type	\N	\N	\N	0	1	\N
37	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	3	线边仓	inline	warehouse_type	\N	\N	\N	0	1	\N
38	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	4	工具仓	tool	warehouse_type	\N	\N	\N	0	1	\N
39	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	1	电	electricity	energy_type	\N	\N	\N	0	1	\N
40	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	2	水	water	energy_type	\N	\N	\N	0	1	\N
41	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	3	气	gas	energy_type	\N	\N	\N	0	1	\N
42	2026-03-31 10:48:01.595598+08	2026-03-31 10:48:01.595598+08	\N	4	蒸汽	steam	energy_type	\N	\N	\N	0	1	\N
\.


--
-- Data for Name: sys_dict_type; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_dict_type (id, created_at, updated_at, deleted_at, dict_name, dict_type, status, remark) FROM stdin;
1	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	用户状态	sys_user_status	1	\N
2	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	角色状态	sys_role_status	1	\N
3	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	菜单状态	sys_menu_status	1	\N
4	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	部门状态	sys_dept_status	1	\N
5	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	岗位状态	sys_post_status	1	\N
6	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	设备状态	equ_status	1	\N
7	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	设备类型	equ_type	1	\N
8	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	点检结果	check_result	1	\N
9	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	保养类型	maint_type	1	\N
10	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	维修状态	repair_status	1	\N
11	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	生产订单状态	pro_order_status	1	\N
12	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	销售订单状态	sales_order_status	1	\N
13	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	派工状态	dispatch_status	1	\N
14	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	仓库状态	warehouse_status	1	\N
15	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	仓库类型	warehouse_type	1	\N
16	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	库位状态	location_status	1	\N
17	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	库存状态	inventory_status	1	\N
18	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	单据状态	doc_status	1	\N
19	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	检验结果	qc_result	1	\N
20	2026-03-31 10:48:01.59385+08	2026-03-31 10:48:01.59385+08	\N	能源类型	energy_type	1	\N
\.


--
-- Data for Name: sys_interface_config; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_interface_config (id, created_at, updated_at, deleted_at, tenant_id, name, code, category, description, direction, method, base_url, path, auth_type, auth_config, request_content_type, request_body_template, response_format, timeout, retry_count, retry_interval, source_type, source_table, source_api, source_filter, source_fields, primary_key, incremental_field, incremental_window, status, remark) FROM stdin;
\.


--
-- Data for Name: sys_interface_exec_log; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_interface_exec_log (id, created_at, updated_at, deleted_at, interface_config_id, config_name, trigger_type, start_time, end_time, duration, request_url, request_method, request_headers, request_body, response_body, response_code, status, error_message, retry_count, records_processed) FROM stdin;
\.


--
-- Data for Name: sys_interface_field_map; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_interface_field_map (id, created_at, updated_at, deleted_at, interface_config_id, field_name, field_type, map_type, map_value, required, default_value, transform_func, sort_order) FROM stdin;
\.


--
-- Data for Name: sys_interface_trigger; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_interface_trigger (id, created_at, updated_at, deleted_at, interface_config_id, trigger_type, cron_expr, event_source, payload_template, condition, fallback_minutes, status) FROM stdin;
\.


--
-- Data for Name: sys_login_log; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_login_log (id, created_at, updated_at, deleted_at, tenant_id, username, ip, login_location, browser, os, status, msg, login_time) FROM stdin;
1	2026-03-31 08:00:00+08	2026-03-31 08:00:00+08	\N	1	admin	192.168.1.100	浙江省杭州市	Chrome	Windows 11	0	登录成功	2026-03-31 08:00:00
2	2026-03-31 08:15:00+08	2026-03-31 08:15:00+08	\N	1	admin	192.168.1.101	浙江省杭州市	Firefox	Windows 10	0	登录成功	2026-03-31 08:15:00
3	2026-03-31 08:30:00+08	2026-03-31 08:30:00+08	\N	1	user01	192.168.1.102	浙江省宁波市	Safari	macOS	0	登录成功	2026-03-31 08:30:00
4	2026-03-31 09:00:00+08	2026-03-31 09:00:00+08	\N	2	hangzhou01	192.168.2.100	浙江省杭州市	Chrome	Windows 11	0	登录成功	2026-03-31 09:00:00
5	2026-03-31 09:15:00+08	2026-03-31 09:15:00+08	\N	2	hangzhou02	192.168.2.101	浙江省杭州市	Edge	Windows 11	0	登录成功	2026-03-31 09:15:00
6	2026-03-31 09:30:00+08	2026-03-31 09:30:00+08	\N	3	ningbo01	192.168.3.100	浙江省宁波市	Chrome	Ubuntu	0	登录成功	2026-03-31 09:30:00
7	2026-03-31 10:00:00+08	2026-03-31 10:00:00+08	\N	1	admin	192.168.1.100	浙江省杭州市	Chrome	Windows 11	1	密码错误	2026-03-31 10:00:00
8	2026-03-31 10:05:00+08	2026-03-31 10:05:00+08	\N	1	unknown	192.168.1.105	浙江省温州市	Chrome	Windows 11	1	用户不存在	2026-03-31 10:05:00
9	2026-03-31 10:30:00+08	2026-03-31 10:30:00+08	\N	4	wenzhou01	192.168.4.100	浙江省温州市	Firefox	Windows 10	0	登录成功	2026-03-31 10:30:00
10	2026-03-31 11:00:00+08	2026-03-31 11:00:00+08	\N	5	jiaxing01	192.168.5.100	浙江省嘉兴市	Chrome	Windows 11	0	登录成功	2026-03-31 11:00:00
11	2026-03-31 11:30:00+08	2026-03-31 11:30:00+08	\N	6	huzhou01	192.168.6.100	浙江省湖州市	Safari	macOS	0	登录成功	2026-03-31 11:30:00
12	2026-03-31 13:00:00+08	2026-03-31 13:00:00+08	\N	1	admin	192.168.1.100	浙江省杭州市	Chrome	Windows 11	0	登录成功	2026-03-31 13:00:00
13	2026-03-31 13:30:00+08	2026-03-31 13:30:00+08	\N	7	shaoxing01	192.168.7.100	浙江省绍兴市	Chrome	Windows 11	0	登录成功	2026-03-31 13:30:00
14	2026-03-31 14:00:00+08	2026-03-31 14:00:00+08	\N	8	jinhua01	192.168.8.100	浙江省金华市	Edge	Windows 11	0	登录成功	2026-03-31 14:00:00
15	2026-03-31 14:30:00+08	2026-03-31 14:30:00+08	\N	9	quzhou01	192.168.9.100	浙江省衢州市	Chrome	Ubuntu	0	登录成功	2026-03-31 14:30:00
16	2026-03-31 15:00:00+08	2026-03-31 15:00:00+08	\N	10	zhoushan01	192.168.10.100	浙江省舟山市	Firefox	Windows 10	0	登录成功	2026-03-31 15:00:00
17	2026-03-31 15:30:00+08	2026-03-31 15:30:00+08	\N	11	taizhou01	192.168.11.100	浙江省台州市	Chrome	Windows 11	0	登录成功	2026-03-31 15:30:00
18	2026-03-31 16:00:00+08	2026-03-31 16:00:00+08	\N	12	lishui01	192.168.12.100	浙江省丽水市	Safari	macOS	0	登录成功	2026-03-31 16:00:00
19	2026-03-31 16:30:00+08	2026-03-31 16:30:00+08	\N	13	shanghai01	192.168.13.100	上海市	Chrome	Windows 11	0	登录成功	2026-03-31 16:30:00
20	2026-03-31 17:00:00+08	2026-03-31 17:00:00+08	\N	14	jiangsu01	192.168.14.100	江苏省南京市	Chrome	Windows 11	0	登录成功	2026-03-31 17:00:00
21	2026-03-31 17:30:00+08	2026-03-31 17:30:00+08	\N	15	anhui01	192.168.15.100	安徽省合肥市	Edge	Windows 11	0	登录成功	2026-03-31 17:30:00
22	2026-03-31 18:00:00+08	2026-03-31 18:00:00+08	\N	1	admin	192.168.1.100	浙江省杭州市	Chrome	Windows 11	0	登录成功	2026-03-31 18:00:00
23	2026-03-30 08:00:00+08	2026-03-30 08:00:00+08	\N	1	admin	192.168.1.100	浙江省杭州市	Chrome	Windows 11	0	登录成功	2026-03-30 08:00:00
24	2026-03-30 18:00:00+08	2026-03-30 18:00:00+08	\N	1	admin	192.168.1.100	浙江省杭州市	Chrome	Windows 11	0	登录成功	2026-03-30 18:00:00
25	2026-03-29 08:00:00+08	2026-03-29 08:00:00+08	\N	1	admin	192.168.1.100	浙江省杭州市	Chrome	Windows 11	0	登录成功	2026-03-29 08:00:00
26	2026-04-22 11:14:52.575337+08	2026-04-22 11:14:52.575337+08	\N	1	admin	::1	\N			1	登录成功	2026-04-22 11:14:52.556414
27	2026-04-23 08:45:43.181689+08	2026-04-23 08:45:43.181689+08	\N	1	admin	::1	\N			1	登录成功	2026-04-23 08:45:43.032926
\.


--
-- Data for Name: sys_menu; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_menu (id, created_at, updated_at, deleted_at, tenant_id, parent_id, menu_name, menu_type, path, component, perms, icon, sort, visible, status, is_frame, is_cache) FROM stdin;
4	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	0	生产执行	M	/production	\N		List	4	1	1	0	0
5	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	0	设备管理	M	/equipment	\N		Monitor	5	1	1	0	0
7	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	0	质量管理	M	/quality	\N		CircleCheck	7	1	1	0	0
9	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	0	追溯管理	M	/trace	\N		Search	9	1	1	0	0
10	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	0	能源管理	M	/energy	\N		Lightning	10	1	1	0	0
101	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	2	用户管理	C	/system/user	system/UserList.vue	system:user:list	User	1	1	1	0	0
102	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	2	角色管理	C	/system/role	system/RoleList.vue	system:role:list	Key	2	1	1	0	0
104	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	2	部门管理	C	/system/dept	system/DeptList.vue	system:dept:list	OfficeBuilding	4	1	1	0	0
106	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	2	岗位管理	C	/system/post	system/PostList.vue	system:post:list	Postcard	6	1	1	0	0
108	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	2	登录日志	C	/system/login-log	system/LoginLogList.vue	system:loginlog:list	Key	8	1	1	0	0
109	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	2	操作日志	C	/system/oper-log	system/OperLogList.vue	system:operlog:list	Document	9	1	1	0	0
110	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	2	系统配置	C	/system/config	system/SystemConfig.vue	system:config:list	Setting	10	1	1	0	0
201	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	3	物料管理	C	/mdm/material	mdm/MaterialList.vue	mdm:material:list	Box	1	1	1	0	0
202	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	3	车间管理	C	/mdm/workshop	mdm/WorkshopList.vue	mdm:workshop:list	OfficeBuilding	2	1	1	0	0
203	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	3	生产线管理	C	/mdm/line	mdm/LineList.vue	mdm:line:list	Connection	3	1	1	0	0
204	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	3	工位管理	C	/mdm/workstation	mdm/WorkstationList.vue	mdm:workstation:list	Grid	4	1	1	0	0
206	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	3	BOM管理	C	/mdm/bom	mdm/BomList.vue	mdm:bom:list	Files	6	1	1	0	0
207	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	3	工序管理	C	/mdm/operation	mdm/OperationList.vue	mdm:operation:list	Operation	7	1	1	0	0
301	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	4	生产工单	C	/production/order	production/ProductionOrderList.vue	production:order:list	List	1	1	1	0	0
302	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	4	销售订单	C	/production/sales-order	production/SalesOrderList.vue	production:salesorder:list	Document	2	1	1	0	0
303	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	4	生产报工	C	/production/report	production/ReportList.vue	production:report:list	DocumentCheck	3	1	1	0	0
304	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	4	派工	C	/production/dispatch	production/DispatchList.vue	production:dispatch:list	Tickets	4	1	1	0	0
401	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	5	设备台账	C	/equipment	equipment/EquipmentList.vue	equipment:list:list	Monitor	1	1	1	0	0
402	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	5	设备点检	C	/equipment/check	equipment/CheckList.vue	equipment:check:list	Check	2	1	1	0	0
403	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	5	设备保养	C	/equipment/maintenance	equipment/MaintenanceList.vue	equipment:maintenance:list	Tools	3	1	1	0	0
404	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	5	设备维修	C	/equipment/repair	equipment/RepairList.vue	equipment:repair:list	Tool	4	1	1	0	0
405	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	5	备件管理	C	/equipment/spare	equipment/SparePartList.vue	equipment:spare:list	Box	5	1	1	0	0
501	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	6	仓库管理	C	/wms/warehouse	wms/WarehouseList.vue	wms:warehouse:list	House	1	1	1	0	0
502	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	6	库位管理	C	/wms/location	wms/LocationList.vue	wms:location:list	Location	2	1	1	0	0
503	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	6	库存管理	C	/wms/inventory	wms/InventoryList.vue	wms:inventory:list	Box	3	1	1	0	0
601	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	7	IQC检验	C	/quality/iqc	quality/IQCList.vue	quality:iqc:list	CircleCheck	1	1	1	0	0
602	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	7	IPQC检验	C	/quality/ipqc	quality/IPQCList.vue	quality:ipqc:list	CircleCheck	2	1	1	0	0
603	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	7	FQC检验	C	/quality/fqc	quality/FQCList.vue	quality:fqc:list	CircleCheck	3	1	1	0	0
103	2026-01-01 08:00:00+08	2026-04-08 16:14:31.768656+08	\N	1	2	菜单管理	C	/system/menu	system/MenuList.vue	system:menu:list	Menu	0	1	1	0	0
107	2026-01-01 08:00:00+08	2026-04-08 16:14:50.313018+08	\N	1	2	租户管理	C	/system/tenant	system/TenantList.vue	system:tenant:list	Building	3	1	1	0	0
105	2026-01-01 08:00:00+08	2026-04-08 16:15:28.467696+08	\N	1	2	字典管理	C	/system/dict	system/DictList.vue	system:dict:list	Document	8	1	1	0	0
2	2026-01-01 08:00:00+08	2026-04-10 14:00:19.228858+08	\N	1	0	系统管理	M	/system	\N		Setting	12	1	1	0	0
3	2026-01-01 08:00:00+08	2026-04-10 14:00:47.133821+08	\N	1	0	主数据	M	/mdm	\N		Box	11	1	1	0	0
6	2026-01-01 08:00:00+08	2026-04-10 14:01:07.007602+08	\N	1	0	仓储管理	M	/wms	\N		House	2	1	1	0	0
8	2026-01-01 08:00:00+08	2026-04-10 14:01:33.367754+08	\N	1	0	APS计划	M	/aps	\N		Calendar	3	1	1	0	0
1	2026-01-01 08:00:00+08	2026-04-16 10:10:35.781276+08	\N	1	0	首页	C	/dashboard	\N		HomeFilled	3	1	1	0	0
604	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	7	OQC检验	C	/quality/oqc	quality/OQCList.vue	quality:oqc:list	CircleCheck	4	1	1	0	0
605	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	7	缺陷代码	C	/quality/defect-code	quality/DefectCodeList.vue	quality:defectcode:list	Warning	5	1	1	0	0
606	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	7	缺陷记录	C	/quality/defect-record	quality/DefectRecordList.vue	quality:defectrecord:list	Document	6	1	1	0	0
607	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	7	NCR处理	C	/quality/ncr	quality/NCRList.vue	quality:ncr:list	Close	7	1	1	0	0
608	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	7	SPC数据	C	/quality/spc	quality/SPCList.vue	quality:spc:list	DataLine	8	1	1	0	0
701	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	8	MPS计划	C	/aps/mps	aps/MPSList.vue	aps:mps:list	Calendar	1	1	1	0	0
702	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	8	MRP计划	C	/aps/mrp	aps/MRPList.vue	aps:mrp:list	Grid	2	1	1	0	0
703	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	8	排程计划	C	/aps/schedule	aps/ScheduleList.vue	aps:schedule:list	List	3	1	1	0	0
801	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	9	追溯查询	C	/trace/query	trace/TraceQuery.vue	trace:query:list	Search	1	1	1	0	0
802	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	9	安东呼叫	C	/trace/andon	trace/AndonCall.vue	trace:andon:list	Bell	2	1	1	0	0
901	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	10	能源监控	C	/energy/monitor	energy/EnergyMonitor.vue	energy:monitor:list	Lightning	1	1	1	0	0
205	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	3	班次管理	C	/mdm/mdm-shift	mdm/ShiftList.vue	mdm:shift:list	Clock	5	1	1	0	0
505	2026-04-04 22:50:54.270317+08	2026-04-08 16:11:27.976644+08	\N	1	6	发货单	C	delivery	wms/DeliveryOrderList.vue	wms:deliveryorder:list	Upload	5	1	1	0	0
210	2026-04-04 22:50:54.26492+08	2026-04-09 08:56:26.816874+08	\N	1	3	客户管理	C	customer	mdm/CustomerList.vue	mdm:customer:list	User	10	0	1	0	0
209	2026-04-04 22:50:54.260905+08	2026-04-08 16:15:08.913502+08	\N	1	3	物料分类	C	material-category	mdm/MaterialCategoryList.vue	mdm:materialcategory:list	Folder	9	0	1	0	0
10001	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	101	新增用户	F		\N	system:user:add	#	1	0	1	0	0
10002	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	101	修改用户	F		\N	system:user:edit	#	2	0	1	0	0
10003	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	101	删除用户	F		\N	system:user:delete	#	3	0	1	0	0
10004	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	101	重置密码	F		\N	system:user:resetPwd	#	4	0	1	0	0
10005	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	101	导出用户	F		\N	system:user:export	#	5	0	1	0	0
10006	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	102	新增角色	F		\N	system:role:add	#	1	0	1	0	0
10007	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	102	修改角色	F		\N	system:role:edit	#	2	0	1	0	0
10008	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	102	删除角色	F		\N	system:role:delete	#	3	0	1	0	0
10009	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	102	分配权限	F		\N	system:role:assign	#	4	0	1	0	0
10010	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	103	新增菜单	F		\N	system:menu:add	#	1	0	1	0	0
10011	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	103	修改菜单	F		\N	system:menu:edit	#	2	0	1	0	0
10012	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	103	删除菜单	F		\N	system:menu:delete	#	3	0	1	0	0
10013	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	104	新增部门	F		\N	system:dept:add	#	1	0	1	0	0
10014	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	104	修改部门	F		\N	system:dept:edit	#	2	0	1	0	0
10015	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	104	删除部门	F		\N	system:dept:delete	#	3	0	1	0	0
10016	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	105	新增字典	F		\N	system:dict:add	#	1	0	1	0	0
10017	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	105	修改字典	F		\N	system:dict:edit	#	2	0	1	0	0
504	2026-04-04 22:50:54.268996+08	2026-04-08 16:11:22.610648+08	\N	1	6	收货单	C	receive	wms/ReceiveOrderList.vue	wms:receiveorder:list	Download	4	1	1	0	0
307	2026-04-04 22:50:54.265891+08	2026-04-08 17:20:39.598569+08	\N	1	4	工单变更	C	order	production/OrderChangeList.vue	production:orderchange:list	Edit	7	0	1	0	0
409	2026-04-04 22:50:54.266859+08	2026-04-08 17:21:53.267016+08	\N	1	5	OEE分析	C	oee	equipment/OEELIst.vue	equipment:oee:list	DataLine	9	0	1	0	0
704	2026-04-04 22:50:54.267931+08	2026-04-08 16:11:42.926837+08	\N	1	8	工作中心	C	work-center	aps/WorkCenterList.vue	aps:workcenter:list	OfficeBuilding	4	1	1	0	0
10018	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	105	删除字典	F		\N	system:dict:delete	#	3	0	1	0	0
10019	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	106	新增岗位	F		\N	system:post:add	#	1	0	1	0	0
10020	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	106	修改岗位	F		\N	system:post:edit	#	2	0	1	0	0
10021	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	106	删除岗位	F		\N	system:post:delete	#	3	0	1	0	0
10022	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	107	新增租户	F		\N	system:tenant:add	#	1	0	1	0	0
10023	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	107	修改租户	F		\N	system:tenant:edit	#	2	0	1	0	0
10024	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	107	删除租户	F		\N	system:tenant:delete	#	3	0	1	0	0
10025	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	108	导出日志	F		\N	system:loginlog:export	#	1	0	1	0	0
10026	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	108	清空日志	F		\N	system:loginlog:clean	#	2	0	1	0	0
10027	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	109	导出日志	F		\N	system:operlog:export	#	1	0	1	0	0
10101	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	201	新增物料	F		\N	mdm:material:add	#	1	0	1	0	0
10102	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	201	修改物料	F		\N	mdm:material:edit	#	2	0	1	0	0
10103	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	201	删除物料	F		\N	mdm:material:delete	#	3	0	1	0	0
10104	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	201	导入物料	F		\N	mdm:material:import	#	4	0	1	0	0
10105	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	201	导出物料	F		\N	mdm:material:export	#	5	0	1	0	0
10106	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	209	新增分类	F		\N	mdm:materialcategory:add	#	1	0	1	0	0
10107	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	209	修改分类	F		\N	mdm:materialcategory:edit	#	2	0	1	0	0
10108	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	209	删除分类	F		\N	mdm:materialcategory:delete	#	3	0	1	0	0
10109	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	202	新增车间	F		\N	mdm:workshop:add	#	1	0	1	0	0
10110	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	202	修改车间	F		\N	mdm:workshop:edit	#	2	0	1	0	0
10111	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	202	删除车间	F		\N	mdm:workshop:delete	#	3	0	1	0	0
10112	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	203	新增生产线	F		\N	mdm:line:add	#	1	0	1	0	0
10113	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	203	修改生产线	F		\N	mdm:line:edit	#	2	0	1	0	0
10114	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	203	删除生产线	F		\N	mdm:line:delete	#	3	0	1	0	0
10115	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	204	新增工位	F		\N	mdm:workstation:add	#	1	0	1	0	0
10116	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	204	修改工位	F		\N	mdm:workstation:edit	#	2	0	1	0	0
10117	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	204	删除工位	F		\N	mdm:workstation:delete	#	3	0	1	0	0
10118	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	205	新增班次	F		\N	mdm:shift:add	#	1	0	1	0	0
10119	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	205	修改班次	F		\N	mdm:shift:edit	#	2	0	1	0	0
10120	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	205	删除班次	F		\N	mdm:shift:delete	#	3	0	1	0	0
10121	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	206	新增BOM	F		\N	mdm:bom:add	#	1	0	1	0	0
10122	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	206	修改BOM	F		\N	mdm:bom:edit	#	2	0	1	0	0
10123	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	206	删除BOM	F		\N	mdm:bom:delete	#	3	0	1	0	0
10124	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	207	新增工序	F		\N	mdm:operation:add	#	1	0	1	0	0
10125	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	207	修改工序	F		\N	mdm:operation:edit	#	2	0	1	0	0
10126	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	207	删除工序	F		\N	mdm:operation:delete	#	3	0	1	0	0
10127	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	210	新增客户	F		\N	mdm:customer:add	#	1	0	1	0	0
10128	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	210	修改客户	F		\N	mdm:customer:edit	#	2	0	1	0	0
10129	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	210	删除客户	F		\N	mdm:customer:delete	#	3	0	1	0	0
10301	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	301	新增工单	F		\N	production:order:add	#	1	0	1	0	0
10302	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	301	修改工单	F		\N	production:order:edit	#	2	0	1	0	0
10303	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	301	删除工单	F		\N	production:order:delete	#	3	0	1	0	0
10304	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	301	审核工单	F		\N	production:order:approve	#	4	0	1	0	0
10305	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	301	下达工单	F		\N	production:order:release	#	5	0	1	0	0
10306	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	302	新增订单	F		\N	production:salesorder:add	#	1	0	1	0	0
10307	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	302	修改订单	F		\N	production:salesorder:edit	#	2	0	1	0	0
10308	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	302	删除订单	F		\N	production:salesorder:delete	#	3	0	1	0	0
10309	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	303	新增报工	F		\N	production:report:add	#	1	0	1	0	0
10310	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	303	修改报工	F		\N	production:report:edit	#	2	0	1	0	0
10311	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	303	删除报工	F		\N	production:report:delete	#	3	0	1	0	0
10312	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	304	新增派工	F		\N	production:dispatch:add	#	1	0	1	0	0
10313	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	304	修改派工	F		\N	production:dispatch:edit	#	2	0	1	0	0
10314	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	304	删除派工	F		\N	production:dispatch:delete	#	3	0	1	0	0
10315	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	307	新增变更	F		\N	production:orderchange:add	#	1	0	1	0	0
10316	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	307	修改变更	F		\N	production:orderchange:edit	#	2	0	1	0	0
10317	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	307	删除变更	F		\N	production:orderchange:delete	#	3	0	1	0	0
10318	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	307	审核变更	F		\N	production:orderchange:approve	#	4	0	1	0	0
10401	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	401	新增设备	F		\N	equipment:list:add	#	1	0	1	0	0
10402	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	401	修改设备	F		\N	equipment:list:edit	#	2	0	1	0	0
10403	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	401	删除设备	F		\N	equipment:list:delete	#	3	0	1	0	0
10404	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	402	新增点检	F		\N	equipment:check:add	#	1	0	1	0	0
10405	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	402	修改点检	F		\N	equipment:check:edit	#	2	0	1	0	0
10406	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	402	删除点检	F		\N	equipment:check:delete	#	3	0	1	0	0
10407	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	403	新增保养	F		\N	equipment:maintenance:add	#	1	0	1	0	0
10408	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	403	修改保养	F		\N	equipment:maintenance:edit	#	2	0	1	0	0
10409	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	403	删除保养	F		\N	equipment:maintenance:delete	#	3	0	1	0	0
10410	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	404	新增维修	F		\N	equipment:repair:add	#	1	0	1	0	0
10411	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	404	修改维修	F		\N	equipment:repair:edit	#	2	0	1	0	0
10412	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	404	删除维修	F		\N	equipment:repair:delete	#	3	0	1	0	0
10413	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	405	新增备件	F		\N	equipment:spare:add	#	1	0	1	0	0
10414	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	405	修改备件	F		\N	equipment:spare:edit	#	2	0	1	0	0
10415	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	405	删除备件	F		\N	equipment:spare:delete	#	3	0	1	0	0
10416	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	409	导出OEE	F		\N	equipment:oee:export	#	1	0	1	0	0
10501	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	501	新增仓库	F		\N	wms:warehouse:add	#	1	0	1	0	0
10502	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	501	修改仓库	F		\N	wms:warehouse:edit	#	2	0	1	0	0
10503	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	501	删除仓库	F		\N	wms:warehouse:delete	#	3	0	1	0	0
10504	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	502	新增库位	F		\N	wms:location:add	#	1	0	1	0	0
10505	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	502	修改库位	F		\N	wms:location:edit	#	2	0	1	0	0
10506	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	502	删除库位	F		\N	wms:location:delete	#	3	0	1	0	0
10507	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	503	导出库存	F		\N	wms:inventory:export	#	1	0	1	0	0
10508	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	503	盘点	F		\N	wms:inventory:check	#	2	0	1	0	0
10509	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	504	新增收货	F		\N	wms:receiveorder:add	#	1	0	1	0	0
10510	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	504	修改收货	F		\N	wms:receiveorder:edit	#	2	0	1	0	0
10511	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	504	删除收货	F		\N	wms:receiveorder:delete	#	3	0	1	0	0
10512	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	504	审核收货	F		\N	wms:receiveorder:approve	#	4	0	1	0	0
10513	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	505	新增发货	F		\N	wms:deliveryorder:add	#	1	0	1	0	0
10514	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	505	修改发货	F		\N	wms:deliveryorder:edit	#	2	0	1	0	0
10515	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	505	删除发货	F		\N	wms:deliveryorder:delete	#	3	0	1	0	0
10516	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	505	审核发货	F		\N	wms:deliveryorder:approve	#	4	0	1	0	0
10601	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	601	新增检验	F		\N	quality:iqc:add	#	1	0	1	0	0
10602	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	601	修改检验	F		\N	quality:iqc:edit	#	2	0	1	0	0
10603	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	601	删除检验	F		\N	quality:iqc:delete	#	3	0	1	0	0
10604	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	601	审核检验	F		\N	quality:iqc:approve	#	4	0	1	0	0
10605	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	602	新增检验	F		\N	quality:ipqc:add	#	1	0	1	0	0
10606	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	602	修改检验	F		\N	quality:ipqc:edit	#	2	0	1	0	0
10607	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	602	删除检验	F		\N	quality:ipqc:delete	#	3	0	1	0	0
10608	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	603	新增检验	F		\N	quality:fqc:add	#	1	0	1	0	0
10609	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	603	修改检验	F		\N	quality:fqc:edit	#	2	0	1	0	0
10610	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	603	删除检验	F		\N	quality:fqc:delete	#	3	0	1	0	0
10611	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	604	新增检验	F		\N	quality:oqc:add	#	1	0	1	0	0
10612	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	604	修改检验	F		\N	quality:oqc:edit	#	2	0	1	0	0
10613	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	604	删除检验	F		\N	quality:oqc:delete	#	3	0	1	0	0
10614	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	605	新增代码	F		\N	quality:defectcode:add	#	1	0	1	0	0
10615	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	605	修改代码	F		\N	quality:defectcode:edit	#	2	0	1	0	0
10616	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	605	删除代码	F		\N	quality:defectcode:delete	#	3	0	1	0	0
10617	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	606	新增记录	F		\N	quality:defectrecord:add	#	1	0	1	0	0
10618	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	606	修改记录	F		\N	quality:defectrecord:edit	#	2	0	1	0	0
10619	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	606	删除记录	F		\N	quality:defectrecord:delete	#	3	0	1	0	0
10620	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	607	新增NCR	F		\N	quality:ncr:add	#	1	0	1	0	0
10621	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	607	修改NCR	F		\N	quality:ncr:edit	#	2	0	1	0	0
10622	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	607	删除NCR	F		\N	quality:ncr:delete	#	3	0	1	0	0
10623	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	608	新增数据	F		\N	quality:spc:add	#	1	0	1	0	0
10624	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	608	修改数据	F		\N	quality:spc:edit	#	2	0	1	0	0
10625	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	608	删除数据	F		\N	quality:spc:delete	#	3	0	1	0	0
10701	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	701	新增MPS	F		\N	aps:mps:add	#	1	0	1	0	0
10702	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	701	修改MPS	F		\N	aps:mps:edit	#	2	0	1	0	0
10703	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	701	删除MPS	F		\N	aps:mps:delete	#	3	0	1	0	0
10704	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	701	下达MPS	F		\N	aps:mps:release	#	4	0	1	0	0
10705	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	702	新增MRP	F		\N	aps:mrp:add	#	1	0	1	0	0
10706	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	702	修改MRP	F		\N	aps:mrp:edit	#	2	0	1	0	0
10707	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	702	删除MRP	F		\N	aps:mrp:delete	#	3	0	1	0	0
10708	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	702	执行MRP	F		\N	aps:mrp:execute	#	4	0	1	0	0
10709	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	703	新增排程	F		\N	aps:schedule:add	#	1	0	1	0	0
10710	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	703	修改排程	F		\N	aps:schedule:edit	#	2	0	1	0	0
10711	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	703	删除排程	F		\N	aps:schedule:delete	#	3	0	1	0	0
10712	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	703	执行排程	F		\N	aps:schedule:execute	#	4	0	1	0	0
10713	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	704	新增工作中心	F		\N	aps:workcenter:add	#	1	0	1	0	0
10714	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	704	修改工作中心	F		\N	aps:workcenter:edit	#	2	0	1	0	0
10715	2026-04-04 23:00:50.561878+08	2026-04-04 23:00:50.561878+08	\N	1	704	删除工作中心	F		\N	aps:workcenter:delete	#	3	0	1	0	0
10417	2026-04-04 23:44:35.068274+08	2026-04-04 23:44:35.068274+08	\N	1	409	新增OEE	F		\N	equipment:oee:add	#	1	0	1	0	0
10418	2026-04-04 23:44:35.068274+08	2026-04-04 23:44:35.068274+08	\N	1	409	查看OEE	F		\N	equipment:oee:view	#	2	0	1	0	0
10419	2026-04-04 23:44:35.068274+08	2026-04-04 23:44:35.068274+08	\N	1	409	删除OEE	F		\N	equipment:oee:delete	#	3	0	1	0	0
10720	2026-04-04 23:45:07.855003+08	2026-04-04 23:45:07.855003+08	\N	1	701	运行MPS	F		\N	aps:mps:run	#	2	0	1	0	0
10721	2026-04-04 23:45:07.855003+08	2026-04-04 23:45:07.855003+08	\N	1	701	MPS明细	F		\N	aps:mps:detail	#	3	0	1	0	0
10420	2026-04-04 23:45:07.855003+08	2026-04-04 23:45:07.855003+08	\N	1	404	开始维修	F		\N	equipment:repair:start	#	3	0	1	0	0
10421	2026-04-04 23:45:07.855003+08	2026-04-04 23:45:07.855003+08	\N	1	404	完成维修	F		\N	equipment:repair:complete	#	4	0	1	0	0
10520	2026-04-04 23:45:07.855003+08	2026-04-04 23:45:07.855003+08	\N	1	503	新增库存	F		\N	wms:inventory:add	#	1	0	1	0	0
10521	2026-04-04 23:45:07.855003+08	2026-04-04 23:45:07.855003+08	\N	1	503	修改库存	F		\N	wms:inventory:edit	#	2	0	1	0	0
10522	2026-04-04 23:45:07.855003+08	2026-04-04 23:45:07.855003+08	\N	1	503	删除库存	F		\N	wms:inventory:delete	#	3	0	1	0	0
10523	2026-04-04 23:45:07.855003+08	2026-04-04 23:45:07.855003+08	\N	1	504	查看收货	F		\N	wms:receiveorder:view	#	5	0	1	0	0
10524	2026-04-04 23:45:07.855003+08	2026-04-04 23:45:07.855003+08	\N	1	504	确认收货	F		\N	wms:receiveorder:confirm	#	6	0	1	0	0
10525	2026-04-04 23:45:07.855003+08	2026-04-04 23:45:07.855003+08	\N	1	505	查看发货	F		\N	wms:deliveryorder:view	#	5	0	1	0	0
10526	2026-04-04 23:45:07.855003+08	2026-04-04 23:45:07.855003+08	\N	1	505	确认发货	F		\N	wms:deliveryorder:confirm	#	6	0	1	0	0
10730	2026-04-04 23:45:37.715022+08	2026-04-04 23:45:37.715022+08	\N	1	702	缺料分析	F		\N	aps:mrp:shortage	#	4	0	1	0	0
10731	2026-04-04 23:45:37.715022+08	2026-04-04 23:45:37.715022+08	\N	1	702	计算结果	F		\N	aps:mrp:results	#	5	0	1	0	0
10740	2026-04-04 23:45:37.715022+08	2026-04-04 23:45:37.715022+08	\N	1	703	甘特图	F		\N	aps:schedule:gantt	#	2	0	1	0	0
10741	2026-04-04 23:45:37.715022+08	2026-04-04 23:45:37.715022+08	\N	1	703	排程结果	F		\N	aps:schedule:results	#	5	0	1	0	0
1000	2026-04-08 17:39:49.819042+08	2026-04-08 17:40:30.753952+08	\N	1	2	AI助手配置	C	/system/ai-config	system/AiConfigView.vue	system:aiconfig:list	ChatDotRound	11	1	1	1	0
11	2026-04-09 14:34:19.351603+08	2026-04-09 14:34:19.351603+08	\N	1	0	流程管理	M	/bpm	\N	\N	Guide	14	1	1	1	0
26	2026-04-09 14:36:07.193022+08	2026-04-09 14:36:07.193022+08	\N	1	0	统一告警	M	/alert	\N	\N	Bell	13	1	1	1	0
52	2026-04-15 08:51:26.429762+08	2026-04-15 08:51:26.429762+08	\N	1	2	Notice	C	/system/notice	\N	\N	Bell	20	1	1	1	0
28	2026-04-09 14:36:07.194745+08	2026-04-09 14:36:07.194745+08	\N	1	25	询价单	C	/scp/rfq	scp/RFQList.vue	\N	PriceTag	2	1	1	1	0
46	2026-04-14 15:51:05.089922+08	2026-04-14 15:51:05.089922+08	\N	1	0	生产报表	M	/report	\N	\N	DataLine	15	1	1	1	0
53	2026-04-15 08:51:37.116701+08	2026-04-15 08:51:37.116701+08	\N	1	2	PrintTemplate	C	/system/print-template	\N	\N	Printer	21	1	1	1	0
29	2026-04-09 14:36:07.195241+08	2026-04-09 14:36:07.195241+08	\N	1	25	供应商报价	C	/scp/supplier-quote	scp/SupplierQuoteList.vue	\N	Document	3	1	1	1	0
30	2026-04-09 14:36:07.195678+08	2026-04-09 14:36:07.195678+08	\N	1	25	销售订单	C	/scp/sales-order	scp/SCPSalesOrderList.vue	\N	DocumentCopy	4	1	1	1	0
31	2026-04-09 14:36:07.196422+08	2026-04-09 14:36:07.196422+08	\N	1	25	供应商绩效	C	/scp/supplier-kpi	scp/SupplierKPIList.vue	\N	DataLine	5	1	1	1	0
32	2026-04-09 14:36:07.196991+08	2026-04-09 14:36:07.196991+08	\N	1	26	告警规则	C	/alert/rules	alert/AlertRulesList.vue	\N	SetUp	1	1	1	1	0
33	2026-04-09 14:36:07.197633+08	2026-04-09 14:36:07.197633+08	\N	1	26	告警记录	C	/alert/records	alert/AlertRecordsList.vue	\N	List	2	1	1	1	0
35	2026-04-09 14:36:07.198679+08	2026-04-09 14:36:07.198679+08	\N	1	26	告警统计	C	/alert/statistics	alert/AlertStatistics.vue	\N	DataAnalysis	4	1	1	1	0
59	2026-04-15 08:53:22.768913+08	2026-04-15 08:53:22.768913+08	\N	1	26	AlertNotification	C	/alert/notification	alert/AlertNotification.vue	\N	Message	5	1	1	1	0
12	2026-04-09 14:34:19.352535+08	2026-04-09 14:34:19.352535+08	\N	1	11	流程模型	C	/bpm/process	bpm/ProcessList.vue	\N	FlowChart	1	1	1	1	0
13	2026-04-09 14:34:19.352535+08	2026-04-09 14:34:19.352535+08	\N	1	11	流程实例	C	/bpm/instance	bpm/InstanceList.vue	\N	Connection	2	1	1	1	0
14	2026-04-09 14:34:19.352535+08	2026-04-09 14:34:19.352535+08	\N	1	11	任务实例	C	/bpm/task	bpm/TaskList.vue	\N	Tickets	3	1	1	1	0
16	2026-04-09 14:34:19.353885+08	2026-04-09 14:34:19.353885+08	\N	1	8	交付分析	C	/aps/delivery-analysis	aps/DeliveryAnalysisList.vue	\N	TrendCharts	5	1	1	1	0
17	2026-04-09 14:34:19.353885+08	2026-04-09 14:34:19.353885+08	\N	1	8	缺料分析	C	/aps/material-shortage	aps/MaterialShortageList.vue	\N	Warning	6	1	1	1	0
18	2026-04-09 14:34:19.353885+08	2026-04-09 14:34:19.353885+08	\N	1	8	缺料规则	C	/aps/shortage-rule	aps/ShortageRuleList.vue	\N	SetUp	7	1	1	1	0
19	2026-04-09 14:34:19.353885+08	2026-04-09 14:34:19.353885+08	\N	1	8	换型矩阵	C	/aps/changeover-matrix	aps/ChangeoverMatrixList.vue	\N	Grid	8	1	1	1	0
20	2026-04-09 14:34:19.353885+08	2026-04-09 14:34:19.353885+08	\N	1	8	产品族	C	/aps/product-family	aps/ProductFamilyList.vue	\N	Collection	9	1	1	1	0
22	2026-04-09 14:34:19.355714+08	2026-04-09 14:34:19.355714+08	\N	1	6	盘点管理	C	/wms/stock-check	wms/StockCheckList.vue	\N	DocumentChecked	5	1	1	1	0
58	2026-04-15 08:53:10.444421+08	2026-04-15 08:53:10.444421+08	\N	1	\N	MaterialTrace	C	/mes/material-trace	mes/MaterialTrace.vue	\N	Search	24	1	1	1	0
44	2026-04-09 14:36:07.20116+08	2026-04-09 14:36:07.20116+08	\N	1	5	量检具管理	C	/equipment/gauge	equipment/GaugeList.vue	\N	Scale	10	1	1	1	0
38	2026-04-09 14:36:07.199699+08	2026-04-09 14:36:07.199699+08	\N	1	5	点检计划	C	/equipment/inspection/plans	equipment/inspection/InspectionPlanList.vue	\N	Schedule	7	1	1	1	0
40	2026-04-09 14:36:07.200206+08	2026-04-09 14:36:07.200206+08	\N	1	5	点检记录	C	/equipment/inspection/records	equipment/inspection/InspectionRecordList.vue	\N	Document	8	1	1	1	0
36	2026-04-09 14:36:07.199196+08	2026-04-09 14:36:07.199196+08	\N	1	5	点检模板	C	/equipment/inspection/templates	equipment/inspection/InspectionTemplateList.vue	\N	Tickets	6	1	1	1	0
42	2026-04-09 14:36:07.200621+08	2026-04-09 14:36:07.200621+08	\N	1	5	点检缺陷	C	/equipment/inspection/defects	equipment/inspection/DefectList.vue	\N	Warning	9	1	1	1	0
56	2026-04-15 08:52:36.822246+08	2026-04-15 08:52:36.822246+08	\N	1	4	CodeRule	C	/production/code-rule	production/CodeRuleList.vue	\N	Key	14	1	1	1	0
54	2026-04-15 08:52:10.031861+08	2026-04-15 08:52:10.031861+08	\N	1	4	ElectronicSOP	C	/production/electronic-sop	production/ElectronicSOPList.vue	\N	Document	12	1	1	1	0
55	2026-04-15 08:52:25.218358+08	2026-04-15 08:52:25.218358+08	\N	1	4	FlowCard	C	/production/flow-card	production/FlowCardList.vue	\N	List	13	1	1	1	0
24	2026-04-09 14:34:19.356632+08	2026-04-09 14:34:19.356632+08	\N	1	7	LPA分层审核	C	/quality/lpa	quality/LPAList.vue	\N	Checked	11	1	1	1	0
23	2026-04-09 14:34:19.356632+08	2026-04-09 14:34:19.356632+08	\N	1	7	QRCI质量闭环	C	/quality/qrci	quality/QRCIList.vue	\N	CircleCheck	10	1	1	1	0
57	2026-04-15 08:52:47.546281+08	2026-04-15 08:52:47.546281+08	\N	1	7	DynamicRule	C	/quality/dynamic-rule	quality/DynamicRuleList.vue	\N	Setting	17	1	1	1	0
51	2026-04-14 15:58:27.495007+08	2026-04-14 15:58:27.495007+08	\N	1	46	AnDong BaoBiao	C	/report/andon	report/AndonReport.vue	\N	Bell	5	1	1	1	0
50	2026-04-14 15:58:27.494456+08	2026-04-14 15:58:27.494456+08	\N	1	46	JiaoFu BaoBiao	C	/report/delivery	report/DeliveryReport.vue	\N	Truck	4	1	1	1	0
49	2026-04-14 15:58:27.493804+08	2026-04-14 15:58:27.493804+08	\N	1	46	OEE BaoBiao	C	/report/oee	report/OEEReport.vue	\N	TrendCharts	3	1	1	1	0
48	2026-04-14 15:58:27.493311+08	2026-04-14 15:58:27.493311+08	\N	1	46	ZhiLiang ZhouBao	C	/report/quality-weekly	report/QualityWeeklyReport.vue	\N	DataAnalysis	2	1	1	1	0
25	2026-04-09 14:36:07.189798+08	2026-04-16 10:10:48.506809+08	\N	1	0	供应链管理	M	/scp	\N		Connection	1	1	1	1	0
69	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	25	客户询价	C	/scp/customer-inquiry	scp/CustomerInquiryList.vue	\N	Message	6	1	1	1	0
70	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	4	生产发料	C	/production/issue	production/ProductionIssueList.vue	\N	Top	1	1	1	1	0
71	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	4	生产退料	C	/production/return	production/ProductionReturnList.vue	\N	Bottom	2	1	1	1	0
72	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	7	检验计划	C	/quality/inspection-plans	quality/InspectionPlanList.vue	\N	SetUp	3	1	1	1	0
73	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	7	AQL配置	C	/quality/aql	quality/AQLList.vue	\N	DataLine	4	1	1	1	0
74	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	7	实验室样品	C	/quality/lab/samples	quality/lab/LabSampleList.vue	\N	Sample	5	1	1	1	0
75	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	7	实验室报告	C	/quality/lab/reports	quality/lab/LabReportList.vue	\N	Document	6	1	1	1	0
76	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	7	检测设备	C	/quality/lab/instrument	quality/lab/LabInstrumentList.vue	\N	Tools	7	1	1	1	0
81	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	\N	工厂日历	C	/eam/factory	eam/FactoryList.vue	\N	Calendar	1	1	1	1	0
82	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	\N	设备组织	C	/eam/equipment-org	eam/EquipmentOrgList.vue	\N	Office	2	1	1	1	0
83	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	\N	停机记录	C	/eam/downtime	eam/DowntimeList.vue	\N	Switch	3	1	1	1	0
84	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	\N	付款申请	C	/fin/payment-request	fin/PaymentRequestList.vue	\N	Money	1	1	1	1	0
85	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	\N	采购结算	C	/fin/purchase-settlement	fin/PurchaseSettlementList.vue	\N	ShoppingCart	2	1	1	1	0
86	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	\N	销售结算	C	/fin/sales-settlement	fin/SalesSettlementList.vue	\N	Sell	3	1	1	1	0
87	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	\N	接口配置	C	/integration/interface-config	integration/InterfaceConfigList.vue	\N	Setting	1	1	1	1	0
88	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	\N	执行日志	C	/integration/execution-log	integration/ExecutionLogList.vue	\N	List	2	1	1	1	0
89	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	\N	AGV任务	C	/agv/task	agv/AGVTaskList.vue	\N	VideoPlay	1	1	1	1	0
90	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	\N	AGV设备	C	/agv/device	agv/AGVDeviceList.vue	\N	Van	2	1	1	1	0
91	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	\N	AGV站点	C	/agv/location	agv/AGVLocationList.vue	\N	Location	3	1	1	1	0
92	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	\N	ASN	C	/supplier/asn	supplier/ASNList.vue	\N	Truck	1	1	1	1	0
27	2026-04-09 14:36:07.193461+08	2026-04-09 14:36:07.193461+08	\N	1	25	采购订单	C	/scp/purchase	scp/PurchaseOrderList.vue	\N	ShoppingCart	1	1	1	1	0
34	2026-04-09 14:36:07.198165+08	2026-04-09 14:36:07.198165+08	\N	1	26	升级规则	C	/alert/escalation	alert/AlertEscalationList.vue	\N	Top	3	1	1	1	0
15	2026-04-09 14:34:19.353885+08	2026-04-09 14:34:19.353885+08	\N	1	8	滚动排程	C	/aps/rolling-config	aps/RollingConfigList.vue	\N	Refresh	4	1	1	1	0
21	2026-04-09 14:34:19.355714+08	2026-04-09 14:34:19.355714+08	\N	1	6	调拨管理	C	/wms/transfer	wms/TransferOrderList.vue	\N	Switch	4	1	1	1	0
77	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	\N	班组管理	C	/mes/team	mes/TeamList.vue	\N	User	1	1	1	1	0
78	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	\N	工艺路线	C	/mes/process-routes	mes/ProcessRouteList.vue	\N	Flow	2	1	1	1	0
79	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	\N	离线作业	C	/mes/offline	mes/OfflineOperationList.vue	\N	VideoPause	3	1	1	1	0
80	2026-04-21 13:02:57.840739+08	2026-04-21 13:02:57.840739+08	\N	1	\N	人员技能	C	/mes/person-skill	mes/PersonSkillList.vue	\N	Star	4	1	1	1	0
47	2026-04-14 15:58:27.490929+08	2026-04-14 15:58:27.490929+08	\N	1	46	ShengChan RiBao	C	/report/production-daily	report/ProductionDailyReport.vue	\N	DataLine	1	1	1	1	0
\.


--
-- Data for Name: sys_notice; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_notice (id, created_at, updated_at, deleted_at, tenant_id, title, content, notice_type, priority, publish_dept, publisher_id, publisher_name, publish_time, effect_time, expire_time, target_type, target_ids, is_top, status, view_count, remark) FROM stdin;
\.


--
-- Data for Name: sys_notice_read_record; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_notice_read_record (id, created_at, updated_at, deleted_at, tenant_id, notice_id, user_id, user_name, read_time) FROM stdin;
\.


--
-- Data for Name: sys_oper_log; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_oper_log (id, created_at, updated_at, deleted_at, tenant_id, title, business_type, method, request_method, operator_type, oper_name, dept_name, oper_url, oper_ip, oper_location, oper_param, json_result, status, error_msg, oper_time) FROM stdin;
1	2026-03-31 08:00:00+08	2026-03-31 08:00:00+08	\N	1	用户登录	login	Login	POST	1	系统管理员	信息中心	/api/v1/login	192.168.1.100	浙江省杭州市	{"username":"admin"}	{"code":0,"message":"success"}	0		2026-03-31 08:00:00
2	2026-03-31 08:30:00+08	2026-03-31 08:30:00+08	\N	1	查询用户列表	query	GetUserList	GET	1	系统管理员	信息中心	/api/v1/system/user/list	192.168.1.100	浙江省杭州市	{}	{"code":0,"data":[]}	0		2026-03-31 08:30:00
3	2026-03-31 09:00:00+08	2026-03-31 09:00:00+08	\N	1	创建用户	create	CreateUser	POST	1	系统管理员	信息中心	/api/v1/system/user	192.168.1.100	浙江省杭州市	{"username":"newuser","nickname":"新用户"}	{"code":0,"message":"success"}	0		2026-03-31 09:00:00
4	2026-03-31 09:30:00+08	2026-03-31 09:30:00+08	\N	2	更新用户	update	UpdateUser	PUT	1	杭州管理员	杭州信息中心	/api/v1/system/user/5	192.168.2.100	浙江省杭州市	{"nickname":"更新后的昵称"}	{"code":0,"message":"success"}	0		2026-03-31 09:30:00
5	2026-03-31 10:00:00+08	2026-03-31 10:00:00+08	\N	1	删除用户	delete	DeleteUser	DELETE	1	系统管理员	信息中心	/api/v1/system/user/6	192.168.1.100	浙江省杭州市	{"id":6}	{"code":0,"message":"success"}	0		2026-03-31 10:00:00
6	2026-03-31 10:30:00+08	2026-03-31 10:30:00+08	\N	1	查询租户列表	query	GetTenantList	GET	1	系统管理员	信息中心	/api/v1/system/tenant/list	192.168.1.100	浙江省杭州市	{}	{"code":0,"data":{"list":[]}}	0		2026-03-31 10:30:00
7	2026-03-31 11:00:00+08	2026-03-31 11:00:00+08	\N	1	创建租户	create	CreateTenant	POST	1	系统管理员	信息中心	/api/v1/system/tenant	192.168.1.100	浙江省杭州市	{"tenant_name":"新租户","tenant_key":"newtenant"}	{"code":0,"message":"success"}	0		2026-03-31 11:00:00
8	2026-03-31 11:30:00+08	2026-03-31 11:30:00+08	\N	3	查询部门列表	query	GetDeptList	GET	1	宁波管理员	宁波信息中心	/api/v1/system/dept/list	192.168.3.100	浙江省宁波市	{}	{"code":0,"data":[]}	0		2026-03-31 11:30:00
9	2026-03-31 13:00:00+08	2026-03-31 13:00:00+08	\N	1	查询角色列表	query	GetRoleList	GET	1	系统管理员	信息中心	/api/v1/system/role/list	192.168.1.100	浙江省杭州市	{}	{"code":0,"data":[]}	0		2026-03-31 13:00:00
10	2026-03-31 13:30:00+08	2026-03-31 13:30:00+08	\N	1	创建角色	create	CreateRole	POST	1	系统管理员	信息中心	/api/v1/system/role	192.168.1.100	浙江省杭州市	{"role_name":"新角色","role_key":"newrole"}	{"code":0,"message":"success"}	0		2026-03-31 13:30:00
11	2026-03-31 14:00:00+08	2026-03-31 14:00:00+08	\N	4	更新角色	update	UpdateRole	PUT	1	温州管理员	温州信息中心	/api/v1/system/role/3	192.168.4.100	浙江省温州市	{"role_name":"更新的角色"}	{"code":0,"message":"success"}	0		2026-03-31 14:00:00
12	2026-03-31 14:30:00+08	2026-03-31 14:30:00+08	\N	1	查询菜单列表	query	GetMenuList	GET	1	系统管理员	信息中心	/api/v1/system/menu/list	192.168.1.100	浙江省杭州市	{}	{"code":0,"data":[]}	0		2026-03-31 14:30:00
13	2026-03-31 15:00:00+08	2026-03-31 15:00:00+08	\N	5	查询物料列表	query	GetMaterialList	GET	1	嘉兴管理员	嘉兴信息中心	/api/v1/mdm/material/list	192.168.5.100	浙江省嘉兴市	{}	{"code":0,"data":[]}	0		2026-03-31 15:00:00
14	2026-03-31 15:30:00+08	2026-03-31 15:30:00+08	\N	1	创建物料	create	CreateMaterial	POST	1	系统管理员	信息中心	/api/v1/mdm/material	192.168.1.100	浙江省杭州市	{"material_name":"新物料","material_code":"M001"}	{"code":0,"message":"success"}	0		2026-03-31 15:30:00
15	2026-03-31 16:00:00+08	2026-03-31 16:00:00+08	\N	6	查询生产线	query	GetLineList	GET	1	湖州管理员	湖州信息中心	/api/v1/mdm/line/list	192.168.6.100	浙江省湖州市	{}	{"code":0,"data":[]}	0		2026-03-31 16:00:00
16	2026-03-31 16:30:00+08	2026-03-31 16:30:00+08	\N	7	查询仓库	query	GetWarehouseList	GET	1	绍兴管理员	绍兴信息中心	/api/v1/wms/warehouse/list	192.168.7.100	浙江省绍兴市	{}	{"code":0,"data":[]}	0		2026-03-31 16:30:00
17	2026-03-31 17:00:00+08	2026-03-31 17:00:00+08	\N	8	创建仓库	create	CreateWarehouse	POST	1	金华管理员	金华信息中心	/api/v1/wms/warehouse	192.168.8.100	浙江省金华市	{"warehouse_name":"新仓库"}	{"code":0,"message":"success"}	0		2026-03-31 17:00:00
18	2026-03-31 17:30:00+08	2026-03-31 17:30:00+08	\N	1	查询MPS计划	query	GetMPSList	GET	1	系统管理员	信息中心	/api/v1/aps/mps/list	192.168.1.100	浙江省杭州市	{}	{"code":0,"data":[]}	0		2026-03-31 17:30:00
19	2026-03-31 18:00:00+08	2026-03-31 18:00:00+08	\N	9	查询设备列表	query	GetEquipmentList	GET	1	衢州管理员	衢州信息中心	/api/v1/equipment/list	192.168.9.100	浙江省衢州市	{}	{"code":0,"data":[]}	0		2026-03-31 18:00:00
20	2026-03-31 08:00:00+08	2026-03-31 08:00:00+08	\N	10	用户登录	login	Login	POST	1	舟山操作员	舟山信息中心	/api/v1/login	192.168.10.100	浙江省舟山市	{"username":"zhoushan01"}	{"code":0,"message":"success"}	0		2026-03-31 08:00:00
21	2026-03-31 08:30:00+08	2026-03-31 08:30:00+08	\N	11	查询IQC检验单	query	GetIQCList	GET	1	台州质检员	台州质检部	/api/v1/quality/iqc/list	192.168.11.100	浙江省台州市	{}	{"code":0,"data":[]}	0		2026-03-31 08:30:00
22	2026-03-31 09:00:00+08	2026-03-31 09:00:00+08	\N	12	创建IQC检验单	create	CreateIQC	POST	1	丽水质检员	丽水质检部	/api/v1/quality/iqc	192.168.12.100	浙江省丽水市	{"material_code":"M001","quantity":100}	{"code":0,"message":"success"}	0		2026-03-31 09:00:00
23	2026-03-31 09:30:00+08	2026-03-31 09:30:00+08	\N	13	查询生产工单	query	GetProductionOrderList	GET	1	上海操作员	上海生产部	/api/v1/production/order/list	192.168.13.100	上海市	{}	{"code":0,"data":[]}	0		2026-03-31 09:30:00
24	2026-03-31 10:00:00+08	2026-03-31 10:00:00+08	\N	14	创建生产工单	create	CreateProductionOrder	POST	1	江苏操作员	江苏生产部	/api/v1/production/order	192.168.14.100	江苏省南京市	{"order_no":"PO20260331001"}	{"code":0,"message":"success"}	0		2026-03-31 10:00:00
\.


--
-- Data for Name: sys_post; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_post (id, created_at, updated_at, deleted_at, tenant_id, post_code, post_name, post_sort, status, remark) FROM stdin;
2	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	CEO	首席执行官	1	1	公司最高管理者
3	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	COO	首席运营官	2	1	运营管理
4	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	CTO	首席技术官	3	1	技术管理
5	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	CFO	财务总监	4	1	财务管理
6	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MGR	部门经理	5	1	部门负责人
7	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	DIR	总监	6	1	部门总监
8	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	LEAD	主管	7	1	组长/主管
9	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	ENG	工程师	8	1	技术工程师
10	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	TECH	技术员	9	1	技术员
11	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	OP	操作工	10	1	生产线操作工
12	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	QC	质检员	11	1	质量检验员
13	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	WH	仓管员	12	1	仓库管理员
14	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	ADMIN	管理员	13	1	行政人员
15	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	SA	销售员	14	1	销售人员
16	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	PUR	采购员	15	1	采购人员
17	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	CEO	首席执行官	1	1	\N
18	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	COO	首席运营官	2	1	\N
19	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	CTO	首席技术官	3	1	\N
20	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	CFO	财务总监	4	1	\N
21	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	MGR	部门经理	5	1	\N
22	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	DIR	总监	6	1	\N
23	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	LEAD	主管	7	1	\N
24	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	ENG	工程师	8	1	\N
25	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	TECH	技术员	9	1	\N
26	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	OP	操作工	10	1	\N
27	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	QC	质检员	11	1	\N
28	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	WH	仓管员	12	1	\N
29	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	ADMIN	管理员	13	1	\N
30	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	SA	销售员	14	1	\N
31	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	PUR	采购员	15	1	\N
\.


--
-- Data for Name: sys_print_template; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_print_template (id, created_at, updated_at, deleted_at, tenant_id, template_code, template_name, template_type, entity_type, content, paper_type, paper_width, paper_height, is_default, status, remark) FROM stdin;
\.


--
-- Data for Name: sys_role; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_role (id, created_at, updated_at, deleted_at, tenant_id, role_name, role_key, role_sort, data_scope, status, remark) FROM stdin;
1	2026-03-31 10:48:01.590258+08	2026-03-31 10:48:01.590258+08	\N	1	超级管理员	admin	1	1	1	\N
2	2026-03-31 10:48:01.590258+08	2026-03-31 10:48:01.590258+08	\N	1	普通用户	user	2	1	1	\N
3	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	超级管理员	super	1	1	1	拥有系统所有权限
4	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	系统管理员	admin	2	2	1	系统管理权限
5	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	部门经理	mgr	3	2	1	管理部门数据
6	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	生产经理	prd_mgr	4	2	1	生产管理权限
7	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	质量经理	qc_mgr	5	2	1	质量管理权限
8	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	普通用户	user	6	5	1	普通操作权限
9	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	操作员	operator	7	5	1	生产线操作员
10	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	质检员	qc	8	5	1	质量检验权限
11	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	仓管员	warehouse	9	5	1	仓库管理权限
12	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	1	访客	guest	10	7	1	只读权限
\.


--
-- Data for Name: sys_role_menu; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_role_menu (role_id, menu_id) FROM stdin;
4	1
4	2
4	101
4	102
4	103
4	104
4	105
4	106
4	107
4	108
4	109
4	110
4	3
4	4
4	5
4	6
4	7
4	601
4	602
4	603
4	8
4	701
4	702
4	703
4	9
4	801
4	802
4	10
4	901
1	1
1	2
1	3
1	4
1	5
1	6
1	7
1	8
1	9
1	10
1	209
1	210
1	307
1	409
1	504
1	505
1	704
1	10001
1	10002
1	10003
1	10004
1	10005
1	10006
1	10007
1	10008
1	10009
1	10010
1	10011
1	10012
1	10013
1	10014
1	10015
1	10016
1	10017
1	10018
1	10019
1	10020
1	10021
1	10022
1	10023
1	10024
1	10025
1	10026
1	10027
1	10101
1	10102
1	10103
1	10104
1	10105
1	10106
1	10107
1	10108
1	10109
1	10110
1	10111
1	10112
1	10113
1	10114
1	10115
1	10116
1	10117
1	10118
1	10119
1	10120
1	10121
1	10122
1	10123
1	10124
1	10125
1	10126
1	10127
1	10128
1	10129
1	10301
1	10302
1	10303
1	10304
1	10305
1	10306
1	10307
1	10308
1	10309
1	10310
1	10311
1	10312
1	10313
1	10314
1	10315
1	10316
1	10317
1	10318
1	10401
1	10402
1	10403
1	10404
1	10405
1	10406
1	10407
1	10408
1	10409
1	10410
1	10411
1	10412
1	10413
1	10414
1	10415
1	10416
1	10501
1	10502
1	10503
1	10504
1	10505
1	10506
1	10507
1	10508
1	10509
1	10510
1	10511
1	10512
1	10513
1	10514
1	10515
1	10516
1	10601
1	10602
1	10603
1	10604
1	10605
1	10606
1	10607
1	10608
1	10609
1	10610
1	10611
1	10612
1	10613
1	10614
1	10615
1	10616
1	10617
1	10618
1	10619
1	10620
1	10621
1	10622
1	10623
1	10624
1	10625
1	10701
1	10702
1	10703
1	10704
1	10705
1	10706
1	10707
1	10708
1	10709
1	10710
1	10711
1	10712
1	10713
1	10714
1	10715
1	10417
1	10418
1	10419
1	10720
1	10721
1	10420
1	10421
1	10520
1	10521
1	10522
1	10523
1	10524
1	10525
1	10526
1	10730
1	10731
1	10740
1	10741
1	11
1	12
1	13
1	14
1	15
1	16
1	17
1	18
1	19
1	20
1	21
1	22
1	23
1	24
1	25
1	27
1	28
1	29
1	30
1	31
1	26
1	32
1	33
1	34
1	35
1	36
1	37
1	38
1	39
1	40
1	41
1	42
1	43
1	44
1	45
1	46
1	52
1	53
1	54
1	55
1	56
1	57
1	58
1	59
1	47
1	48
1	49
1	50
1	51
1	69
1	70
1	71
1	72
1	73
1	74
1	75
1	76
1	77
1	78
1	79
1	80
1	81
1	82
1	83
1	84
1	85
1	86
1	87
1	88
1	89
1	90
1	91
1	92
\.


--
-- Data for Name: sys_role_perm; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_role_perm (role_id, perm) FROM stdin;
1	user:add
1	user:edit
1	user:delete
\.


--
-- Data for Name: sys_tenant; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_tenant (id, created_at, updated_at, deleted_at, tenant_name, tenant_key, contact, phone, email, status, expire_time, package_id, province, city, district, address, manager, contact_name, contact_phone, contact_email, factory_type, employee_count, area_size, annual_capacity, remark) FROM stdin;
1	2026-03-31 10:48:01.587236+08	2026-03-31 10:48:01.587236+08	\N	默认租户	default	\N	\N	\N	1	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
2	2026-01-01 08:00:00+08	2026-01-01 08:00:00+08	\N	峰梅动力科技有限公司	fengmei	张总	13800138001	zhang@fengmei.com	1	2027-12-31 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
3	2026-01-02 08:00:00+08	2026-01-02 08:00:00+08	\N	杭州分公司	hangzhou	李经理	13800138002	li@fengmei.com	1	2027-12-31 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
4	2026-01-03 08:00:00+08	2026-01-03 08:00:00+08	\N	宁波分公司	ningbo	王经理	13800138003	wang@fengmei.com	1	2027-12-31 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
5	2026-01-04 08:00:00+08	2026-01-04 08:00:00+08	\N	温州分公司	wenzhou	陈经理	13800138004	chen@fengmei.com	1	2027-12-31 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
6	2026-01-05 08:00:00+08	2026-01-05 08:00:00+08	\N	嘉兴分公司	jiaxing	刘经理	13800138005	liu@fengmei.com	1	2027-12-31 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
7	2026-01-06 08:00:00+08	2026-01-06 08:00:00+08	\N	湖州分公司	huzhou	周经理	13800138006	zhou@fengmei.com	1	2027-12-31 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
8	2026-01-07 08:00:00+08	2026-01-07 08:00:00+08	\N	绍兴分公司	shaoxing	吴经理	13800138007	wu@fengmei.com	1	2027-12-31 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
9	2026-01-08 08:00:00+08	2026-01-08 08:00:00+08	\N	金华分公司	jinhua	徐经理	13800138008	xu@fengmei.com	1	2027-12-31 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
10	2026-01-09 08:00:00+08	2026-01-09 08:00:00+08	\N	衢州分公司	quzhou	孙经理	13800138009	sun@fengmei.com	1	2027-12-31 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
11	2026-01-10 08:00:00+08	2026-01-10 08:00:00+08	\N	舟山分公司	zhoushan	马经理	13800138010	ma@fengmei.com	1	2027-12-31 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
12	2026-01-11 08:00:00+08	2026-01-11 08:00:00+08	\N	台州分公司	taizhou	朱经理	13800138011	zhu@fengmei.com	1	2027-12-31 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
13	2026-01-12 08:00:00+08	2026-01-12 08:00:00+08	\N	丽水分公司	lishui	何经理	13800138012	he@fengmei.com	1	2027-12-31 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
14	2026-01-13 08:00:00+08	2026-01-13 08:00:00+08	\N	上海分公司	shanghai	罗经理	13800138013	luo@fengmei.com	1	2027-12-31 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
15	2026-01-14 08:00:00+08	2026-01-14 08:00:00+08	\N	江苏分公司	jiangsu	宋经理	13800138014	song@fengmei.com	1	2027-12-31 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
16	2026-01-15 08:00:00+08	2026-01-15 08:00:00+08	\N	安徽分公司	anhui	唐经理	13800138015	tang@fengmei.com	1	2027-12-31 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
17	2026-01-16 08:00:00+08	2026-01-16 08:00:00+08	\N	福建分公司	fujian	冯经理	13800138016	feng@fengmei.com	1	2027-12-31 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
18	2026-01-17 08:00:00+08	2026-01-17 08:00:00+08	\N	江西分公司	jiangxi	许经理	13800138017	xu2@fengmei.com	1	2027-12-31 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
19	2026-01-18 08:00:00+08	2026-01-18 08:00:00+08	\N	山东分公司	shandong	蒋经理	13800138018	jiang@fengmei.com	1	2027-12-31 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
20	2026-01-19 08:00:00+08	2026-01-19 08:00:00+08	\N	测试租户A	test_a	测试联系人	13800138019	test@fengmei.com	1	2026-06-30 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
21	2026-01-20 08:00:00+08	2026-01-20 08:00:00+08	\N	禁用租户	disabled	禁用联系人	13800138020	disabled@fengmei.com	0	2026-03-31 23:59:59	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N	\N
\.


--
-- Data for Name: sys_user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_user (id, created_at, updated_at, deleted_at, tenant_id, username, nickname, password, email, phone, avatar, dept_id, status, login_ip, login_date) FROM stdin;
1	2026-03-31 10:48:01.588815+08	2026-03-31 10:48:01.588815+08	\N	1	admin	管理员	$2a$10$TlqyWNgrYfCl470V2R6Rhe3eG6yJZcoopIkO4B8vfNgyv48O6Xn7a	\N	\N	\N	\N	1	\N	\N
2	2026-04-01 15:43:22.755878+08	2026-04-01 15:43:22.755878+08	\N	1	nhg	nhg	$2a$10$8.crKXR4UMGyhrDYKH8LnuCTj1vqwaS9Bg5gqRIdn.MegETVrNoLa			\N	\N	1	\N	\N
\.


--
-- Data for Name: sys_user_role; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sys_user_role (user_id, role_id) FROM stdin;
1	1
2	4
\.


--
-- Data for Name: tra_andon_call; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tra_andon_call (id, created_at, updated_at, deleted_at, tenant_id, call_no, line_id, line_name, station_id, station_name, call_type, call_level, call_desc, call_user_id, call_user_name, call_time, response_user_id, response_time, resolve_time, status, remark) FROM stdin;
\.


--
-- Data for Name: tra_data_collection; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tra_data_collection (id, created_at, updated_at, deleted_at, tenant_id, equipment_id, equipment_code, station_id, data_type, data_key, data_value, unit, collect_time) FROM stdin;
\.


--
-- Data for Name: tra_serial_number; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tra_serial_number (id, created_at, updated_at, deleted_at, tenant_id, serial_number, material_id, material_code, material_name, batch_no, line_id, line_name, order_id, order_no, production_date, status) FROM stdin;
\.


--
-- Data for Name: tra_trace_record; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tra_trace_record (id, created_at, updated_at, deleted_at, tenant_id, trace_no, material_id, material_code, material_name, serial_number, batch_no, process_id, process_name, station_id, station_name, operator_id, operator_name, operate_time, operate_type, input_qty, output_qty, reject_qty) FROM stdin;
\.


--
-- Data for Name: trace_records; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.trace_records (id, created_at, updated_at, deleted_at, tenant_id, trace_no, material_id, material_code, material_name, serial_number, batch_no, process_id, process_name, station_id, station_name, operator_id, operator_name, operate_time, operate_type, input_qty, output_qty, reject_qty) FROM stdin;
\.


--
-- Data for Name: transfer_orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transfer_orders (id, created_at, updated_at, deleted_at, tenant_id, transfer_no, from_warehouse_id, to_warehouse_id, transfer_date, transfer_user_id, status, remark) FROM stdin;
\.


--
-- Data for Name: visual_inspection_result; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.visual_inspection_result (id, created_at, updated_at, task_id, detection_time, result, confidence, defect_type, defect_location, defect_image_url, ai_analysis, manual_review_result, manual_review_by, manual_review_at, remark, tenant_id) FROM stdin;
\.


--
-- Data for Name: visual_inspection_task; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.visual_inspection_task (id, created_at, updated_at, task_no, task_type, product_id, product_code, product_name, production_order_id, workshop_id, image_url, image_hash, detection_standard, ai_model_version, status, priority, requested_by, requested_at, completed_at, remark, tenant_id, created_by) FROM stdin;
\.


--
-- Data for Name: warehouses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.warehouses (id, created_at, updated_at, deleted_at, tenant_id, warehouse_code, warehouse_name, warehouse_type, address, manager, phone, status) FROM stdin;
\.


--
-- Data for Name: wms_agv_device; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_agv_device (id, tenant_id, agv_code, agv_name, agv_type, status, current_location, battery_level, max_load, ext_config, last_heartbeat, remark, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: wms_agv_location_mapping; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_agv_location_mapping (id, tenant_id, location_code, location_name, location_type, agv_location_code, x_coord, y_coord, priority, enabled, ext_data, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: wms_agv_task; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_agv_task (id, tenant_id, task_no, task_type, status, priority, source_location_id, source_location, target_location_id, target_location, material_id, material_code, material_name, quantity, unit, assigned_agv_code, assigned_agv_name, started_at, completed_at, error_message, related_order_no, related_order_type, ext_data, created_by, created_at, updated_at) FROM stdin;
\.


--
-- Data for Name: wms_delivery_order; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_delivery_order (id, created_at, updated_at, deleted_at, tenant_id, delivery_no, customer_id, customer_name, warehouse_id, delivery_date, delivery_user_id, status, remark) FROM stdin;
\.


--
-- Data for Name: wms_delivery_order_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_delivery_order_item (id, created_at, updated_at, deleted_at, delivery_id, material_id, material_code, material_name, quantity, shipped_qty, unit, batch_no) FROM stdin;
\.


--
-- Data for Name: wms_inventory; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_inventory (id, created_at, updated_at, deleted_at, tenant_id, material_id, material_code, material_name, warehouse_id, location_id, quantity, available_qty, allocated_qty, locked_qty, batch_no) FROM stdin;
1	2026-04-04 01:33:43.430159+08	2026-04-04 01:33:43.430159+08	\N	1	1	MAT-001	钢板A3	4	1	1000.0000	1000.0000	0.0000	0.0000	BATCH-20260401
2	2026-04-04 01:33:43.430159+08	2026-04-04 01:33:43.430159+08	\N	1	2	MAT-002	钢板A4	4	2	800.0000	800.0000	0.0000	0.0000	BATCH-20260402
3	2026-04-04 01:33:43.430159+08	2026-04-04 01:33:43.430159+08	\N	1	3	MAT-003	铝合金板	4	3	500.0000	500.0000	0.0000	0.0000	BATCH-20260403
4	2026-04-04 01:33:43.430159+08	2026-04-04 01:33:43.430159+08	\N	1	5	MAT-005	塑料粒子PP	4	4	2000.0000	2000.0000	0.0000	0.0000	BATCH-20260404
5	2026-04-04 01:33:43.430159+08	2026-04-04 01:33:43.430159+08	\N	1	1	MAT-001	钢板A3	5	5	200.0000	200.0000	0.0000	0.0000	BATCH-FG-001
6	2026-04-04 01:33:43.430159+08	2026-04-04 01:33:43.430159+08	\N	1	2	MAT-002	钢板A4	5	6	150.0000	150.0000	0.0000	0.0000	BATCH-FG-002
\.


--
-- Data for Name: wms_inventory_record; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_inventory_record (id, created_at, updated_at, deleted_at, tenant_id, record_no, record_type, material_id, material_code, material_name, warehouse_id, location_id, quantity, batch_no, source_no, operator_id, operator_name, operate_time, remark) FROM stdin;
\.


--
-- Data for Name: wms_kanban_pull; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_kanban_pull (id, created_at, updated_at, deleted_at, tenant_id, kanban_no, material_id, material_code, material_name, from_location_id, to_location_id, kanban_qty, trigger_qty, current_qty, lead_time, status, remark) FROM stdin;
\.


--
-- Data for Name: wms_location; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_location (id, created_at, updated_at, deleted_at, tenant_id, location_code, location_name, warehouse_id, zone_code, "row", col, layer, location_type, capacity, status) FROM stdin;
1	2026-04-04 01:32:43.13006+08	2026-04-04 01:32:43.13006+08	\N	1	WH01-A01-001	原料仓A区01排01位	4	A	1	1	1	存储区	\N	1
2	2026-04-04 01:32:43.13006+08	2026-04-04 01:32:43.13006+08	\N	1	WH01-A01-002	原料仓A区01排02位	4	A	1	2	1	存储区	\N	1
3	2026-04-04 01:32:43.13006+08	2026-04-04 01:32:43.13006+08	\N	1	WH01-A02-001	原料仓A区02排01位	4	A	2	1	1	存储区	\N	1
4	2026-04-04 01:32:43.13006+08	2026-04-04 01:32:43.13006+08	\N	1	WH01-B01-001	原料仓B区01排01位	4	B	1	1	1	存储区	\N	1
5	2026-04-04 01:32:43.13006+08	2026-04-04 01:32:43.13006+08	\N	1	WH02-A01-001	成品仓A区01排01位	5	A	1	1	1	存储区	\N	1
6	2026-04-04 01:32:43.13006+08	2026-04-04 01:32:43.13006+08	\N	1	WH02-A01-002	成品仓A区01排02位	5	A	1	2	1	存储区	\N	1
7	2026-04-04 01:32:43.13006+08	2026-04-04 01:32:43.13006+08	\N	1	WH03-001	线边仓工位1	6	A	1	1	1	线边仓	\N	1
8	2026-04-04 01:32:43.13006+08	2026-04-04 01:32:43.13006+08	\N	1	WH03-002	线边仓工位2	6	A	1	2	1	线边仓	\N	1
\.


--
-- Data for Name: wms_production_complete; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_production_complete (id, created_at, updated_at, complete_no, production_order_id, order_no, workshop_id, workstation_id, complete_qty, qualified_qty, status, complete_time, operator_id, operator_name, remark, tenant_id) FROM stdin;
\.


--
-- Data for Name: wms_production_complete_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_production_complete_item (id, complete_id, line_no, material_id, material_code, material_name, unit, complete_qty, qualified_qty, warehouse_id, location_id, batch_no, tenant_id) FROM stdin;
\.


--
-- Data for Name: wms_production_issue; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_production_issue (id, created_at, updated_at, issue_no, issue_type, production_order_id, order_no, workstation_id, workshop_id, status, pick_status, request_by, request_time, issued_by, issued_time, remark, tenant_id, created_by) FROM stdin;
\.


--
-- Data for Name: wms_production_issue_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_production_issue_item (id, issue_id, line_no, material_id, material_code, material_name, unit, required_qty, picked_qty, issued_qty, warehouse_id, location_id, batch_no, remark, tenant_id) FROM stdin;
\.


--
-- Data for Name: wms_production_return; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_production_return (id, created_at, updated_at, return_no, production_order_id, order_no, workstation_id, workshop_id, status, return_type, request_by, request_time, approved_by, approved_time, returned_by, returned_time, remark, tenant_id, created_by) FROM stdin;
\.


--
-- Data for Name: wms_production_return_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_production_return_item (id, return_id, line_no, material_id, material_code, material_name, unit, issued_qty, return_qty, warehouse_id, location_id, batch_no, reason, remark, tenant_id) FROM stdin;
\.


--
-- Data for Name: wms_production_stock_in; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_production_stock_in (id, created_at, updated_at, stock_in_no, complete_id, complete_no, warehouse_id, location_id, status, stock_in_time, operator_id, operator_name, tenant_id) FROM stdin;
\.


--
-- Data for Name: wms_production_stock_in_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_production_stock_in_item (id, stock_in_id, line_no, material_id, material_code, material_name, unit, stock_in_qty, warehouse_id, location_id, batch_no, tenant_id) FROM stdin;
\.


--
-- Data for Name: wms_purchase_return; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_purchase_return (id, created_at, updated_at, return_no, purchase_order_id, supplier_id, supplier_name, warehouse_id, return_date, status, return_type, request_by, request_time, approved_by, approved_time, returned_by, returned_time, remark, tenant_id, created_by) FROM stdin;
\.


--
-- Data for Name: wms_purchase_return_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_purchase_return_item (id, return_id, line_no, material_id, material_code, material_name, unit, return_qty, reason, tenant_id) FROM stdin;
\.


--
-- Data for Name: wms_receive_order; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_receive_order (id, created_at, updated_at, deleted_at, tenant_id, receive_no, supplier_id, supplier_name, warehouse_id, receive_date, receive_user_id, status, remark) FROM stdin;
\.


--
-- Data for Name: wms_receive_order_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_receive_order_item (id, created_at, updated_at, deleted_at, receive_id, material_id, material_code, material_name, quantity, received_qty, unit, batch_no) FROM stdin;
\.


--
-- Data for Name: wms_sales_return; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_sales_return (id, created_at, updated_at, return_no, sales_order_id, customer_id, customer_name, warehouse_id, return_date, status, return_type, request_by, request_time, approved_by, approved_time, returned_by, returned_time, remark, tenant_id, created_by) FROM stdin;
\.


--
-- Data for Name: wms_sales_return_item; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_sales_return_item (id, return_id, line_no, material_id, material_code, material_name, unit, return_qty, reason, tenant_id) FROM stdin;
\.


--
-- Data for Name: wms_side_location; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_side_location (id, created_at, updated_at, deleted_at, tenant_id, location_code, location_name, workshop_id, workshop_name, line_id, line_name, station_id, station_name, location_type, max_capacity, current_qty, status, remark) FROM stdin;
\.


--
-- Data for Name: wms_stock_check; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_stock_check (id, created_at, updated_at, deleted_at, tenant_id, check_no, warehouse_id, check_date, check_user_id, status, remark, check_type, warehouse_name, area_ids, location_ids, plan_start_date, plan_end_date, actual_start_date, actual_end_date, include_zero_stock, include_expired_stock, is_blind_mode, checker_id, checker_name, audit_id, audit_name, total_locations, counted_locations, total_materials, variance_count, variance_rate, approval_status, approved_by, approved_time, approval_comment, workshop_id) FROM stdin;
\.


--
-- Data for Name: wms_transfer_order; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_transfer_order (id, created_at, updated_at, deleted_at, tenant_id, transfer_no, from_warehouse_id, to_warehouse_id, transfer_date, transfer_user_id, status, remark, transfer_type, from_warehouse_name, from_warehouse_type, to_warehouse_name, to_warehouse_type, from_workstation_id, from_workstation_name, to_workstation_id, to_workstation_name, transfer_reason, source_order_id, source_order_no, requester_id, requester_name, request_time, approver_id, approved_time, approval_comment, actual_start_time, actual_complete_time, operator_id, operator_name, logistics_provider, tracking_no, total_amount, currency, workshop_id) FROM stdin;
\.


--
-- Data for Name: wms_warehouse; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.wms_warehouse (id, created_at, updated_at, deleted_at, tenant_id, warehouse_code, warehouse_name, warehouse_type, address, manager, phone, status) FROM stdin;
1	2026-03-31 10:48:01.606878+08	2026-03-31 10:48:01.606878+08	\N	1	WH001	原材料仓	raw	\N	仓管A	13900139001	1
2	2026-03-31 10:48:01.606878+08	2026-03-31 10:48:01.606878+08	\N	1	WH002	成品仓	finished	\N	仓管B	13900139002	1
3	2026-03-31 10:48:01.606878+08	2026-03-31 10:48:01.606878+08	\N	1	WH003	线边仓	inline	\N	仓管C	13900139003	1
4	2026-04-04 01:31:41.079174+08	2026-04-04 01:31:41.079174+08	\N	1	WH01	原料仓库	原材料	\N	\N	\N	1
5	2026-04-04 01:31:41.079174+08	2026-04-04 01:31:41.079174+08	\N	1	WH02	成品仓库	成品	\N	\N	\N	1
6	2026-04-04 01:31:41.079174+08	2026-04-04 01:31:41.079174+08	\N	1	WH03	线边仓	线边仓	\N	\N	\N	1
\.


--
-- Data for Name: work_centers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.work_centers (id, created_at, updated_at, deleted_at, tenant_id, work_center_code, work_center_name, workshop_id, capacity, status) FROM stdin;
\.


--
-- Data for Name: workshops; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.workshops (id, created_at, updated_at, deleted_at, tenant_id, workshop_code, workshop_name, workshop_type, manager, phone, address, status) FROM stdin;
\.


--
-- Data for Name: workstations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.workstations (id, created_at, updated_at, deleted_at, tenant_id, station_code, station_name, line_id, station_type, status) FROM stdin;
\.


--
-- Name: ai_chat_conversation_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ai_chat_conversation_id_seq', 38, true);


--
-- Name: ai_chat_message_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ai_chat_message_id_seq', 80, true);


--
-- Name: ai_config_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ai_config_id_seq', 1, true);


--
-- Name: alert_escalation_rule_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.alert_escalation_rule_id_seq', 1, false);


--
-- Name: alert_notification_log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.alert_notification_log_id_seq', 1, false);


--
-- Name: alert_record_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.alert_record_id_seq', 1, false);


--
-- Name: alert_rule_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.alert_rule_id_seq', 1, false);


--
-- Name: andon_call_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.andon_call_id_seq', 1, false);


--
-- Name: andon_calls_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.andon_calls_id_seq', 1, false);


--
-- Name: andon_escalation_log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.andon_escalation_log_id_seq', 1, false);


--
-- Name: andon_escalation_rule_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.andon_escalation_rule_id_seq', 1, false);


--
-- Name: andon_notification_log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.andon_notification_log_id_seq', 1, false);


--
-- Name: andon_report_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.andon_report_id_seq', 1, false);


--
-- Name: aps_capacity_analysis_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.aps_capacity_analysis_id_seq', 1, false);


--
-- Name: aps_changeover_matrix_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.aps_changeover_matrix_id_seq', 1, false);


--
-- Name: aps_delivery_rate_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.aps_delivery_rate_id_seq', 1, false);


--
-- Name: aps_jit_demand_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.aps_jit_demand_id_seq', 1, false);


--
-- Name: aps_mps_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.aps_mps_id_seq', 1, false);


--
-- Name: aps_mrp_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.aps_mrp_id_seq', 1, false);


--
-- Name: aps_mrp_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.aps_mrp_item_id_seq', 1, false);


--
-- Name: aps_resource_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.aps_resource_id_seq', 1, false);


--
-- Name: aps_rolling_schedule_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.aps_rolling_schedule_id_seq', 1, false);


--
-- Name: aps_schedule_plan_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.aps_schedule_plan_id_seq', 1, false);


--
-- Name: aps_schedule_result_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.aps_schedule_result_id_seq', 1, false);


--
-- Name: aps_work_center_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.aps_work_center_id_seq', 1, false);


--
-- Name: aps_working_calendar_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.aps_working_calendar_id_seq', 1, false);


--
-- Name: bom_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bom_items_id_seq', 1, false);


--
-- Name: boms_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.boms_id_seq', 1, false);


--
-- Name: bpm_approval_record_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bpm_approval_record_id_seq', 1, false);


--
-- Name: bpm_biz_mapping_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bpm_biz_mapping_id_seq', 1, false);


--
-- Name: bpm_delegate_record_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bpm_delegate_record_id_seq', 1, false);


--
-- Name: bpm_form_definition_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bpm_form_definition_id_seq', 1, false);


--
-- Name: bpm_form_field_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bpm_form_field_id_seq', 1, false);


--
-- Name: bpm_form_version_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bpm_form_version_id_seq', 1, false);


--
-- Name: bpm_model_version_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bpm_model_version_id_seq', 1, false);


--
-- Name: bpm_node_definition_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bpm_node_definition_id_seq', 1, false);


--
-- Name: bpm_node_timeout_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bpm_node_timeout_id_seq', 1, false);


--
-- Name: bpm_order_day_workorder_map_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bpm_order_day_workorder_map_id_seq', 1, false);


--
-- Name: bpm_process_instance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bpm_process_instance_id_seq', 1, false);


--
-- Name: bpm_process_model_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bpm_process_model_id_seq', 1, false);


--
-- Name: bpm_reminder_record_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bpm_reminder_record_id_seq', 1, false);


--
-- Name: bpm_sequence_flow_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bpm_sequence_flow_id_seq', 1, false);


--
-- Name: bpm_task_assignment_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bpm_task_assignment_id_seq', 1, false);


--
-- Name: bpm_task_instance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.bpm_task_instance_id_seq', 1, false);


--
-- Name: container_master_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.container_master_id_seq', 1, false);


--
-- Name: container_movement_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.container_movement_id_seq', 1, false);


--
-- Name: data_collections_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.data_collections_id_seq', 1, false);


--
-- Name: dc_collect_record_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.dc_collect_record_id_seq', 1, false);


--
-- Name: dc_data_point_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.dc_data_point_id_seq', 1, false);


--
-- Name: dc_scan_log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.dc_scan_log_id_seq', 1, false);


--
-- Name: defect_codes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.defect_codes_id_seq', 1, false);


--
-- Name: defect_records_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.defect_records_id_seq', 1, false);


--
-- Name: delivery_order_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.delivery_order_items_id_seq', 1, false);


--
-- Name: delivery_orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.delivery_orders_id_seq', 1, false);


--
-- Name: delivery_report_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.delivery_report_id_seq', 1, false);


--
-- Name: dispatches_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.dispatches_id_seq', 1, false);


--
-- Name: eam_equipment_downtime_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.eam_equipment_downtime_id_seq', 1, false);


--
-- Name: eam_repair_flow_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.eam_repair_flow_id_seq', 1, false);


--
-- Name: eam_repair_job_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.eam_repair_job_id_seq', 1, false);


--
-- Name: eam_repair_std_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.eam_repair_std_id_seq', 1, false);


--
-- Name: ene_energy_record_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ene_energy_record_id_seq', 1, false);


--
-- Name: energy_records_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.energy_records_id_seq', 1, false);


--
-- Name: equ_equipment_check_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.equ_equipment_check_id_seq', 1, false);


--
-- Name: equ_equipment_document_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.equ_equipment_document_id_seq', 1, false);


--
-- Name: equ_equipment_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.equ_equipment_id_seq', 11, true);


--
-- Name: equ_equipment_maintenance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.equ_equipment_maintenance_id_seq', 1, false);


--
-- Name: equ_equipment_part_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.equ_equipment_part_id_seq', 1, false);


--
-- Name: equ_equipment_repair_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.equ_equipment_repair_id_seq', 1, false);


--
-- Name: equ_gauge_calibration_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.equ_gauge_calibration_id_seq', 1, false);


--
-- Name: equ_gauge_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.equ_gauge_id_seq', 1, false);


--
-- Name: equ_mold_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.equ_mold_id_seq', 1, false);


--
-- Name: equ_mold_maintenance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.equ_mold_maintenance_id_seq', 1, false);


--
-- Name: equ_mold_repair_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.equ_mold_repair_id_seq', 1, false);


--
-- Name: equ_oee_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.equ_oee_id_seq', 1, false);


--
-- Name: equ_oee_record_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.equ_oee_record_id_seq', 1, false);


--
-- Name: equ_spare_part_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.equ_spare_part_id_seq', 1, false);


--
-- Name: equ_teep_data_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.equ_teep_data_id_seq', 1, false);


--
-- Name: fin_payment_request_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.fin_payment_request_id_seq', 1, false);


--
-- Name: fin_purchase_advance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.fin_purchase_advance_id_seq', 1, false);


--
-- Name: fin_purchase_settlement_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.fin_purchase_settlement_id_seq', 1, false);


--
-- Name: fin_purchase_settlement_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.fin_purchase_settlement_item_id_seq', 1, false);


--
-- Name: fin_sales_receipt_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.fin_sales_receipt_id_seq', 1, false);


--
-- Name: fin_sales_settlement_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.fin_sales_settlement_id_seq', 1, false);


--
-- Name: fin_sales_settlement_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.fin_sales_settlement_item_id_seq', 1, false);


--
-- Name: fin_statement_detail_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.fin_statement_detail_id_seq', 1, false);


--
-- Name: fin_supplier_statement_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.fin_supplier_statement_id_seq', 1, false);


--
-- Name: fqcs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.fqcs_id_seq', 1, false);


--
-- Name: integration_erp_mappings_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.integration_erp_mappings_id_seq', 1, false);


--
-- Name: integration_erp_sync_logs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.integration_erp_sync_logs_id_seq', 1, false);


--
-- Name: inventories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.inventories_id_seq', 1, false);


--
-- Name: inventory_records_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.inventory_records_id_seq', 1, false);


--
-- Name: ip_qcs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ip_qcs_id_seq', 1, false);


--
-- Name: iqc_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.iqc_items_id_seq', 1, false);


--
-- Name: iqcs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.iqcs_id_seq', 1, false);


--
-- Name: lab_calibration_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.lab_calibration_id_seq', 1, false);


--
-- Name: lab_instrument_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.lab_instrument_id_seq', 1, false);


--
-- Name: lab_report_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.lab_report_id_seq', 1, false);


--
-- Name: lab_sample_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.lab_sample_id_seq', 1, false);


--
-- Name: lab_test_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.lab_test_item_id_seq', 1, false);


--
-- Name: locations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.locations_id_seq', 1, false);


--
-- Name: login_logs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.login_logs_id_seq', 1, false);


--
-- Name: material_categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.material_categories_id_seq', 1, false);


--
-- Name: materials_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.materials_id_seq', 1, false);


--
-- Name: mdm_bom_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mdm_bom_id_seq', 2, true);


--
-- Name: mdm_bom_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mdm_bom_item_id_seq', 11, true);


--
-- Name: mdm_customer_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mdm_customer_id_seq', 6, true);


--
-- Name: mdm_material_category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mdm_material_category_id_seq', 1, false);


--
-- Name: mdm_material_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mdm_material_id_seq', 45, true);


--
-- Name: mdm_operation_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mdm_operation_id_seq', 6, true);


--
-- Name: mdm_production_line_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mdm_production_line_id_seq', 8, true);


--
-- Name: mdm_shift_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mdm_shift_id_seq', 3, true);


--
-- Name: mdm_supplier_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mdm_supplier_id_seq', 78, true);


--
-- Name: mdm_supplier_material_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mdm_supplier_material_id_seq', 1, false);


--
-- Name: mdm_workshop_config_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mdm_workshop_config_id_seq', 1, false);


--
-- Name: mdm_workshop_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mdm_workshop_id_seq', 3, true);


--
-- Name: mdm_workstation_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mdm_workstation_id_seq', 13, true);


--
-- Name: mes_code_rule_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_code_rule_id_seq', 1, false);


--
-- Name: mes_code_rule_record_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_code_rule_record_id_seq', 1, false);


--
-- Name: mes_container_lifecycle_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_container_lifecycle_id_seq', 1, false);


--
-- Name: mes_container_maintenance_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_container_maintenance_id_seq', 1, false);


--
-- Name: mes_electronic_sop_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_electronic_sop_id_seq', 1, false);


--
-- Name: mes_first_last_inspect_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_first_last_inspect_id_seq', 1, false);


--
-- Name: mes_flow_card_detail_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_flow_card_detail_id_seq', 1, false);


--
-- Name: mes_flow_card_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_flow_card_id_seq', 1, false);


--
-- Name: mes_job_report_log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_job_report_log_id_seq', 1, false);


--
-- Name: mes_order_day_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_order_day_id_seq', 1, false);


--
-- Name: mes_order_day_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_order_day_item_id_seq', 1, false);


--
-- Name: mes_order_month_audit_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_order_month_audit_id_seq', 1, false);


--
-- Name: mes_order_month_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_order_month_id_seq', 1, false);


--
-- Name: mes_order_month_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_order_month_item_id_seq', 1, false);


--
-- Name: mes_package_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_package_id_seq', 1, false);


--
-- Name: mes_person_skill_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_person_skill_id_seq', 1, false);


--
-- Name: mes_person_skill_score_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_person_skill_score_id_seq', 1, false);


--
-- Name: mes_process_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_process_id_seq', 1, false);


--
-- Name: mes_process_operation_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_process_operation_id_seq', 1, false);


--
-- Name: mes_team_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_team_id_seq', 1, false);


--
-- Name: mes_team_member_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_team_member_id_seq', 1, false);


--
-- Name: mes_team_shift_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_team_shift_id_seq', 1, false);


--
-- Name: mes_work_scheduling_detail_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_work_scheduling_detail_id_seq', 1, false);


--
-- Name: mes_work_scheduling_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mes_work_scheduling_id_seq', 1, false);


--
-- Name: mps_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mps_id_seq', 1, false);


--
-- Name: mrp_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mrp_items_id_seq', 1, false);


--
-- Name: mrps_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.mrps_id_seq', 1, false);


--
-- Name: ncrs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.ncrs_id_seq', 1, false);


--
-- Name: notification_channel_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.notification_channel_id_seq', 1, false);


--
-- Name: oee_report_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.oee_report_id_seq', 1, false);


--
-- Name: oper_logs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.oper_logs_id_seq', 1, false);


--
-- Name: oqcs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.oqcs_id_seq', 1, false);


--
-- Name: pro_dispatch_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.pro_dispatch_id_seq', 1, false);


--
-- Name: pro_production_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.pro_production_order_id_seq', 1, false);


--
-- Name: pro_production_report_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.pro_production_report_id_seq', 1, false);


--
-- Name: pro_sales_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.pro_sales_order_id_seq', 1, false);


--
-- Name: pro_sales_order_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.pro_sales_order_item_id_seq', 1, false);


--
-- Name: processes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.processes_id_seq', 7, true);


--
-- Name: production_daily_report_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.production_daily_report_id_seq', 1, false);


--
-- Name: production_lines_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.production_lines_id_seq', 1, false);


--
-- Name: production_offline_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.production_offline_id_seq', 1, false);


--
-- Name: production_offline_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.production_offline_item_id_seq', 1, false);


--
-- Name: production_orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.production_orders_id_seq', 1, false);


--
-- Name: production_reports_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.production_reports_id_seq', 1, false);


--
-- Name: qc_aql_level_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.qc_aql_level_id_seq', 1, false);


--
-- Name: qc_aql_table_row_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.qc_aql_table_row_id_seq', 1, false);


--
-- Name: qc_defect_code_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.qc_defect_code_id_seq', 1, false);


--
-- Name: qc_defect_record_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.qc_defect_record_id_seq', 1, false);


--
-- Name: qc_fqc_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.qc_fqc_id_seq', 1, false);


--
-- Name: qc_inspection_characteristic_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.qc_inspection_characteristic_id_seq', 1, false);


--
-- Name: qc_inspection_feature_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.qc_inspection_feature_id_seq', 1, false);


--
-- Name: qc_ipqc_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.qc_ipqc_id_seq', 1, false);


--
-- Name: qc_iqc_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.qc_iqc_id_seq', 1, false);


--
-- Name: qc_iqc_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.qc_iqc_item_id_seq', 1, false);


--
-- Name: qc_ncr_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.qc_ncr_id_seq', 1, false);


--
-- Name: qc_oqc_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.qc_oqc_id_seq', 1, false);


--
-- Name: qc_sampling_plan_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.qc_sampling_plan_id_seq', 1, false);


--
-- Name: qc_spc_data_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.qc_spc_data_id_seq', 1, false);


--
-- Name: qms_sampling_plan_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.qms_sampling_plan_id_seq', 1, false);


--
-- Name: qms_sampling_record_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.qms_sampling_record_id_seq', 1, false);


--
-- Name: qms_sampling_rule_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.qms_sampling_rule_id_seq', 1, false);


--
-- Name: quality_weekly_report_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.quality_weekly_report_id_seq', 1, false);


--
-- Name: receive_order_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.receive_order_items_id_seq', 1, false);


--
-- Name: receive_orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.receive_orders_id_seq', 1, false);


--
-- Name: resources_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.resources_id_seq', 1, false);


--
-- Name: route_operations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.route_operations_id_seq', 14, true);


--
-- Name: routes_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.routes_id_seq', 2, true);


--
-- Name: sales_order_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sales_order_items_id_seq', 1, false);


--
-- Name: sales_orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sales_orders_id_seq', 1, false);


--
-- Name: schedule_plans_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.schedule_plans_id_seq', 1, true);


--
-- Name: schedule_results_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.schedule_results_id_seq', 1, false);


--
-- Name: scp_customer_inquiry_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_customer_inquiry_id_seq', 1, false);


--
-- Name: scp_inquiry_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_inquiry_item_id_seq', 1, false);


--
-- Name: scp_mrs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_mrs_id_seq', 1, false);


--
-- Name: scp_mrs_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_mrs_item_id_seq', 1, false);


--
-- Name: scp_po_change_log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_po_change_log_id_seq', 1, false);


--
-- Name: scp_purchase_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_purchase_order_id_seq', 11, true);


--
-- Name: scp_purchase_order_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_purchase_order_item_id_seq', 20, true);


--
-- Name: scp_purchase_plan_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_purchase_plan_id_seq', 1, false);


--
-- Name: scp_purchase_plan_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_purchase_plan_item_id_seq', 1, false);


--
-- Name: scp_qad_sync_log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_qad_sync_log_id_seq', 1, false);


--
-- Name: scp_quote_comparison_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_quote_comparison_id_seq', 1, false);


--
-- Name: scp_quote_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_quote_item_id_seq', 20, true);


--
-- Name: scp_rfq_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_rfq_id_seq', 11, true);


--
-- Name: scp_rfq_invite_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_rfq_invite_id_seq', 14, true);


--
-- Name: scp_rfq_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_rfq_item_id_seq', 20, true);


--
-- Name: scp_sales_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_sales_order_id_seq', 11, true);


--
-- Name: scp_sales_order_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_sales_order_item_id_seq', 20, true);


--
-- Name: scp_so_change_log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_so_change_log_id_seq', 1, false);


--
-- Name: scp_supplier_bank_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_supplier_bank_id_seq', 1, false);


--
-- Name: scp_supplier_contact_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_supplier_contact_id_seq', 1, false);


--
-- Name: scp_supplier_delivery_record_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_supplier_delivery_record_id_seq', 1, false);


--
-- Name: scp_supplier_grade_standard_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_supplier_grade_standard_id_seq', 1, false);


--
-- Name: scp_supplier_kpi_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_supplier_kpi_id_seq', 20, true);


--
-- Name: scp_supplier_purchase_info_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_supplier_purchase_info_id_seq', 1, false);


--
-- Name: scp_supplier_quality_record_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_supplier_quality_record_id_seq', 1, false);


--
-- Name: scp_supplier_quote_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.scp_supplier_quote_id_seq', 11, true);


--
-- Name: serial_numbers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.serial_numbers_id_seq', 1, false);


--
-- Name: shifts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.shifts_id_seq', 1, false);


--
-- Name: spc_data_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.spc_data_id_seq', 1, false);


--
-- Name: stock_checks_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.stock_checks_id_seq', 1, false);


--
-- Name: supplier_asn_items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.supplier_asn_items_id_seq', 1, false);


--
-- Name: supplier_asns_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.supplier_asns_id_seq', 1, false);


--
-- Name: sys_dept_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_dept_id_seq', 43, true);


--
-- Name: sys_dict_data_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_dict_data_id_seq', 42, true);


--
-- Name: sys_dict_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_dict_type_id_seq', 21, true);


--
-- Name: sys_interface_config_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_interface_config_id_seq', 1, false);


--
-- Name: sys_interface_exec_log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_interface_exec_log_id_seq', 1, false);


--
-- Name: sys_interface_field_map_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_interface_field_map_id_seq', 1, false);


--
-- Name: sys_interface_trigger_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_interface_trigger_id_seq', 1, false);


--
-- Name: sys_login_log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_login_log_id_seq', 27, true);


--
-- Name: sys_menu_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_menu_id_seq', 92, true);


--
-- Name: sys_notice_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_notice_id_seq', 1, false);


--
-- Name: sys_notice_read_record_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_notice_read_record_id_seq', 1, false);


--
-- Name: sys_oper_log_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_oper_log_id_seq', 24, true);


--
-- Name: sys_post_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_post_id_seq', 46, true);


--
-- Name: sys_print_template_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_print_template_id_seq', 1, false);


--
-- Name: sys_role_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_role_id_seq', 22, true);


--
-- Name: sys_tenant_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_tenant_id_seq', 21, true);


--
-- Name: sys_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sys_user_id_seq', 2, true);


--
-- Name: tra_andon_call_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tra_andon_call_id_seq', 1, false);


--
-- Name: tra_data_collection_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tra_data_collection_id_seq', 1, false);


--
-- Name: tra_serial_number_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tra_serial_number_id_seq', 1, false);


--
-- Name: tra_trace_record_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tra_trace_record_id_seq', 1, false);


--
-- Name: trace_records_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.trace_records_id_seq', 1, false);


--
-- Name: transfer_orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transfer_orders_id_seq', 1, false);


--
-- Name: visual_inspection_result_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.visual_inspection_result_id_seq', 1, false);


--
-- Name: visual_inspection_task_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.visual_inspection_task_id_seq', 1, false);


--
-- Name: warehouses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.warehouses_id_seq', 1, false);


--
-- Name: wms_agv_device_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_agv_device_id_seq', 1, false);


--
-- Name: wms_agv_location_mapping_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_agv_location_mapping_id_seq', 1, false);


--
-- Name: wms_agv_task_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_agv_task_id_seq', 1, false);


--
-- Name: wms_delivery_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_delivery_order_id_seq', 1, false);


--
-- Name: wms_delivery_order_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_delivery_order_item_id_seq', 1, false);


--
-- Name: wms_inventory_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_inventory_id_seq', 6, true);


--
-- Name: wms_inventory_record_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_inventory_record_id_seq', 1, false);


--
-- Name: wms_kanban_pull_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_kanban_pull_id_seq', 1, false);


--
-- Name: wms_location_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_location_id_seq', 8, true);


--
-- Name: wms_production_complete_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_production_complete_id_seq', 1, false);


--
-- Name: wms_production_complete_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_production_complete_item_id_seq', 1, false);


--
-- Name: wms_production_issue_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_production_issue_id_seq', 1, false);


--
-- Name: wms_production_issue_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_production_issue_item_id_seq', 1, false);


--
-- Name: wms_production_return_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_production_return_id_seq', 1, false);


--
-- Name: wms_production_return_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_production_return_item_id_seq', 1, false);


--
-- Name: wms_production_stock_in_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_production_stock_in_id_seq', 1, false);


--
-- Name: wms_production_stock_in_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_production_stock_in_item_id_seq', 1, false);


--
-- Name: wms_purchase_return_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_purchase_return_id_seq', 1, false);


--
-- Name: wms_purchase_return_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_purchase_return_item_id_seq', 1, false);


--
-- Name: wms_receive_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_receive_order_id_seq', 1, false);


--
-- Name: wms_receive_order_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_receive_order_item_id_seq', 1, false);


--
-- Name: wms_sales_return_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_sales_return_id_seq', 1, false);


--
-- Name: wms_sales_return_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_sales_return_item_id_seq', 1, false);


--
-- Name: wms_side_location_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_side_location_id_seq', 1, false);


--
-- Name: wms_stock_check_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_stock_check_id_seq', 1, false);


--
-- Name: wms_transfer_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_transfer_order_id_seq', 1, false);


--
-- Name: wms_warehouse_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.wms_warehouse_id_seq', 8, true);


--
-- Name: work_centers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.work_centers_id_seq', 1, false);


--
-- Name: workshops_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.workshops_id_seq', 1, false);


--
-- Name: workstations_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.workstations_id_seq', 1, false);


--
-- Name: ai_chat_conversation ai_chat_conversation_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ai_chat_conversation
    ADD CONSTRAINT ai_chat_conversation_pkey PRIMARY KEY (id);


--
-- Name: ai_chat_message ai_chat_message_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ai_chat_message
    ADD CONSTRAINT ai_chat_message_pkey PRIMARY KEY (id);


--
-- Name: ai_config ai_config_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ai_config
    ADD CONSTRAINT ai_config_pkey PRIMARY KEY (id);


--
-- Name: alert_escalation_rule alert_escalation_rule_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alert_escalation_rule
    ADD CONSTRAINT alert_escalation_rule_pkey PRIMARY KEY (id);


--
-- Name: alert_notification_log alert_notification_log_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alert_notification_log
    ADD CONSTRAINT alert_notification_log_pkey PRIMARY KEY (id);


--
-- Name: alert_record alert_record_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alert_record
    ADD CONSTRAINT alert_record_pkey PRIMARY KEY (id);


--
-- Name: alert_rule alert_rule_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.alert_rule
    ADD CONSTRAINT alert_rule_pkey PRIMARY KEY (id);


--
-- Name: andon_call andon_call_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.andon_call
    ADD CONSTRAINT andon_call_pkey PRIMARY KEY (id);


--
-- Name: andon_calls andon_calls_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.andon_calls
    ADD CONSTRAINT andon_calls_pkey PRIMARY KEY (id);


--
-- Name: andon_escalation_log andon_escalation_log_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.andon_escalation_log
    ADD CONSTRAINT andon_escalation_log_pkey PRIMARY KEY (id);


--
-- Name: andon_escalation_rule andon_escalation_rule_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.andon_escalation_rule
    ADD CONSTRAINT andon_escalation_rule_pkey PRIMARY KEY (id);


--
-- Name: andon_notification_log andon_notification_log_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.andon_notification_log
    ADD CONSTRAINT andon_notification_log_pkey PRIMARY KEY (id);


--
-- Name: andon_report andon_report_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.andon_report
    ADD CONSTRAINT andon_report_pkey PRIMARY KEY (id);


--
-- Name: aps_capacity_analysis aps_capacity_analysis_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_capacity_analysis
    ADD CONSTRAINT aps_capacity_analysis_pkey PRIMARY KEY (id);


--
-- Name: aps_changeover_matrix aps_changeover_matrix_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_changeover_matrix
    ADD CONSTRAINT aps_changeover_matrix_pkey PRIMARY KEY (id);


--
-- Name: aps_delivery_rate aps_delivery_rate_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_delivery_rate
    ADD CONSTRAINT aps_delivery_rate_pkey PRIMARY KEY (id);


--
-- Name: aps_jit_demand aps_jit_demand_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_jit_demand
    ADD CONSTRAINT aps_jit_demand_pkey PRIMARY KEY (id);


--
-- Name: aps_mps aps_mps_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_mps
    ADD CONSTRAINT aps_mps_pkey PRIMARY KEY (id);


--
-- Name: aps_mps aps_mps_tenant_id_mps_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_mps
    ADD CONSTRAINT aps_mps_tenant_id_mps_no_key UNIQUE (tenant_id, mps_no);


--
-- Name: aps_mrp_item aps_mrp_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_mrp_item
    ADD CONSTRAINT aps_mrp_item_pkey PRIMARY KEY (id);


--
-- Name: aps_mrp aps_mrp_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_mrp
    ADD CONSTRAINT aps_mrp_pkey PRIMARY KEY (id);


--
-- Name: aps_mrp aps_mrp_tenant_id_mrp_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_mrp
    ADD CONSTRAINT aps_mrp_tenant_id_mrp_no_key UNIQUE (tenant_id, mrp_no);


--
-- Name: aps_resource aps_resource_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_resource
    ADD CONSTRAINT aps_resource_pkey PRIMARY KEY (id);


--
-- Name: aps_resource aps_resource_tenant_id_resource_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_resource
    ADD CONSTRAINT aps_resource_tenant_id_resource_code_key UNIQUE (tenant_id, resource_code);


--
-- Name: aps_rolling_schedule aps_rolling_schedule_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_rolling_schedule
    ADD CONSTRAINT aps_rolling_schedule_pkey PRIMARY KEY (id);


--
-- Name: aps_schedule_plan aps_schedule_plan_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_schedule_plan
    ADD CONSTRAINT aps_schedule_plan_pkey PRIMARY KEY (id);


--
-- Name: aps_schedule_plan aps_schedule_plan_tenant_id_plan_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_schedule_plan
    ADD CONSTRAINT aps_schedule_plan_tenant_id_plan_no_key UNIQUE (tenant_id, plan_no);


--
-- Name: aps_schedule_result aps_schedule_result_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_schedule_result
    ADD CONSTRAINT aps_schedule_result_pkey PRIMARY KEY (id);


--
-- Name: aps_work_center aps_work_center_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_work_center
    ADD CONSTRAINT aps_work_center_pkey PRIMARY KEY (id);


--
-- Name: aps_work_center aps_work_center_tenant_id_work_center_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_work_center
    ADD CONSTRAINT aps_work_center_tenant_id_work_center_code_key UNIQUE (tenant_id, work_center_code);


--
-- Name: aps_working_calendar aps_working_calendar_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_working_calendar
    ADD CONSTRAINT aps_working_calendar_pkey PRIMARY KEY (id);


--
-- Name: bom_items bom_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bom_items
    ADD CONSTRAINT bom_items_pkey PRIMARY KEY (id);


--
-- Name: boms boms_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.boms
    ADD CONSTRAINT boms_pkey PRIMARY KEY (id);


--
-- Name: bpm_approval_record bpm_approval_record_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_approval_record
    ADD CONSTRAINT bpm_approval_record_pkey PRIMARY KEY (id);


--
-- Name: bpm_biz_mapping bpm_biz_mapping_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_biz_mapping
    ADD CONSTRAINT bpm_biz_mapping_pkey PRIMARY KEY (id);


--
-- Name: bpm_delegate_record bpm_delegate_record_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_delegate_record
    ADD CONSTRAINT bpm_delegate_record_pkey PRIMARY KEY (id);


--
-- Name: bpm_form_definition bpm_form_definition_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_form_definition
    ADD CONSTRAINT bpm_form_definition_pkey PRIMARY KEY (id);


--
-- Name: bpm_form_field bpm_form_field_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_form_field
    ADD CONSTRAINT bpm_form_field_pkey PRIMARY KEY (id);


--
-- Name: bpm_form_version bpm_form_version_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_form_version
    ADD CONSTRAINT bpm_form_version_pkey PRIMARY KEY (id);


--
-- Name: bpm_model_version bpm_model_version_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_model_version
    ADD CONSTRAINT bpm_model_version_pkey PRIMARY KEY (id);


--
-- Name: bpm_node_definition bpm_node_definition_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_node_definition
    ADD CONSTRAINT bpm_node_definition_pkey PRIMARY KEY (id);


--
-- Name: bpm_node_timeout bpm_node_timeout_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_node_timeout
    ADD CONSTRAINT bpm_node_timeout_pkey PRIMARY KEY (id);


--
-- Name: bpm_order_day_workorder_map bpm_order_day_workorder_map_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_order_day_workorder_map
    ADD CONSTRAINT bpm_order_day_workorder_map_pkey PRIMARY KEY (id);


--
-- Name: bpm_process_instance bpm_process_instance_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_process_instance
    ADD CONSTRAINT bpm_process_instance_pkey PRIMARY KEY (id);


--
-- Name: bpm_process_model bpm_process_model_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_process_model
    ADD CONSTRAINT bpm_process_model_pkey PRIMARY KEY (id);


--
-- Name: bpm_reminder_record bpm_reminder_record_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_reminder_record
    ADD CONSTRAINT bpm_reminder_record_pkey PRIMARY KEY (id);


--
-- Name: bpm_sequence_flow bpm_sequence_flow_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_sequence_flow
    ADD CONSTRAINT bpm_sequence_flow_pkey PRIMARY KEY (id);


--
-- Name: bpm_task_assignment bpm_task_assignment_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_task_assignment
    ADD CONSTRAINT bpm_task_assignment_pkey PRIMARY KEY (id);


--
-- Name: bpm_task_instance bpm_task_instance_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_task_instance
    ADD CONSTRAINT bpm_task_instance_pkey PRIMARY KEY (id);


--
-- Name: container_master container_master_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.container_master
    ADD CONSTRAINT container_master_pkey PRIMARY KEY (id);


--
-- Name: container_movement container_movement_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.container_movement
    ADD CONSTRAINT container_movement_pkey PRIMARY KEY (id);


--
-- Name: data_collections data_collections_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.data_collections
    ADD CONSTRAINT data_collections_pkey PRIMARY KEY (id);


--
-- Name: dc_collect_record dc_collect_record_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dc_collect_record
    ADD CONSTRAINT dc_collect_record_pkey PRIMARY KEY (id);


--
-- Name: dc_data_point dc_data_point_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dc_data_point
    ADD CONSTRAINT dc_data_point_pkey PRIMARY KEY (id);


--
-- Name: dc_scan_log dc_scan_log_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dc_scan_log
    ADD CONSTRAINT dc_scan_log_pkey PRIMARY KEY (id);


--
-- Name: defect_codes defect_codes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.defect_codes
    ADD CONSTRAINT defect_codes_pkey PRIMARY KEY (id);


--
-- Name: defect_records defect_records_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.defect_records
    ADD CONSTRAINT defect_records_pkey PRIMARY KEY (id);


--
-- Name: delivery_order_items delivery_order_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.delivery_order_items
    ADD CONSTRAINT delivery_order_items_pkey PRIMARY KEY (id);


--
-- Name: delivery_orders delivery_orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.delivery_orders
    ADD CONSTRAINT delivery_orders_pkey PRIMARY KEY (id);


--
-- Name: delivery_report delivery_report_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.delivery_report
    ADD CONSTRAINT delivery_report_pkey PRIMARY KEY (id);


--
-- Name: dispatches dispatches_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dispatches
    ADD CONSTRAINT dispatches_pkey PRIMARY KEY (id);


--
-- Name: eam_equipment_downtime eam_equipment_downtime_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.eam_equipment_downtime
    ADD CONSTRAINT eam_equipment_downtime_pkey PRIMARY KEY (id);


--
-- Name: eam_repair_flow eam_repair_flow_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.eam_repair_flow
    ADD CONSTRAINT eam_repair_flow_pkey PRIMARY KEY (id);


--
-- Name: eam_repair_job eam_repair_job_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.eam_repair_job
    ADD CONSTRAINT eam_repair_job_pkey PRIMARY KEY (id);


--
-- Name: eam_repair_std eam_repair_std_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.eam_repair_std
    ADD CONSTRAINT eam_repair_std_pkey PRIMARY KEY (id);


--
-- Name: ene_energy_record ene_energy_record_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ene_energy_record
    ADD CONSTRAINT ene_energy_record_pkey PRIMARY KEY (id);


--
-- Name: energy_records energy_records_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.energy_records
    ADD CONSTRAINT energy_records_pkey PRIMARY KEY (id);


--
-- Name: equ_equipment_check equ_equipment_check_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_equipment_check
    ADD CONSTRAINT equ_equipment_check_pkey PRIMARY KEY (id);


--
-- Name: equ_equipment_document equ_equipment_document_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_equipment_document
    ADD CONSTRAINT equ_equipment_document_pkey PRIMARY KEY (id);


--
-- Name: equ_equipment_maintenance equ_equipment_maintenance_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_equipment_maintenance
    ADD CONSTRAINT equ_equipment_maintenance_pkey PRIMARY KEY (id);


--
-- Name: equ_equipment_part equ_equipment_part_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_equipment_part
    ADD CONSTRAINT equ_equipment_part_pkey PRIMARY KEY (id);


--
-- Name: equ_equipment equ_equipment_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_equipment
    ADD CONSTRAINT equ_equipment_pkey PRIMARY KEY (id);


--
-- Name: equ_equipment_repair equ_equipment_repair_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_equipment_repair
    ADD CONSTRAINT equ_equipment_repair_pkey PRIMARY KEY (id);


--
-- Name: equ_equipment equ_equipment_tenant_id_equipment_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_equipment
    ADD CONSTRAINT equ_equipment_tenant_id_equipment_code_key UNIQUE (tenant_id, equipment_code);


--
-- Name: equ_gauge_calibration equ_gauge_calibration_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_gauge_calibration
    ADD CONSTRAINT equ_gauge_calibration_pkey PRIMARY KEY (id);


--
-- Name: equ_gauge equ_gauge_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_gauge
    ADD CONSTRAINT equ_gauge_pkey PRIMARY KEY (id);


--
-- Name: equ_mold_maintenance equ_mold_maintenance_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_mold_maintenance
    ADD CONSTRAINT equ_mold_maintenance_pkey PRIMARY KEY (id);


--
-- Name: equ_mold equ_mold_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_mold
    ADD CONSTRAINT equ_mold_pkey PRIMARY KEY (id);


--
-- Name: equ_mold_repair equ_mold_repair_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_mold_repair
    ADD CONSTRAINT equ_mold_repair_pkey PRIMARY KEY (id);


--
-- Name: equ_oee equ_oee_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_oee
    ADD CONSTRAINT equ_oee_pkey PRIMARY KEY (id);


--
-- Name: equ_oee_record equ_oee_record_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_oee_record
    ADD CONSTRAINT equ_oee_record_pkey PRIMARY KEY (id);


--
-- Name: equ_spare_part equ_spare_part_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_spare_part
    ADD CONSTRAINT equ_spare_part_pkey PRIMARY KEY (id);


--
-- Name: equ_spare_part equ_spare_part_tenant_id_spare_part_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_spare_part
    ADD CONSTRAINT equ_spare_part_tenant_id_spare_part_code_key UNIQUE (tenant_id, spare_part_code);


--
-- Name: equ_teep_data equ_teep_data_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_teep_data
    ADD CONSTRAINT equ_teep_data_pkey PRIMARY KEY (id);


--
-- Name: fin_payment_request fin_payment_request_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fin_payment_request
    ADD CONSTRAINT fin_payment_request_pkey PRIMARY KEY (id);


--
-- Name: fin_purchase_advance fin_purchase_advance_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fin_purchase_advance
    ADD CONSTRAINT fin_purchase_advance_pkey PRIMARY KEY (id);


--
-- Name: fin_purchase_settlement_item fin_purchase_settlement_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fin_purchase_settlement_item
    ADD CONSTRAINT fin_purchase_settlement_item_pkey PRIMARY KEY (id);


--
-- Name: fin_purchase_settlement fin_purchase_settlement_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fin_purchase_settlement
    ADD CONSTRAINT fin_purchase_settlement_pkey PRIMARY KEY (id);


--
-- Name: fin_sales_receipt fin_sales_receipt_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fin_sales_receipt
    ADD CONSTRAINT fin_sales_receipt_pkey PRIMARY KEY (id);


--
-- Name: fin_sales_settlement_item fin_sales_settlement_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fin_sales_settlement_item
    ADD CONSTRAINT fin_sales_settlement_item_pkey PRIMARY KEY (id);


--
-- Name: fin_sales_settlement fin_sales_settlement_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fin_sales_settlement
    ADD CONSTRAINT fin_sales_settlement_pkey PRIMARY KEY (id);


--
-- Name: fin_statement_detail fin_statement_detail_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fin_statement_detail
    ADD CONSTRAINT fin_statement_detail_pkey PRIMARY KEY (id);


--
-- Name: fin_supplier_statement fin_supplier_statement_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fin_supplier_statement
    ADD CONSTRAINT fin_supplier_statement_pkey PRIMARY KEY (id);


--
-- Name: fqcs fqcs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fqcs
    ADD CONSTRAINT fqcs_pkey PRIMARY KEY (id);


--
-- Name: ai_chat_conversation idx_ai_chat_conversation_session_id; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ai_chat_conversation
    ADD CONSTRAINT idx_ai_chat_conversation_session_id UNIQUE (session_id);


--
-- Name: andon_call idx_andon_call_call_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.andon_call
    ADD CONSTRAINT idx_andon_call_call_no UNIQUE (call_no);


--
-- Name: andon_calls idx_andon_calls_call_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.andon_calls
    ADD CONSTRAINT idx_andon_calls_call_no UNIQUE (call_no);


--
-- Name: aps_mps idx_aps_mps_mps_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_mps
    ADD CONSTRAINT idx_aps_mps_mps_no UNIQUE (mps_no);


--
-- Name: aps_mrp idx_aps_mrp_mrp_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_mrp
    ADD CONSTRAINT idx_aps_mrp_mrp_no UNIQUE (mrp_no);


--
-- Name: aps_resource idx_aps_resource_resource_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_resource
    ADD CONSTRAINT idx_aps_resource_resource_code UNIQUE (resource_code);


--
-- Name: aps_schedule_plan idx_aps_schedule_plan_plan_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_schedule_plan
    ADD CONSTRAINT idx_aps_schedule_plan_plan_no UNIQUE (plan_no);


--
-- Name: aps_work_center idx_aps_work_center_work_center_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.aps_work_center
    ADD CONSTRAINT idx_aps_work_center_work_center_code UNIQUE (work_center_code);


--
-- Name: boms idx_boms_bom_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.boms
    ADD CONSTRAINT idx_boms_bom_code UNIQUE (bom_code);


--
-- Name: dc_data_point idx_dc_data_point_point_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dc_data_point
    ADD CONSTRAINT idx_dc_data_point_point_code UNIQUE (point_code);


--
-- Name: defect_codes idx_defect_codes_defect_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.defect_codes
    ADD CONSTRAINT idx_defect_codes_defect_code UNIQUE (defect_code);


--
-- Name: defect_records idx_defect_records_record_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.defect_records
    ADD CONSTRAINT idx_defect_records_record_no UNIQUE (record_no);


--
-- Name: delivery_orders idx_delivery_orders_delivery_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.delivery_orders
    ADD CONSTRAINT idx_delivery_orders_delivery_no UNIQUE (delivery_no);


--
-- Name: eam_repair_flow idx_eam_repair_flow_flow_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.eam_repair_flow
    ADD CONSTRAINT idx_eam_repair_flow_flow_code UNIQUE (flow_code);


--
-- Name: eam_repair_job idx_eam_repair_job_job_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.eam_repair_job
    ADD CONSTRAINT idx_eam_repair_job_job_code UNIQUE (job_code);


--
-- Name: eam_repair_std idx_eam_repair_std_std_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.eam_repair_std
    ADD CONSTRAINT idx_eam_repair_std_std_code UNIQUE (std_code);


--
-- Name: equ_gauge idx_equ_gauge_gauge_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_gauge
    ADD CONSTRAINT idx_equ_gauge_gauge_code UNIQUE (gauge_code);


--
-- Name: equ_mold idx_equ_mold_mold_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.equ_mold
    ADD CONSTRAINT idx_equ_mold_mold_code UNIQUE (mold_code);


--
-- Name: fqcs idx_fqcs_fqc_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.fqcs
    ADD CONSTRAINT idx_fqcs_fqc_no UNIQUE (fqc_no);


--
-- Name: inventory_records idx_inventory_records_record_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.inventory_records
    ADD CONSTRAINT idx_inventory_records_record_no UNIQUE (record_no);


--
-- Name: ip_qcs idx_ip_qcs_ip_qc_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ip_qcs
    ADD CONSTRAINT idx_ip_qcs_ip_qc_no UNIQUE (ip_qc_no);


--
-- Name: iqcs idx_iqcs_iqc_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.iqcs
    ADD CONSTRAINT idx_iqcs_iqc_no UNIQUE (iqc_no);


--
-- Name: locations idx_locations_location_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.locations
    ADD CONSTRAINT idx_locations_location_code UNIQUE (location_code);


--
-- Name: materials idx_materials_material_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.materials
    ADD CONSTRAINT idx_materials_material_code UNIQUE (material_code);


--
-- Name: mdm_bom idx_mdm_bom_bom_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_bom
    ADD CONSTRAINT idx_mdm_bom_bom_code UNIQUE (bom_code);


--
-- Name: mdm_material idx_mdm_material_material_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_material
    ADD CONSTRAINT idx_mdm_material_material_code UNIQUE (material_code);


--
-- Name: mdm_operation idx_mdm_operation_operation_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_operation
    ADD CONSTRAINT idx_mdm_operation_operation_code UNIQUE (operation_code);


--
-- Name: mdm_production_line idx_mdm_production_line_line_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_production_line
    ADD CONSTRAINT idx_mdm_production_line_line_code UNIQUE (line_code);


--
-- Name: mdm_shift idx_mdm_shift_shift_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_shift
    ADD CONSTRAINT idx_mdm_shift_shift_code UNIQUE (shift_code);


--
-- Name: mdm_supplier idx_mdm_supplier_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_supplier
    ADD CONSTRAINT idx_mdm_supplier_code UNIQUE (code);


--
-- Name: mdm_workshop idx_mdm_workshop_workshop_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_workshop
    ADD CONSTRAINT idx_mdm_workshop_workshop_code UNIQUE (workshop_code);


--
-- Name: mdm_workstation idx_mdm_workstation_station_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_workstation
    ADD CONSTRAINT idx_mdm_workstation_station_code UNIQUE (station_code);


--
-- Name: mes_process idx_mes_process_process_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_process
    ADD CONSTRAINT idx_mes_process_process_code UNIQUE (process_code);


--
-- Name: mes_team idx_mes_team_team_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_team
    ADD CONSTRAINT idx_mes_team_team_code UNIQUE (team_code);


--
-- Name: mps idx_mps_mps_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mps
    ADD CONSTRAINT idx_mps_mps_no UNIQUE (mps_no);


--
-- Name: mrps idx_mrps_mrp_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mrps
    ADD CONSTRAINT idx_mrps_mrp_no UNIQUE (mrp_no);


--
-- Name: ncrs idx_ncrs_ncr_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ncrs
    ADD CONSTRAINT idx_ncrs_ncr_no UNIQUE (ncr_no);


--
-- Name: oqcs idx_oqcs_oqc_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.oqcs
    ADD CONSTRAINT idx_oqcs_oqc_no UNIQUE (oqc_no);


--
-- Name: pro_production_order idx_pro_production_order_order_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pro_production_order
    ADD CONSTRAINT idx_pro_production_order_order_no UNIQUE (order_no);


--
-- Name: pro_sales_order idx_pro_sales_order_order_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pro_sales_order
    ADD CONSTRAINT idx_pro_sales_order_order_no UNIQUE (order_no);


--
-- Name: processes idx_processes_process_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.processes
    ADD CONSTRAINT idx_processes_process_code UNIQUE (process_code);


--
-- Name: production_lines idx_production_lines_line_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.production_lines
    ADD CONSTRAINT idx_production_lines_line_code UNIQUE (line_code);


--
-- Name: production_orders idx_production_orders_order_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.production_orders
    ADD CONSTRAINT idx_production_orders_order_no UNIQUE (order_no);


--
-- Name: qc_defect_code idx_qc_defect_code_defect_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_defect_code
    ADD CONSTRAINT idx_qc_defect_code_defect_code UNIQUE (defect_code);


--
-- Name: qc_defect_record idx_qc_defect_record_record_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_defect_record
    ADD CONSTRAINT idx_qc_defect_record_record_no UNIQUE (record_no);


--
-- Name: qc_fqc idx_qc_fqc_fqc_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_fqc
    ADD CONSTRAINT idx_qc_fqc_fqc_no UNIQUE (fqc_no);


--
-- Name: qc_iqc idx_qc_iqc_iqc_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_iqc
    ADD CONSTRAINT idx_qc_iqc_iqc_no UNIQUE (iqc_no);


--
-- Name: qc_ncr idx_qc_ncr_ncr_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_ncr
    ADD CONSTRAINT idx_qc_ncr_ncr_no UNIQUE (ncr_no);


--
-- Name: qc_oqc idx_qc_oqc_oqc_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_oqc
    ADD CONSTRAINT idx_qc_oqc_oqc_no UNIQUE (oqc_no);


--
-- Name: qc_sampling_plan idx_qc_sampling_plan_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_sampling_plan
    ADD CONSTRAINT idx_qc_sampling_plan_code UNIQUE (code);


--
-- Name: qms_sampling_plan idx_qms_sampling_plan_plan_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qms_sampling_plan
    ADD CONSTRAINT idx_qms_sampling_plan_plan_code UNIQUE (plan_code);


--
-- Name: receive_orders idx_receive_orders_receive_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.receive_orders
    ADD CONSTRAINT idx_receive_orders_receive_no UNIQUE (receive_no);


--
-- Name: resources idx_resources_resource_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.resources
    ADD CONSTRAINT idx_resources_resource_code UNIQUE (resource_code);


--
-- Name: routes idx_routes_route_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.routes
    ADD CONSTRAINT idx_routes_route_code UNIQUE (route_code);


--
-- Name: sales_orders idx_sales_orders_order_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sales_orders
    ADD CONSTRAINT idx_sales_orders_order_no UNIQUE (order_no);


--
-- Name: schedule_plans idx_schedule_plans_plan_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedule_plans
    ADD CONSTRAINT idx_schedule_plans_plan_no UNIQUE (plan_no);


--
-- Name: serial_numbers idx_serial_numbers_serial_number; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.serial_numbers
    ADD CONSTRAINT idx_serial_numbers_serial_number UNIQUE (serial_number);


--
-- Name: shifts idx_shifts_shift_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shifts
    ADD CONSTRAINT idx_shifts_shift_code UNIQUE (shift_code);


--
-- Name: stock_checks idx_stock_checks_check_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stock_checks
    ADD CONSTRAINT idx_stock_checks_check_no UNIQUE (check_no);


--
-- Name: sys_interface_config idx_sys_interface_config_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_interface_config
    ADD CONSTRAINT idx_sys_interface_config_code UNIQUE (code);


--
-- Name: tra_andon_call idx_tra_andon_call_call_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tra_andon_call
    ADD CONSTRAINT idx_tra_andon_call_call_no UNIQUE (call_no);


--
-- Name: tra_serial_number idx_tra_serial_number_serial_number; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tra_serial_number
    ADD CONSTRAINT idx_tra_serial_number_serial_number UNIQUE (serial_number);


--
-- Name: tra_trace_record idx_tra_trace_record_trace_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tra_trace_record
    ADD CONSTRAINT idx_tra_trace_record_trace_no UNIQUE (trace_no);


--
-- Name: trace_records idx_trace_records_trace_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trace_records
    ADD CONSTRAINT idx_trace_records_trace_no UNIQUE (trace_no);


--
-- Name: transfer_orders idx_transfer_orders_transfer_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transfer_orders
    ADD CONSTRAINT idx_transfer_orders_transfer_no UNIQUE (transfer_no);


--
-- Name: warehouses idx_warehouses_warehouse_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.warehouses
    ADD CONSTRAINT idx_warehouses_warehouse_code UNIQUE (warehouse_code);


--
-- Name: wms_delivery_order idx_wms_delivery_order_delivery_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_delivery_order
    ADD CONSTRAINT idx_wms_delivery_order_delivery_no UNIQUE (delivery_no);


--
-- Name: wms_inventory_record idx_wms_inventory_record_record_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_inventory_record
    ADD CONSTRAINT idx_wms_inventory_record_record_no UNIQUE (record_no);


--
-- Name: wms_kanban_pull idx_wms_kanban_pull_kanban_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_kanban_pull
    ADD CONSTRAINT idx_wms_kanban_pull_kanban_no UNIQUE (kanban_no);


--
-- Name: wms_location idx_wms_location_location_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_location
    ADD CONSTRAINT idx_wms_location_location_code UNIQUE (location_code);


--
-- Name: wms_receive_order idx_wms_receive_order_receive_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_receive_order
    ADD CONSTRAINT idx_wms_receive_order_receive_no UNIQUE (receive_no);


--
-- Name: wms_side_location idx_wms_side_location_location_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_side_location
    ADD CONSTRAINT idx_wms_side_location_location_code UNIQUE (location_code);


--
-- Name: wms_stock_check idx_wms_stock_check_check_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_stock_check
    ADD CONSTRAINT idx_wms_stock_check_check_no UNIQUE (check_no);


--
-- Name: wms_transfer_order idx_wms_transfer_order_transfer_no; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_transfer_order
    ADD CONSTRAINT idx_wms_transfer_order_transfer_no UNIQUE (transfer_no);


--
-- Name: wms_warehouse idx_wms_warehouse_warehouse_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_warehouse
    ADD CONSTRAINT idx_wms_warehouse_warehouse_code UNIQUE (warehouse_code);


--
-- Name: work_centers idx_work_centers_work_center_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.work_centers
    ADD CONSTRAINT idx_work_centers_work_center_code UNIQUE (work_center_code);


--
-- Name: workshops idx_workshops_workshop_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workshops
    ADD CONSTRAINT idx_workshops_workshop_code UNIQUE (workshop_code);


--
-- Name: workstations idx_workstations_station_code; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workstations
    ADD CONSTRAINT idx_workstations_station_code UNIQUE (station_code);


--
-- Name: integration_erp_mappings integration_erp_mappings_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.integration_erp_mappings
    ADD CONSTRAINT integration_erp_mappings_pkey PRIMARY KEY (id);


--
-- Name: integration_erp_sync_logs integration_erp_sync_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.integration_erp_sync_logs
    ADD CONSTRAINT integration_erp_sync_logs_pkey PRIMARY KEY (id);


--
-- Name: inventories inventories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.inventories
    ADD CONSTRAINT inventories_pkey PRIMARY KEY (id);


--
-- Name: inventory_records inventory_records_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.inventory_records
    ADD CONSTRAINT inventory_records_pkey PRIMARY KEY (id);


--
-- Name: ip_qcs ip_qcs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ip_qcs
    ADD CONSTRAINT ip_qcs_pkey PRIMARY KEY (id);


--
-- Name: iqc_items iqc_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.iqc_items
    ADD CONSTRAINT iqc_items_pkey PRIMARY KEY (id);


--
-- Name: iqcs iqcs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.iqcs
    ADD CONSTRAINT iqcs_pkey PRIMARY KEY (id);


--
-- Name: lab_calibration lab_calibration_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lab_calibration
    ADD CONSTRAINT lab_calibration_pkey PRIMARY KEY (id);


--
-- Name: lab_instrument lab_instrument_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lab_instrument
    ADD CONSTRAINT lab_instrument_pkey PRIMARY KEY (id);


--
-- Name: lab_report lab_report_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lab_report
    ADD CONSTRAINT lab_report_pkey PRIMARY KEY (id);


--
-- Name: lab_sample lab_sample_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lab_sample
    ADD CONSTRAINT lab_sample_pkey PRIMARY KEY (id);


--
-- Name: lab_test_item lab_test_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lab_test_item
    ADD CONSTRAINT lab_test_item_pkey PRIMARY KEY (id);


--
-- Name: locations locations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.locations
    ADD CONSTRAINT locations_pkey PRIMARY KEY (id);


--
-- Name: login_logs login_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.login_logs
    ADD CONSTRAINT login_logs_pkey PRIMARY KEY (id);


--
-- Name: material_categories material_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.material_categories
    ADD CONSTRAINT material_categories_pkey PRIMARY KEY (id);


--
-- Name: materials materials_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.materials
    ADD CONSTRAINT materials_pkey PRIMARY KEY (id);


--
-- Name: mdm_bom_item mdm_bom_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_bom_item
    ADD CONSTRAINT mdm_bom_item_pkey PRIMARY KEY (id);


--
-- Name: mdm_bom mdm_bom_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_bom
    ADD CONSTRAINT mdm_bom_pkey PRIMARY KEY (id);


--
-- Name: mdm_bom mdm_bom_tenant_id_bom_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_bom
    ADD CONSTRAINT mdm_bom_tenant_id_bom_code_key UNIQUE (tenant_id, bom_code);


--
-- Name: mdm_customer mdm_customer_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_customer
    ADD CONSTRAINT mdm_customer_pkey PRIMARY KEY (id);


--
-- Name: mdm_material_category mdm_material_category_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_material_category
    ADD CONSTRAINT mdm_material_category_pkey PRIMARY KEY (id);


--
-- Name: mdm_material mdm_material_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_material
    ADD CONSTRAINT mdm_material_pkey PRIMARY KEY (id);


--
-- Name: mdm_material mdm_material_tenant_id_material_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_material
    ADD CONSTRAINT mdm_material_tenant_id_material_code_key UNIQUE (tenant_id, material_code);


--
-- Name: mdm_operation mdm_operation_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_operation
    ADD CONSTRAINT mdm_operation_pkey PRIMARY KEY (id);


--
-- Name: mdm_operation mdm_operation_tenant_id_operation_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_operation
    ADD CONSTRAINT mdm_operation_tenant_id_operation_code_key UNIQUE (tenant_id, operation_code);


--
-- Name: mdm_production_line mdm_production_line_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_production_line
    ADD CONSTRAINT mdm_production_line_pkey PRIMARY KEY (id);


--
-- Name: mdm_production_line mdm_production_line_tenant_id_line_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_production_line
    ADD CONSTRAINT mdm_production_line_tenant_id_line_code_key UNIQUE (tenant_id, line_code);


--
-- Name: mdm_shift mdm_shift_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_shift
    ADD CONSTRAINT mdm_shift_pkey PRIMARY KEY (id);


--
-- Name: mdm_shift mdm_shift_tenant_id_shift_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_shift
    ADD CONSTRAINT mdm_shift_tenant_id_shift_code_key UNIQUE (tenant_id, shift_code);


--
-- Name: mdm_supplier_material mdm_supplier_material_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_supplier_material
    ADD CONSTRAINT mdm_supplier_material_pkey PRIMARY KEY (id);


--
-- Name: mdm_supplier mdm_supplier_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_supplier
    ADD CONSTRAINT mdm_supplier_pkey PRIMARY KEY (id);


--
-- Name: mdm_workshop_config mdm_workshop_config_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_workshop_config
    ADD CONSTRAINT mdm_workshop_config_pkey PRIMARY KEY (id);


--
-- Name: mdm_workshop mdm_workshop_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_workshop
    ADD CONSTRAINT mdm_workshop_pkey PRIMARY KEY (id);


--
-- Name: mdm_workshop mdm_workshop_tenant_id_workshop_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_workshop
    ADD CONSTRAINT mdm_workshop_tenant_id_workshop_code_key UNIQUE (tenant_id, workshop_code);


--
-- Name: mdm_workstation mdm_workstation_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_workstation
    ADD CONSTRAINT mdm_workstation_pkey PRIMARY KEY (id);


--
-- Name: mdm_workstation mdm_workstation_tenant_id_station_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mdm_workstation
    ADD CONSTRAINT mdm_workstation_tenant_id_station_code_key UNIQUE (tenant_id, station_code);


--
-- Name: mes_code_rule mes_code_rule_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_code_rule
    ADD CONSTRAINT mes_code_rule_pkey PRIMARY KEY (id);


--
-- Name: mes_code_rule_record mes_code_rule_record_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_code_rule_record
    ADD CONSTRAINT mes_code_rule_record_pkey PRIMARY KEY (id);


--
-- Name: mes_container_lifecycle mes_container_lifecycle_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_container_lifecycle
    ADD CONSTRAINT mes_container_lifecycle_pkey PRIMARY KEY (id);


--
-- Name: mes_container_maintenance mes_container_maintenance_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_container_maintenance
    ADD CONSTRAINT mes_container_maintenance_pkey PRIMARY KEY (id);


--
-- Name: mes_electronic_sop mes_electronic_sop_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_electronic_sop
    ADD CONSTRAINT mes_electronic_sop_pkey PRIMARY KEY (id);


--
-- Name: mes_first_last_inspect mes_first_last_inspect_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_first_last_inspect
    ADD CONSTRAINT mes_first_last_inspect_pkey PRIMARY KEY (id);


--
-- Name: mes_flow_card_detail mes_flow_card_detail_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_flow_card_detail
    ADD CONSTRAINT mes_flow_card_detail_pkey PRIMARY KEY (id);


--
-- Name: mes_flow_card mes_flow_card_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_flow_card
    ADD CONSTRAINT mes_flow_card_pkey PRIMARY KEY (id);


--
-- Name: mes_job_report_log mes_job_report_log_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_job_report_log
    ADD CONSTRAINT mes_job_report_log_pkey PRIMARY KEY (id);


--
-- Name: mes_order_day_item mes_order_day_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_order_day_item
    ADD CONSTRAINT mes_order_day_item_pkey PRIMARY KEY (id);


--
-- Name: mes_order_day mes_order_day_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_order_day
    ADD CONSTRAINT mes_order_day_pkey PRIMARY KEY (id);


--
-- Name: mes_order_month_audit mes_order_month_audit_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_order_month_audit
    ADD CONSTRAINT mes_order_month_audit_pkey PRIMARY KEY (id);


--
-- Name: mes_order_month_item mes_order_month_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_order_month_item
    ADD CONSTRAINT mes_order_month_item_pkey PRIMARY KEY (id);


--
-- Name: mes_order_month mes_order_month_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_order_month
    ADD CONSTRAINT mes_order_month_pkey PRIMARY KEY (id);


--
-- Name: mes_package mes_package_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_package
    ADD CONSTRAINT mes_package_pkey PRIMARY KEY (id);


--
-- Name: mes_person_skill mes_person_skill_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_person_skill
    ADD CONSTRAINT mes_person_skill_pkey PRIMARY KEY (id);


--
-- Name: mes_person_skill_score mes_person_skill_score_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_person_skill_score
    ADD CONSTRAINT mes_person_skill_score_pkey PRIMARY KEY (id);


--
-- Name: mes_process_operation mes_process_operation_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_process_operation
    ADD CONSTRAINT mes_process_operation_pkey PRIMARY KEY (id);


--
-- Name: mes_process mes_process_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_process
    ADD CONSTRAINT mes_process_pkey PRIMARY KEY (id);


--
-- Name: mes_team_member mes_team_member_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_team_member
    ADD CONSTRAINT mes_team_member_pkey PRIMARY KEY (id);


--
-- Name: mes_team mes_team_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_team
    ADD CONSTRAINT mes_team_pkey PRIMARY KEY (id);


--
-- Name: mes_team_shift mes_team_shift_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_team_shift
    ADD CONSTRAINT mes_team_shift_pkey PRIMARY KEY (id);


--
-- Name: mes_work_scheduling_detail mes_work_scheduling_detail_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_work_scheduling_detail
    ADD CONSTRAINT mes_work_scheduling_detail_pkey PRIMARY KEY (id);


--
-- Name: mes_work_scheduling mes_work_scheduling_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_work_scheduling
    ADD CONSTRAINT mes_work_scheduling_pkey PRIMARY KEY (id);


--
-- Name: mps mps_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mps
    ADD CONSTRAINT mps_pkey PRIMARY KEY (id);


--
-- Name: mrp_items mrp_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mrp_items
    ADD CONSTRAINT mrp_items_pkey PRIMARY KEY (id);


--
-- Name: mrps mrps_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mrps
    ADD CONSTRAINT mrps_pkey PRIMARY KEY (id);


--
-- Name: ncrs ncrs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.ncrs
    ADD CONSTRAINT ncrs_pkey PRIMARY KEY (id);


--
-- Name: notification_channel notification_channel_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.notification_channel
    ADD CONSTRAINT notification_channel_pkey PRIMARY KEY (id);


--
-- Name: oee_report oee_report_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.oee_report
    ADD CONSTRAINT oee_report_pkey PRIMARY KEY (id);


--
-- Name: oper_logs oper_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.oper_logs
    ADD CONSTRAINT oper_logs_pkey PRIMARY KEY (id);


--
-- Name: oqcs oqcs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.oqcs
    ADD CONSTRAINT oqcs_pkey PRIMARY KEY (id);


--
-- Name: pro_dispatch pro_dispatch_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pro_dispatch
    ADD CONSTRAINT pro_dispatch_pkey PRIMARY KEY (id);


--
-- Name: pro_production_order pro_production_order_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pro_production_order
    ADD CONSTRAINT pro_production_order_pkey PRIMARY KEY (id);


--
-- Name: pro_production_order pro_production_order_tenant_id_order_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pro_production_order
    ADD CONSTRAINT pro_production_order_tenant_id_order_no_key UNIQUE (tenant_id, order_no);


--
-- Name: pro_production_report pro_production_report_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pro_production_report
    ADD CONSTRAINT pro_production_report_pkey PRIMARY KEY (id);


--
-- Name: pro_sales_order_item pro_sales_order_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pro_sales_order_item
    ADD CONSTRAINT pro_sales_order_item_pkey PRIMARY KEY (id);


--
-- Name: pro_sales_order pro_sales_order_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pro_sales_order
    ADD CONSTRAINT pro_sales_order_pkey PRIMARY KEY (id);


--
-- Name: pro_sales_order pro_sales_order_tenant_id_order_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.pro_sales_order
    ADD CONSTRAINT pro_sales_order_tenant_id_order_no_key UNIQUE (tenant_id, order_no);


--
-- Name: processes processes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.processes
    ADD CONSTRAINT processes_pkey PRIMARY KEY (id);


--
-- Name: production_daily_report production_daily_report_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.production_daily_report
    ADD CONSTRAINT production_daily_report_pkey PRIMARY KEY (id);


--
-- Name: production_lines production_lines_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.production_lines
    ADD CONSTRAINT production_lines_pkey PRIMARY KEY (id);


--
-- Name: production_offline_item production_offline_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.production_offline_item
    ADD CONSTRAINT production_offline_item_pkey PRIMARY KEY (id);


--
-- Name: production_offline production_offline_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.production_offline
    ADD CONSTRAINT production_offline_pkey PRIMARY KEY (id);


--
-- Name: production_orders production_orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.production_orders
    ADD CONSTRAINT production_orders_pkey PRIMARY KEY (id);


--
-- Name: production_reports production_reports_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.production_reports
    ADD CONSTRAINT production_reports_pkey PRIMARY KEY (id);


--
-- Name: qc_aql_level qc_aql_level_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_aql_level
    ADD CONSTRAINT qc_aql_level_pkey PRIMARY KEY (id);


--
-- Name: qc_aql_table_row qc_aql_table_row_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_aql_table_row
    ADD CONSTRAINT qc_aql_table_row_pkey PRIMARY KEY (id);


--
-- Name: qc_defect_code qc_defect_code_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_defect_code
    ADD CONSTRAINT qc_defect_code_pkey PRIMARY KEY (id);


--
-- Name: qc_defect_code qc_defect_code_tenant_id_defect_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_defect_code
    ADD CONSTRAINT qc_defect_code_tenant_id_defect_code_key UNIQUE (tenant_id, defect_code);


--
-- Name: qc_defect_record qc_defect_record_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_defect_record
    ADD CONSTRAINT qc_defect_record_pkey PRIMARY KEY (id);


--
-- Name: qc_defect_record qc_defect_record_tenant_id_record_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_defect_record
    ADD CONSTRAINT qc_defect_record_tenant_id_record_no_key UNIQUE (tenant_id, record_no);


--
-- Name: qc_fqc qc_fqc_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_fqc
    ADD CONSTRAINT qc_fqc_pkey PRIMARY KEY (id);


--
-- Name: qc_fqc qc_fqc_tenant_id_fqc_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_fqc
    ADD CONSTRAINT qc_fqc_tenant_id_fqc_no_key UNIQUE (tenant_id, fqc_no);


--
-- Name: qc_inspection_characteristic qc_inspection_characteristic_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_inspection_characteristic
    ADD CONSTRAINT qc_inspection_characteristic_pkey PRIMARY KEY (id);


--
-- Name: qc_inspection_feature qc_inspection_feature_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_inspection_feature
    ADD CONSTRAINT qc_inspection_feature_pkey PRIMARY KEY (id);


--
-- Name: qc_ipqc qc_ipqc_ip_qc_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_ipqc
    ADD CONSTRAINT qc_ipqc_ip_qc_no_key UNIQUE (ip_qc_no);


--
-- Name: qc_ipqc qc_ipqc_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_ipqc
    ADD CONSTRAINT qc_ipqc_pkey PRIMARY KEY (id);


--
-- Name: qc_ipqc qc_ipqc_tenant_id_ipqc_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_ipqc
    ADD CONSTRAINT qc_ipqc_tenant_id_ipqc_no_key UNIQUE (tenant_id, ipqc_no);


--
-- Name: qc_iqc_item qc_iqc_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_iqc_item
    ADD CONSTRAINT qc_iqc_item_pkey PRIMARY KEY (id);


--
-- Name: qc_iqc qc_iqc_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_iqc
    ADD CONSTRAINT qc_iqc_pkey PRIMARY KEY (id);


--
-- Name: qc_iqc qc_iqc_tenant_id_iqc_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_iqc
    ADD CONSTRAINT qc_iqc_tenant_id_iqc_no_key UNIQUE (tenant_id, iqc_no);


--
-- Name: qc_ncr qc_ncr_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_ncr
    ADD CONSTRAINT qc_ncr_pkey PRIMARY KEY (id);


--
-- Name: qc_ncr qc_ncr_tenant_id_ncr_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_ncr
    ADD CONSTRAINT qc_ncr_tenant_id_ncr_no_key UNIQUE (tenant_id, ncr_no);


--
-- Name: qc_oqc qc_oqc_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_oqc
    ADD CONSTRAINT qc_oqc_pkey PRIMARY KEY (id);


--
-- Name: qc_oqc qc_oqc_tenant_id_oqc_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_oqc
    ADD CONSTRAINT qc_oqc_tenant_id_oqc_no_key UNIQUE (tenant_id, oqc_no);


--
-- Name: qc_sampling_plan qc_sampling_plan_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_sampling_plan
    ADD CONSTRAINT qc_sampling_plan_pkey PRIMARY KEY (id);


--
-- Name: qc_spc_data qc_spc_data_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qc_spc_data
    ADD CONSTRAINT qc_spc_data_pkey PRIMARY KEY (id);


--
-- Name: qms_sampling_plan qms_sampling_plan_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qms_sampling_plan
    ADD CONSTRAINT qms_sampling_plan_pkey PRIMARY KEY (id);


--
-- Name: qms_sampling_record qms_sampling_record_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qms_sampling_record
    ADD CONSTRAINT qms_sampling_record_pkey PRIMARY KEY (id);


--
-- Name: qms_sampling_rule qms_sampling_rule_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.qms_sampling_rule
    ADD CONSTRAINT qms_sampling_rule_pkey PRIMARY KEY (id);


--
-- Name: quality_weekly_report quality_weekly_report_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.quality_weekly_report
    ADD CONSTRAINT quality_weekly_report_pkey PRIMARY KEY (id);


--
-- Name: receive_order_items receive_order_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.receive_order_items
    ADD CONSTRAINT receive_order_items_pkey PRIMARY KEY (id);


--
-- Name: receive_orders receive_orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.receive_orders
    ADD CONSTRAINT receive_orders_pkey PRIMARY KEY (id);


--
-- Name: resources resources_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.resources
    ADD CONSTRAINT resources_pkey PRIMARY KEY (id);


--
-- Name: route_operations route_operations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.route_operations
    ADD CONSTRAINT route_operations_pkey PRIMARY KEY (id);


--
-- Name: routes routes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.routes
    ADD CONSTRAINT routes_pkey PRIMARY KEY (id);


--
-- Name: sales_order_items sales_order_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sales_order_items
    ADD CONSTRAINT sales_order_items_pkey PRIMARY KEY (id);


--
-- Name: sales_orders sales_orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sales_orders
    ADD CONSTRAINT sales_orders_pkey PRIMARY KEY (id);


--
-- Name: schedule_plans schedule_plans_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedule_plans
    ADD CONSTRAINT schedule_plans_pkey PRIMARY KEY (id);


--
-- Name: schedule_results schedule_results_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schedule_results
    ADD CONSTRAINT schedule_results_pkey PRIMARY KEY (id);


--
-- Name: scp_customer_inquiry scp_customer_inquiry_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_customer_inquiry
    ADD CONSTRAINT scp_customer_inquiry_pkey PRIMARY KEY (id);


--
-- Name: scp_inquiry_item scp_inquiry_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_inquiry_item
    ADD CONSTRAINT scp_inquiry_item_pkey PRIMARY KEY (id);


--
-- Name: scp_mrs_item scp_mrs_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_mrs_item
    ADD CONSTRAINT scp_mrs_item_pkey PRIMARY KEY (id);


--
-- Name: scp_mrs scp_mrs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_mrs
    ADD CONSTRAINT scp_mrs_pkey PRIMARY KEY (id);


--
-- Name: scp_po_change_log scp_po_change_log_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_po_change_log
    ADD CONSTRAINT scp_po_change_log_pkey PRIMARY KEY (id);


--
-- Name: scp_purchase_order_item scp_purchase_order_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_purchase_order_item
    ADD CONSTRAINT scp_purchase_order_item_pkey PRIMARY KEY (id);


--
-- Name: scp_purchase_order scp_purchase_order_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_purchase_order
    ADD CONSTRAINT scp_purchase_order_pkey PRIMARY KEY (id);


--
-- Name: scp_purchase_plan_item scp_purchase_plan_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_purchase_plan_item
    ADD CONSTRAINT scp_purchase_plan_item_pkey PRIMARY KEY (id);


--
-- Name: scp_purchase_plan scp_purchase_plan_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_purchase_plan
    ADD CONSTRAINT scp_purchase_plan_pkey PRIMARY KEY (id);


--
-- Name: scp_qad_sync_log scp_qad_sync_log_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_qad_sync_log
    ADD CONSTRAINT scp_qad_sync_log_pkey PRIMARY KEY (id);


--
-- Name: scp_quote_comparison scp_quote_comparison_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_quote_comparison
    ADD CONSTRAINT scp_quote_comparison_pkey PRIMARY KEY (id);


--
-- Name: scp_quote_item scp_quote_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_quote_item
    ADD CONSTRAINT scp_quote_item_pkey PRIMARY KEY (id);


--
-- Name: scp_rfq_invite scp_rfq_invite_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_rfq_invite
    ADD CONSTRAINT scp_rfq_invite_pkey PRIMARY KEY (id);


--
-- Name: scp_rfq_item scp_rfq_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_rfq_item
    ADD CONSTRAINT scp_rfq_item_pkey PRIMARY KEY (id);


--
-- Name: scp_rfq scp_rfq_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_rfq
    ADD CONSTRAINT scp_rfq_pkey PRIMARY KEY (id);


--
-- Name: scp_sales_order_item scp_sales_order_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_sales_order_item
    ADD CONSTRAINT scp_sales_order_item_pkey PRIMARY KEY (id);


--
-- Name: scp_sales_order scp_sales_order_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_sales_order
    ADD CONSTRAINT scp_sales_order_pkey PRIMARY KEY (id);


--
-- Name: scp_so_change_log scp_so_change_log_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_so_change_log
    ADD CONSTRAINT scp_so_change_log_pkey PRIMARY KEY (id);


--
-- Name: scp_supplier_bank scp_supplier_bank_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_supplier_bank
    ADD CONSTRAINT scp_supplier_bank_pkey PRIMARY KEY (id);


--
-- Name: scp_supplier_contact scp_supplier_contact_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_supplier_contact
    ADD CONSTRAINT scp_supplier_contact_pkey PRIMARY KEY (id);


--
-- Name: scp_supplier_delivery_record scp_supplier_delivery_record_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_supplier_delivery_record
    ADD CONSTRAINT scp_supplier_delivery_record_pkey PRIMARY KEY (id);


--
-- Name: scp_supplier_grade_standard scp_supplier_grade_standard_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_supplier_grade_standard
    ADD CONSTRAINT scp_supplier_grade_standard_pkey PRIMARY KEY (id);


--
-- Name: scp_supplier_kpi scp_supplier_kpi_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_supplier_kpi
    ADD CONSTRAINT scp_supplier_kpi_pkey PRIMARY KEY (id);


--
-- Name: scp_supplier_purchase_info scp_supplier_purchase_info_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_supplier_purchase_info
    ADD CONSTRAINT scp_supplier_purchase_info_pkey PRIMARY KEY (id);


--
-- Name: scp_supplier_quality_record scp_supplier_quality_record_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_supplier_quality_record
    ADD CONSTRAINT scp_supplier_quality_record_pkey PRIMARY KEY (id);


--
-- Name: scp_supplier_quote scp_supplier_quote_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_supplier_quote
    ADD CONSTRAINT scp_supplier_quote_pkey PRIMARY KEY (id);


--
-- Name: serial_numbers serial_numbers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.serial_numbers
    ADD CONSTRAINT serial_numbers_pkey PRIMARY KEY (id);


--
-- Name: shifts shifts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.shifts
    ADD CONSTRAINT shifts_pkey PRIMARY KEY (id);


--
-- Name: spc_data spc_data_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.spc_data
    ADD CONSTRAINT spc_data_pkey PRIMARY KEY (id);


--
-- Name: stock_checks stock_checks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.stock_checks
    ADD CONSTRAINT stock_checks_pkey PRIMARY KEY (id);


--
-- Name: supplier_asn_items supplier_asn_items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.supplier_asn_items
    ADD CONSTRAINT supplier_asn_items_pkey PRIMARY KEY (id);


--
-- Name: supplier_asns supplier_asns_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.supplier_asns
    ADD CONSTRAINT supplier_asns_pkey PRIMARY KEY (id);


--
-- Name: sys_dept sys_dept_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_dept
    ADD CONSTRAINT sys_dept_pkey PRIMARY KEY (id);


--
-- Name: sys_dict_data sys_dict_data_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_dict_data
    ADD CONSTRAINT sys_dict_data_pkey PRIMARY KEY (id);


--
-- Name: sys_dict_type sys_dict_type_dict_type_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_dict_type
    ADD CONSTRAINT sys_dict_type_dict_type_key UNIQUE (dict_type);


--
-- Name: sys_dict_type sys_dict_type_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_dict_type
    ADD CONSTRAINT sys_dict_type_pkey PRIMARY KEY (id);


--
-- Name: sys_interface_config sys_interface_config_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_interface_config
    ADD CONSTRAINT sys_interface_config_pkey PRIMARY KEY (id);


--
-- Name: sys_interface_exec_log sys_interface_exec_log_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_interface_exec_log
    ADD CONSTRAINT sys_interface_exec_log_pkey PRIMARY KEY (id);


--
-- Name: sys_interface_field_map sys_interface_field_map_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_interface_field_map
    ADD CONSTRAINT sys_interface_field_map_pkey PRIMARY KEY (id);


--
-- Name: sys_interface_trigger sys_interface_trigger_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_interface_trigger
    ADD CONSTRAINT sys_interface_trigger_pkey PRIMARY KEY (id);


--
-- Name: sys_login_log sys_login_log_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_login_log
    ADD CONSTRAINT sys_login_log_pkey PRIMARY KEY (id);


--
-- Name: sys_menu sys_menu_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_menu
    ADD CONSTRAINT sys_menu_pkey PRIMARY KEY (id);


--
-- Name: sys_notice sys_notice_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_notice
    ADD CONSTRAINT sys_notice_pkey PRIMARY KEY (id);


--
-- Name: sys_notice_read_record sys_notice_read_record_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_notice_read_record
    ADD CONSTRAINT sys_notice_read_record_pkey PRIMARY KEY (id);


--
-- Name: sys_oper_log sys_oper_log_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_oper_log
    ADD CONSTRAINT sys_oper_log_pkey PRIMARY KEY (id);


--
-- Name: sys_post sys_post_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_post
    ADD CONSTRAINT sys_post_pkey PRIMARY KEY (id);


--
-- Name: sys_print_template sys_print_template_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_print_template
    ADD CONSTRAINT sys_print_template_pkey PRIMARY KEY (id);


--
-- Name: sys_role_menu sys_role_menu_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_role_menu
    ADD CONSTRAINT sys_role_menu_pkey PRIMARY KEY (role_id, menu_id);


--
-- Name: sys_role_perm sys_role_perm_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_role_perm
    ADD CONSTRAINT sys_role_perm_pkey PRIMARY KEY (role_id, perm);


--
-- Name: sys_role sys_role_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_role
    ADD CONSTRAINT sys_role_pkey PRIMARY KEY (id);


--
-- Name: sys_tenant sys_tenant_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_tenant
    ADD CONSTRAINT sys_tenant_pkey PRIMARY KEY (id);


--
-- Name: sys_tenant sys_tenant_tenant_key_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_tenant
    ADD CONSTRAINT sys_tenant_tenant_key_key UNIQUE (tenant_key);


--
-- Name: sys_user sys_user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_user
    ADD CONSTRAINT sys_user_pkey PRIMARY KEY (id);


--
-- Name: sys_user_role sys_user_role_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_user_role
    ADD CONSTRAINT sys_user_role_pkey PRIMARY KEY (user_id, role_id);


--
-- Name: tra_andon_call tra_andon_call_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tra_andon_call
    ADD CONSTRAINT tra_andon_call_pkey PRIMARY KEY (id);


--
-- Name: tra_andon_call tra_andon_call_tenant_id_call_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tra_andon_call
    ADD CONSTRAINT tra_andon_call_tenant_id_call_no_key UNIQUE (tenant_id, call_no);


--
-- Name: tra_data_collection tra_data_collection_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tra_data_collection
    ADD CONSTRAINT tra_data_collection_pkey PRIMARY KEY (id);


--
-- Name: tra_serial_number tra_serial_number_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tra_serial_number
    ADD CONSTRAINT tra_serial_number_pkey PRIMARY KEY (id);


--
-- Name: tra_serial_number tra_serial_number_tenant_id_serial_number_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tra_serial_number
    ADD CONSTRAINT tra_serial_number_tenant_id_serial_number_key UNIQUE (tenant_id, serial_number);


--
-- Name: tra_trace_record tra_trace_record_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tra_trace_record
    ADD CONSTRAINT tra_trace_record_pkey PRIMARY KEY (id);


--
-- Name: tra_trace_record tra_trace_record_tenant_id_trace_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tra_trace_record
    ADD CONSTRAINT tra_trace_record_tenant_id_trace_no_key UNIQUE (tenant_id, trace_no);


--
-- Name: trace_records trace_records_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.trace_records
    ADD CONSTRAINT trace_records_pkey PRIMARY KEY (id);


--
-- Name: transfer_orders transfer_orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transfer_orders
    ADD CONSTRAINT transfer_orders_pkey PRIMARY KEY (id);


--
-- Name: visual_inspection_result visual_inspection_result_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.visual_inspection_result
    ADD CONSTRAINT visual_inspection_result_pkey PRIMARY KEY (id);


--
-- Name: visual_inspection_task visual_inspection_task_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.visual_inspection_task
    ADD CONSTRAINT visual_inspection_task_pkey PRIMARY KEY (id);


--
-- Name: warehouses warehouses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.warehouses
    ADD CONSTRAINT warehouses_pkey PRIMARY KEY (id);


--
-- Name: wms_agv_device wms_agv_device_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_agv_device
    ADD CONSTRAINT wms_agv_device_pkey PRIMARY KEY (id);


--
-- Name: wms_agv_location_mapping wms_agv_location_mapping_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_agv_location_mapping
    ADD CONSTRAINT wms_agv_location_mapping_pkey PRIMARY KEY (id);


--
-- Name: wms_agv_task wms_agv_task_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_agv_task
    ADD CONSTRAINT wms_agv_task_pkey PRIMARY KEY (id);


--
-- Name: wms_delivery_order_item wms_delivery_order_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_delivery_order_item
    ADD CONSTRAINT wms_delivery_order_item_pkey PRIMARY KEY (id);


--
-- Name: wms_delivery_order wms_delivery_order_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_delivery_order
    ADD CONSTRAINT wms_delivery_order_pkey PRIMARY KEY (id);


--
-- Name: wms_delivery_order wms_delivery_order_tenant_id_delivery_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_delivery_order
    ADD CONSTRAINT wms_delivery_order_tenant_id_delivery_no_key UNIQUE (tenant_id, delivery_no);


--
-- Name: wms_inventory wms_inventory_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_inventory
    ADD CONSTRAINT wms_inventory_pkey PRIMARY KEY (id);


--
-- Name: wms_inventory_record wms_inventory_record_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_inventory_record
    ADD CONSTRAINT wms_inventory_record_pkey PRIMARY KEY (id);


--
-- Name: wms_inventory_record wms_inventory_record_tenant_id_record_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_inventory_record
    ADD CONSTRAINT wms_inventory_record_tenant_id_record_no_key UNIQUE (tenant_id, record_no);


--
-- Name: wms_kanban_pull wms_kanban_pull_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_kanban_pull
    ADD CONSTRAINT wms_kanban_pull_pkey PRIMARY KEY (id);


--
-- Name: wms_location wms_location_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_location
    ADD CONSTRAINT wms_location_pkey PRIMARY KEY (id);


--
-- Name: wms_location wms_location_tenant_id_location_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_location
    ADD CONSTRAINT wms_location_tenant_id_location_code_key UNIQUE (tenant_id, location_code);


--
-- Name: wms_production_complete_item wms_production_complete_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_complete_item
    ADD CONSTRAINT wms_production_complete_item_pkey PRIMARY KEY (id);


--
-- Name: wms_production_complete wms_production_complete_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_complete
    ADD CONSTRAINT wms_production_complete_pkey PRIMARY KEY (id);


--
-- Name: wms_production_issue_item wms_production_issue_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_issue_item
    ADD CONSTRAINT wms_production_issue_item_pkey PRIMARY KEY (id);


--
-- Name: wms_production_issue wms_production_issue_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_issue
    ADD CONSTRAINT wms_production_issue_pkey PRIMARY KEY (id);


--
-- Name: wms_production_return_item wms_production_return_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_return_item
    ADD CONSTRAINT wms_production_return_item_pkey PRIMARY KEY (id);


--
-- Name: wms_production_return wms_production_return_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_return
    ADD CONSTRAINT wms_production_return_pkey PRIMARY KEY (id);


--
-- Name: wms_production_stock_in_item wms_production_stock_in_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_stock_in_item
    ADD CONSTRAINT wms_production_stock_in_item_pkey PRIMARY KEY (id);


--
-- Name: wms_production_stock_in wms_production_stock_in_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_stock_in
    ADD CONSTRAINT wms_production_stock_in_pkey PRIMARY KEY (id);


--
-- Name: wms_purchase_return_item wms_purchase_return_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_purchase_return_item
    ADD CONSTRAINT wms_purchase_return_item_pkey PRIMARY KEY (id);


--
-- Name: wms_purchase_return wms_purchase_return_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_purchase_return
    ADD CONSTRAINT wms_purchase_return_pkey PRIMARY KEY (id);


--
-- Name: wms_receive_order_item wms_receive_order_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_receive_order_item
    ADD CONSTRAINT wms_receive_order_item_pkey PRIMARY KEY (id);


--
-- Name: wms_receive_order wms_receive_order_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_receive_order
    ADD CONSTRAINT wms_receive_order_pkey PRIMARY KEY (id);


--
-- Name: wms_receive_order wms_receive_order_tenant_id_receive_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_receive_order
    ADD CONSTRAINT wms_receive_order_tenant_id_receive_no_key UNIQUE (tenant_id, receive_no);


--
-- Name: wms_sales_return_item wms_sales_return_item_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_sales_return_item
    ADD CONSTRAINT wms_sales_return_item_pkey PRIMARY KEY (id);


--
-- Name: wms_sales_return wms_sales_return_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_sales_return
    ADD CONSTRAINT wms_sales_return_pkey PRIMARY KEY (id);


--
-- Name: wms_side_location wms_side_location_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_side_location
    ADD CONSTRAINT wms_side_location_pkey PRIMARY KEY (id);


--
-- Name: wms_stock_check wms_stock_check_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_stock_check
    ADD CONSTRAINT wms_stock_check_pkey PRIMARY KEY (id);


--
-- Name: wms_stock_check wms_stock_check_tenant_id_check_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_stock_check
    ADD CONSTRAINT wms_stock_check_tenant_id_check_no_key UNIQUE (tenant_id, check_no);


--
-- Name: wms_transfer_order wms_transfer_order_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_transfer_order
    ADD CONSTRAINT wms_transfer_order_pkey PRIMARY KEY (id);


--
-- Name: wms_transfer_order wms_transfer_order_tenant_id_transfer_no_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_transfer_order
    ADD CONSTRAINT wms_transfer_order_tenant_id_transfer_no_key UNIQUE (tenant_id, transfer_no);


--
-- Name: wms_warehouse wms_warehouse_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_warehouse
    ADD CONSTRAINT wms_warehouse_pkey PRIMARY KEY (id);


--
-- Name: wms_warehouse wms_warehouse_tenant_id_warehouse_code_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_warehouse
    ADD CONSTRAINT wms_warehouse_tenant_id_warehouse_code_key UNIQUE (tenant_id, warehouse_code);


--
-- Name: work_centers work_centers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.work_centers
    ADD CONSTRAINT work_centers_pkey PRIMARY KEY (id);


--
-- Name: workshops workshops_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workshops
    ADD CONSTRAINT workshops_pkey PRIMARY KEY (id);


--
-- Name: workstations workstations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workstations
    ADD CONSTRAINT workstations_pkey PRIMARY KEY (id);


--
-- Name: idx_ai_chat_conversation_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_ai_chat_conversation_deleted_at ON public.ai_chat_conversation USING btree (deleted_at);


--
-- Name: idx_ai_chat_conversation_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_ai_chat_conversation_tenant_id ON public.ai_chat_conversation USING btree (tenant_id);


--
-- Name: idx_ai_chat_conversation_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_ai_chat_conversation_user_id ON public.ai_chat_conversation USING btree (user_id);


--
-- Name: idx_ai_chat_message_conversation_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_ai_chat_message_conversation_id ON public.ai_chat_message USING btree (conversation_id);


--
-- Name: idx_ai_chat_message_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_ai_chat_message_tenant_id ON public.ai_chat_message USING btree (tenant_id);


--
-- Name: idx_ai_chat_message_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_ai_chat_message_user_id ON public.ai_chat_message USING btree (user_id);


--
-- Name: idx_ai_config_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_ai_config_tenant_id ON public.ai_config USING btree (tenant_id);


--
-- Name: idx_alert_escalation_rule_rule_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_alert_escalation_rule_rule_code ON public.alert_escalation_rule USING btree (rule_code);


--
-- Name: idx_alert_escalation_rule_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_alert_escalation_rule_tenant_id ON public.alert_escalation_rule USING btree (tenant_id);


--
-- Name: idx_alert_notification_log_alert_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_alert_notification_log_alert_id ON public.alert_notification_log USING btree (alert_id);


--
-- Name: idx_alert_notification_log_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_alert_notification_log_tenant_id ON public.alert_notification_log USING btree (tenant_id);


--
-- Name: idx_alert_record_alert_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_alert_record_alert_no ON public.alert_record USING btree (alert_no);


--
-- Name: idx_alert_record_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_alert_record_tenant_id ON public.alert_record USING btree (tenant_id);


--
-- Name: idx_alert_rule_rule_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_alert_rule_rule_code ON public.alert_rule USING btree (rule_code);


--
-- Name: idx_alert_rule_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_alert_rule_tenant_id ON public.alert_rule USING btree (tenant_id);


--
-- Name: idx_andon_call_andon_type; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_call_andon_type ON public.andon_call USING btree (andon_type);


--
-- Name: idx_andon_call_call_time; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_call_call_time ON public.andon_call USING btree (call_time);


--
-- Name: idx_andon_call_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_call_deleted_at ON public.andon_call USING btree (deleted_at);


--
-- Name: idx_andon_call_production_line_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_call_production_line_id ON public.andon_call USING btree (production_line_id);


--
-- Name: idx_andon_call_related_ncr_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_call_related_ncr_id ON public.andon_call USING btree (related_ncr_id);


--
-- Name: idx_andon_call_related_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_call_related_order_id ON public.andon_call USING btree (related_order_id);


--
-- Name: idx_andon_call_related_repair_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_call_related_repair_id ON public.andon_call USING btree (related_repair_id);


--
-- Name: idx_andon_call_status; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_call_status ON public.andon_call USING btree (status);


--
-- Name: idx_andon_call_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_call_tenant_id ON public.andon_call USING btree (tenant_id);


--
-- Name: idx_andon_call_workshop_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_call_workshop_id ON public.andon_call USING btree (workshop_id);


--
-- Name: idx_andon_call_workstation_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_call_workstation_id ON public.andon_call USING btree (workstation_id);


--
-- Name: idx_andon_calls_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_calls_deleted_at ON public.andon_calls USING btree (deleted_at);


--
-- Name: idx_andon_calls_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_calls_tenant_id ON public.andon_calls USING btree (tenant_id);


--
-- Name: idx_andon_escalation_log_call_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_escalation_log_call_id ON public.andon_escalation_log USING btree (call_id);


--
-- Name: idx_andon_escalation_log_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_escalation_log_deleted_at ON public.andon_escalation_log USING btree (deleted_at);


--
-- Name: idx_andon_escalation_rule_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_escalation_rule_deleted_at ON public.andon_escalation_rule USING btree (deleted_at);


--
-- Name: idx_andon_escalation_rule_rule_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_andon_escalation_rule_rule_code ON public.andon_escalation_rule USING btree (rule_code);


--
-- Name: idx_andon_escalation_rule_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_escalation_rule_tenant_id ON public.andon_escalation_rule USING btree (tenant_id);


--
-- Name: idx_andon_escalation_rule_workshop_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_escalation_rule_workshop_id ON public.andon_escalation_rule USING btree (workshop_id);


--
-- Name: idx_andon_notification_log_call_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_notification_log_call_id ON public.andon_notification_log USING btree (call_id);


--
-- Name: idx_andon_notification_log_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_notification_log_deleted_at ON public.andon_notification_log USING btree (deleted_at);


--
-- Name: idx_andon_notification_log_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_notification_log_tenant_id ON public.andon_notification_log USING btree (tenant_id);


--
-- Name: idx_andon_report_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_report_deleted_at ON public.andon_report USING btree (deleted_at);


--
-- Name: idx_andon_report_line_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_report_line_id ON public.andon_report USING btree (line_id);


--
-- Name: idx_andon_report_report_date; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_report_report_date ON public.andon_report USING btree (report_date);


--
-- Name: idx_andon_report_station_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_report_station_id ON public.andon_report USING btree (station_id);


--
-- Name: idx_andon_report_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_report_tenant_id ON public.andon_report USING btree (tenant_id);


--
-- Name: idx_andon_report_workshop_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_andon_report_workshop_id ON public.andon_report USING btree (workshop_id);


--
-- Name: idx_aps_capacity_analysis_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_capacity_analysis_deleted_at ON public.aps_capacity_analysis USING btree (deleted_at);


--
-- Name: idx_aps_capacity_analysis_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_capacity_analysis_tenant_id ON public.aps_capacity_analysis USING btree (tenant_id);


--
-- Name: idx_aps_changeover_matrix_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_changeover_matrix_deleted_at ON public.aps_changeover_matrix USING btree (deleted_at);


--
-- Name: idx_aps_changeover_matrix_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_changeover_matrix_tenant_id ON public.aps_changeover_matrix USING btree (tenant_id);


--
-- Name: idx_aps_delivery_rate_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_delivery_rate_deleted_at ON public.aps_delivery_rate USING btree (deleted_at);


--
-- Name: idx_aps_delivery_rate_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_delivery_rate_tenant_id ON public.aps_delivery_rate USING btree (tenant_id);


--
-- Name: idx_aps_jit_demand_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_jit_demand_deleted_at ON public.aps_jit_demand USING btree (deleted_at);


--
-- Name: idx_aps_jit_demand_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_jit_demand_tenant_id ON public.aps_jit_demand USING btree (tenant_id);


--
-- Name: idx_aps_mps_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_mps_deleted_at ON public.aps_mps USING btree (deleted_at);


--
-- Name: idx_aps_mps_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_mps_tenant ON public.aps_mps USING btree (tenant_id);


--
-- Name: idx_aps_mps_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_mps_tenant_id ON public.aps_mps USING btree (tenant_id);


--
-- Name: idx_aps_mrp_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_mrp_deleted_at ON public.aps_mrp USING btree (deleted_at);


--
-- Name: idx_aps_mrp_item_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_mrp_item_deleted_at ON public.aps_mrp_item USING btree (deleted_at);


--
-- Name: idx_aps_mrp_item_mrp; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_mrp_item_mrp ON public.aps_mrp_item USING btree (mrp_id);


--
-- Name: idx_aps_mrp_item_mrp_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_mrp_item_mrp_id ON public.aps_mrp_item USING btree (mrp_id);


--
-- Name: idx_aps_mrp_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_mrp_tenant ON public.aps_mrp USING btree (tenant_id);


--
-- Name: idx_aps_mrp_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_mrp_tenant_id ON public.aps_mrp USING btree (tenant_id);


--
-- Name: idx_aps_resource_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_resource_deleted_at ON public.aps_resource USING btree (deleted_at);


--
-- Name: idx_aps_resource_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_resource_tenant ON public.aps_resource USING btree (tenant_id);


--
-- Name: idx_aps_resource_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_resource_tenant_id ON public.aps_resource USING btree (tenant_id);


--
-- Name: idx_aps_rolling_schedule_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_rolling_schedule_deleted_at ON public.aps_rolling_schedule USING btree (deleted_at);


--
-- Name: idx_aps_rolling_schedule_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_rolling_schedule_tenant_id ON public.aps_rolling_schedule USING btree (tenant_id);


--
-- Name: idx_aps_schedule_plan_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_schedule_plan_deleted_at ON public.aps_schedule_plan USING btree (deleted_at);


--
-- Name: idx_aps_schedule_plan_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_schedule_plan_tenant ON public.aps_schedule_plan USING btree (tenant_id);


--
-- Name: idx_aps_schedule_plan_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_schedule_plan_tenant_id ON public.aps_schedule_plan USING btree (tenant_id);


--
-- Name: idx_aps_schedule_result_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_schedule_result_deleted_at ON public.aps_schedule_result USING btree (deleted_at);


--
-- Name: idx_aps_schedule_result_order; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_schedule_result_order ON public.aps_schedule_result USING btree (order_id);


--
-- Name: idx_aps_schedule_result_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_schedule_result_order_id ON public.aps_schedule_result USING btree (order_id);


--
-- Name: idx_aps_schedule_result_plan; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_schedule_result_plan ON public.aps_schedule_result USING btree (plan_id);


--
-- Name: idx_aps_schedule_result_plan_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_schedule_result_plan_id ON public.aps_schedule_result USING btree (plan_id);


--
-- Name: idx_aps_work_center_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_work_center_deleted_at ON public.aps_work_center USING btree (deleted_at);


--
-- Name: idx_aps_work_center_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_work_center_tenant ON public.aps_work_center USING btree (tenant_id);


--
-- Name: idx_aps_work_center_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_work_center_tenant_id ON public.aps_work_center USING btree (tenant_id);


--
-- Name: idx_aps_working_calendar_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_working_calendar_deleted_at ON public.aps_working_calendar USING btree (deleted_at);


--
-- Name: idx_aps_working_calendar_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_working_calendar_tenant_id ON public.aps_working_calendar USING btree (tenant_id);


--
-- Name: idx_aps_working_calendar_workshop_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_aps_working_calendar_workshop_id ON public.aps_working_calendar USING btree (workshop_id);


--
-- Name: idx_bom_items_bom_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bom_items_bom_id ON public.bom_items USING btree (bom_id);


--
-- Name: idx_bom_items_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bom_items_deleted_at ON public.bom_items USING btree (deleted_at);


--
-- Name: idx_boms_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_boms_deleted_at ON public.boms USING btree (deleted_at);


--
-- Name: idx_boms_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_boms_tenant_id ON public.boms USING btree (tenant_id);


--
-- Name: idx_bpm_approval_record_instance_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_approval_record_instance_id ON public.bpm_approval_record USING btree (instance_id);


--
-- Name: idx_bpm_approval_record_task_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_approval_record_task_id ON public.bpm_approval_record USING btree (task_id);


--
-- Name: idx_bpm_approval_record_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_approval_record_tenant_id ON public.bpm_approval_record USING btree (tenant_id);


--
-- Name: idx_bpm_biz_mapping_biz_type; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_bpm_biz_mapping_biz_type ON public.bpm_biz_mapping USING btree (biz_type);


--
-- Name: idx_bpm_biz_mapping_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_biz_mapping_tenant_id ON public.bpm_biz_mapping USING btree (tenant_id);


--
-- Name: idx_bpm_delegate_record_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_delegate_record_tenant_id ON public.bpm_delegate_record USING btree (tenant_id);


--
-- Name: idx_bpm_form_definition_form_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_bpm_form_definition_form_code ON public.bpm_form_definition USING btree (form_code);


--
-- Name: idx_bpm_form_definition_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_form_definition_tenant_id ON public.bpm_form_definition USING btree (tenant_id);


--
-- Name: idx_bpm_form_field_form_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_form_field_form_id ON public.bpm_form_field USING btree (form_id);


--
-- Name: idx_bpm_form_field_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_form_field_tenant_id ON public.bpm_form_field USING btree (tenant_id);


--
-- Name: idx_bpm_form_version_form_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_form_version_form_id ON public.bpm_form_version USING btree (form_id);


--
-- Name: idx_bpm_form_version_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_form_version_tenant_id ON public.bpm_form_version USING btree (tenant_id);


--
-- Name: idx_bpm_model_version_model_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_model_version_model_id ON public.bpm_model_version USING btree (model_id);


--
-- Name: idx_bpm_model_version_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_model_version_tenant_id ON public.bpm_model_version USING btree (tenant_id);


--
-- Name: idx_bpm_node_definition_model_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_node_definition_model_id ON public.bpm_node_definition USING btree (model_id);


--
-- Name: idx_bpm_node_definition_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_node_definition_tenant_id ON public.bpm_node_definition USING btree (tenant_id);


--
-- Name: idx_bpm_node_timeout_model_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_node_timeout_model_id ON public.bpm_node_timeout USING btree (model_id);


--
-- Name: idx_bpm_node_timeout_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_node_timeout_tenant_id ON public.bpm_node_timeout USING btree (tenant_id);


--
-- Name: idx_bpm_order_day_workorder_map_day_plan_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_order_day_workorder_map_day_plan_id ON public.bpm_order_day_workorder_map USING btree (day_plan_id);


--
-- Name: idx_bpm_order_day_workorder_map_day_plan_item_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_order_day_workorder_map_day_plan_item_id ON public.bpm_order_day_workorder_map USING btree (day_plan_item_id);


--
-- Name: idx_bpm_order_day_workorder_map_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_order_day_workorder_map_deleted_at ON public.bpm_order_day_workorder_map USING btree (deleted_at);


--
-- Name: idx_bpm_order_day_workorder_map_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_order_day_workorder_map_tenant_id ON public.bpm_order_day_workorder_map USING btree (tenant_id);


--
-- Name: idx_bpm_order_day_workorder_map_work_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_order_day_workorder_map_work_order_id ON public.bpm_order_day_workorder_map USING btree (work_order_id);


--
-- Name: idx_bpm_process_instance_instance_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_bpm_process_instance_instance_no ON public.bpm_process_instance USING btree (instance_no);


--
-- Name: idx_bpm_process_instance_model_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_process_instance_model_id ON public.bpm_process_instance USING btree (model_id);


--
-- Name: idx_bpm_process_instance_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_process_instance_tenant_id ON public.bpm_process_instance USING btree (tenant_id);


--
-- Name: idx_bpm_process_model_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_process_model_tenant_id ON public.bpm_process_model USING btree (tenant_id);


--
-- Name: idx_bpm_reminder_record_instance_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_reminder_record_instance_id ON public.bpm_reminder_record USING btree (instance_id);


--
-- Name: idx_bpm_reminder_record_task_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_reminder_record_task_id ON public.bpm_reminder_record USING btree (task_id);


--
-- Name: idx_bpm_reminder_record_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_reminder_record_tenant_id ON public.bpm_reminder_record USING btree (tenant_id);


--
-- Name: idx_bpm_sequence_flow_model_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_sequence_flow_model_id ON public.bpm_sequence_flow USING btree (model_id);


--
-- Name: idx_bpm_sequence_flow_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_sequence_flow_tenant_id ON public.bpm_sequence_flow USING btree (tenant_id);


--
-- Name: idx_bpm_task_assignment_model_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_task_assignment_model_id ON public.bpm_task_assignment USING btree (model_id);


--
-- Name: idx_bpm_task_assignment_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_task_assignment_tenant_id ON public.bpm_task_assignment USING btree (tenant_id);


--
-- Name: idx_bpm_task_instance_instance_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_bpm_task_instance_instance_id ON public.bpm_task_instance USING btree (instance_id);


--
-- Name: idx_bpm_task_instance_task_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_bpm_task_instance_task_no ON public.bpm_task_instance USING btree (task_no);


--
-- Name: idx_container_master_barcode; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_container_master_barcode ON public.container_master USING btree (barcode);


--
-- Name: idx_container_master_container_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_container_master_container_code ON public.container_master USING btree (container_code);


--
-- Name: idx_container_master_customer_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_container_master_customer_id ON public.container_master USING btree (customer_id);


--
-- Name: idx_container_master_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_container_master_tenant_id ON public.container_master USING btree (tenant_id);


--
-- Name: idx_container_movement_container_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_container_movement_container_id ON public.container_movement USING btree (container_id);


--
-- Name: idx_container_movement_operator_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_container_movement_operator_id ON public.container_movement USING btree (operator_id);


--
-- Name: idx_container_movement_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_container_movement_tenant_id ON public.container_movement USING btree (tenant_id);


--
-- Name: idx_data_collections_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_data_collections_deleted_at ON public.data_collections USING btree (deleted_at);


--
-- Name: idx_data_collections_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_data_collections_tenant_id ON public.data_collections USING btree (tenant_id);


--
-- Name: idx_dc_collect_record_collect_time; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_dc_collect_record_collect_time ON public.dc_collect_record USING btree (collect_time);


--
-- Name: idx_dc_collect_record_point_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_dc_collect_record_point_id ON public.dc_collect_record USING btree (point_id);


--
-- Name: idx_dc_collect_record_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_dc_collect_record_tenant_id ON public.dc_collect_record USING btree (tenant_id);


--
-- Name: idx_dc_data_point_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_dc_data_point_tenant_id ON public.dc_data_point USING btree (tenant_id);


--
-- Name: idx_dc_scan_log_related_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_dc_scan_log_related_id ON public.dc_scan_log USING btree (related_id);


--
-- Name: idx_dc_scan_log_scan_time; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_dc_scan_log_scan_time ON public.dc_scan_log USING btree (scan_time);


--
-- Name: idx_dc_scan_log_scan_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_dc_scan_log_scan_user_id ON public.dc_scan_log USING btree (scan_user_id);


--
-- Name: idx_dc_scan_log_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_dc_scan_log_tenant_id ON public.dc_scan_log USING btree (tenant_id);


--
-- Name: idx_dc_scan_log_workshop_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_dc_scan_log_workshop_id ON public.dc_scan_log USING btree (workshop_id);


--
-- Name: idx_dc_scan_log_workstation_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_dc_scan_log_workstation_id ON public.dc_scan_log USING btree (workstation_id);


--
-- Name: idx_defect_codes_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_defect_codes_deleted_at ON public.defect_codes USING btree (deleted_at);


--
-- Name: idx_defect_codes_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_defect_codes_tenant_id ON public.defect_codes USING btree (tenant_id);


--
-- Name: idx_defect_records_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_defect_records_deleted_at ON public.defect_records USING btree (deleted_at);


--
-- Name: idx_defect_records_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_defect_records_tenant_id ON public.defect_records USING btree (tenant_id);


--
-- Name: idx_delivery_order_items_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_delivery_order_items_deleted_at ON public.delivery_order_items USING btree (deleted_at);


--
-- Name: idx_delivery_order_items_delivery_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_delivery_order_items_delivery_id ON public.delivery_order_items USING btree (delivery_id);


--
-- Name: idx_delivery_orders_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_delivery_orders_deleted_at ON public.delivery_orders USING btree (deleted_at);


--
-- Name: idx_delivery_orders_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_delivery_orders_tenant_id ON public.delivery_orders USING btree (tenant_id);


--
-- Name: idx_delivery_report_customer_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_delivery_report_customer_id ON public.delivery_report USING btree (customer_id);


--
-- Name: idx_delivery_report_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_delivery_report_deleted_at ON public.delivery_report USING btree (deleted_at);


--
-- Name: idx_delivery_report_report_month; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_delivery_report_report_month ON public.delivery_report USING btree (report_month);


--
-- Name: idx_delivery_report_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_delivery_report_tenant_id ON public.delivery_report USING btree (tenant_id);


--
-- Name: idx_dispatches_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_dispatches_deleted_at ON public.dispatches USING btree (deleted_at);


--
-- Name: idx_dispatches_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_dispatches_order_id ON public.dispatches USING btree (order_id);


--
-- Name: idx_dispatches_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_dispatches_tenant_id ON public.dispatches USING btree (tenant_id);


--
-- Name: idx_eam_equipment_downtime_downtime_type; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_eam_equipment_downtime_downtime_type ON public.eam_equipment_downtime USING btree (downtime_type);


--
-- Name: idx_eam_equipment_downtime_equipment_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_eam_equipment_downtime_equipment_id ON public.eam_equipment_downtime USING btree (equipment_id);


--
-- Name: idx_eam_equipment_downtime_start_time; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_eam_equipment_downtime_start_time ON public.eam_equipment_downtime USING btree (start_time);


--
-- Name: idx_eam_equipment_downtime_status; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_eam_equipment_downtime_status ON public.eam_equipment_downtime USING btree (status);


--
-- Name: idx_eam_equipment_downtime_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_eam_equipment_downtime_tenant_id ON public.eam_equipment_downtime USING btree (tenant_id);


--
-- Name: idx_eam_equipment_downtime_work_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_eam_equipment_downtime_work_order_id ON public.eam_equipment_downtime USING btree (work_order_id);


--
-- Name: idx_eam_repair_flow_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_eam_repair_flow_code ON public.eam_repair_flow USING btree (flow_code);


--
-- Name: idx_eam_repair_flow_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_eam_repair_flow_deleted_at ON public.eam_repair_flow USING btree (deleted_at);


--
-- Name: idx_eam_repair_flow_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_eam_repair_flow_tenant_id ON public.eam_repair_flow USING btree (tenant_id);


--
-- Name: idx_eam_repair_job_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_eam_repair_job_code ON public.eam_repair_job USING btree (job_code);


--
-- Name: idx_eam_repair_job_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_eam_repair_job_deleted_at ON public.eam_repair_job USING btree (deleted_at);


--
-- Name: idx_eam_repair_job_equipment_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_eam_repair_job_equipment_id ON public.eam_repair_job USING btree (equipment_id);


--
-- Name: idx_eam_repair_job_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_eam_repair_job_tenant_id ON public.eam_repair_job USING btree (tenant_id);


--
-- Name: idx_eam_repair_std_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_eam_repair_std_code ON public.eam_repair_std USING btree (std_code);


--
-- Name: idx_eam_repair_std_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_eam_repair_std_deleted_at ON public.eam_repair_std USING btree (deleted_at);


--
-- Name: idx_eam_repair_std_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_eam_repair_std_tenant_id ON public.eam_repair_std USING btree (tenant_id);


--
-- Name: idx_ene_energy_record_date; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_ene_energy_record_date ON public.ene_energy_record USING btree (record_date);


--
-- Name: idx_ene_energy_record_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_ene_energy_record_deleted_at ON public.ene_energy_record USING btree (deleted_at);


--
-- Name: idx_ene_energy_record_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_ene_energy_record_tenant ON public.ene_energy_record USING btree (tenant_id);


--
-- Name: idx_ene_energy_record_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_ene_energy_record_tenant_id ON public.ene_energy_record USING btree (tenant_id);


--
-- Name: idx_energy_records_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_energy_records_deleted_at ON public.energy_records USING btree (deleted_at);


--
-- Name: idx_energy_records_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_energy_records_tenant_id ON public.energy_records USING btree (tenant_id);


--
-- Name: idx_equ_equipment_check_equipment; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_equipment_check_equipment ON public.equ_equipment_check USING btree (equipment_id);


--
-- Name: idx_equ_equipment_check_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_equipment_check_tenant ON public.equ_equipment_check USING btree (tenant_id);


--
-- Name: idx_equ_equipment_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_equipment_code ON public.equ_equipment USING btree (equipment_code);


--
-- Name: idx_equ_equipment_document_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_equipment_document_deleted_at ON public.equ_equipment_document USING btree (deleted_at);


--
-- Name: idx_equ_equipment_document_equipment_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_equipment_document_equipment_id ON public.equ_equipment_document USING btree (equipment_id);


--
-- Name: idx_equ_equipment_document_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_equipment_document_tenant_id ON public.equ_equipment_document USING btree (tenant_id);


--
-- Name: idx_equ_equipment_maintenance_equipment; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_equipment_maintenance_equipment ON public.equ_equipment_maintenance USING btree (equipment_id);


--
-- Name: idx_equ_equipment_maintenance_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_equipment_maintenance_tenant ON public.equ_equipment_maintenance USING btree (tenant_id);


--
-- Name: idx_equ_equipment_part_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_equipment_part_deleted_at ON public.equ_equipment_part USING btree (deleted_at);


--
-- Name: idx_equ_equipment_part_equipment_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_equipment_part_equipment_id ON public.equ_equipment_part USING btree (equipment_id);


--
-- Name: idx_equ_equipment_part_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_equipment_part_tenant_id ON public.equ_equipment_part USING btree (tenant_id);


--
-- Name: idx_equ_equipment_repair_equipment; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_equipment_repair_equipment ON public.equ_equipment_repair USING btree (equipment_id);


--
-- Name: idx_equ_equipment_repair_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_equipment_repair_tenant ON public.equ_equipment_repair USING btree (tenant_id);


--
-- Name: idx_equ_equipment_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_equipment_tenant ON public.equ_equipment USING btree (tenant_id);


--
-- Name: idx_equ_gauge_calibration_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_gauge_calibration_deleted_at ON public.equ_gauge_calibration USING btree (deleted_at);


--
-- Name: idx_equ_gauge_calibration_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_gauge_calibration_tenant_id ON public.equ_gauge_calibration USING btree (tenant_id);


--
-- Name: idx_equ_gauge_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_gauge_deleted_at ON public.equ_gauge USING btree (deleted_at);


--
-- Name: idx_equ_gauge_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_gauge_tenant_id ON public.equ_gauge USING btree (tenant_id);


--
-- Name: idx_equ_mold_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_mold_deleted_at ON public.equ_mold USING btree (deleted_at);


--
-- Name: idx_equ_mold_maintenance_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_mold_maintenance_deleted_at ON public.equ_mold_maintenance USING btree (deleted_at);


--
-- Name: idx_equ_mold_maintenance_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_mold_maintenance_tenant_id ON public.equ_mold_maintenance USING btree (tenant_id);


--
-- Name: idx_equ_mold_repair_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_mold_repair_deleted_at ON public.equ_mold_repair USING btree (deleted_at);


--
-- Name: idx_equ_mold_repair_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_mold_repair_tenant_id ON public.equ_mold_repair USING btree (tenant_id);


--
-- Name: idx_equ_mold_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_mold_tenant_id ON public.equ_mold USING btree (tenant_id);


--
-- Name: idx_equ_oee_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_oee_deleted_at ON public.equ_oee USING btree (deleted_at);


--
-- Name: idx_equ_oee_record_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_oee_record_deleted_at ON public.equ_oee_record USING btree (deleted_at);


--
-- Name: idx_equ_oee_record_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_oee_record_tenant_id ON public.equ_oee_record USING btree (tenant_id);


--
-- Name: idx_equ_oee_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_oee_tenant_id ON public.equ_oee USING btree (tenant_id);


--
-- Name: idx_equ_spare_part_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_spare_part_tenant ON public.equ_spare_part USING btree (tenant_id);


--
-- Name: idx_equ_teep_data_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_teep_data_deleted_at ON public.equ_teep_data USING btree (deleted_at);


--
-- Name: idx_equ_teep_data_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_equ_teep_data_tenant_id ON public.equ_teep_data USING btree (tenant_id);


--
-- Name: idx_fin_payment_request_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_payment_request_deleted_at ON public.fin_payment_request USING btree (deleted_at);


--
-- Name: idx_fin_payment_request_request_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_fin_payment_request_request_no ON public.fin_payment_request USING btree (request_no);


--
-- Name: idx_fin_payment_request_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_payment_request_tenant_id ON public.fin_payment_request USING btree (tenant_id);


--
-- Name: idx_fin_purchase_advance_advance_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_fin_purchase_advance_advance_no ON public.fin_purchase_advance USING btree (advance_no);


--
-- Name: idx_fin_purchase_advance_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_purchase_advance_deleted_at ON public.fin_purchase_advance USING btree (deleted_at);


--
-- Name: idx_fin_purchase_advance_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_purchase_advance_tenant_id ON public.fin_purchase_advance USING btree (tenant_id);


--
-- Name: idx_fin_purchase_settlement_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_purchase_settlement_deleted_at ON public.fin_purchase_settlement USING btree (deleted_at);


--
-- Name: idx_fin_purchase_settlement_item_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_purchase_settlement_item_deleted_at ON public.fin_purchase_settlement_item USING btree (deleted_at);


--
-- Name: idx_fin_purchase_settlement_item_settlement_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_purchase_settlement_item_settlement_id ON public.fin_purchase_settlement_item USING btree (settlement_id);


--
-- Name: idx_fin_purchase_settlement_item_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_purchase_settlement_item_tenant_id ON public.fin_purchase_settlement_item USING btree (tenant_id);


--
-- Name: idx_fin_purchase_settlement_settlement_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_fin_purchase_settlement_settlement_no ON public.fin_purchase_settlement USING btree (settlement_no);


--
-- Name: idx_fin_purchase_settlement_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_purchase_settlement_tenant_id ON public.fin_purchase_settlement USING btree (tenant_id);


--
-- Name: idx_fin_sales_receipt_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_sales_receipt_deleted_at ON public.fin_sales_receipt USING btree (deleted_at);


--
-- Name: idx_fin_sales_receipt_receipt_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_fin_sales_receipt_receipt_no ON public.fin_sales_receipt USING btree (receipt_no);


--
-- Name: idx_fin_sales_receipt_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_sales_receipt_tenant_id ON public.fin_sales_receipt USING btree (tenant_id);


--
-- Name: idx_fin_sales_settlement_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_sales_settlement_deleted_at ON public.fin_sales_settlement USING btree (deleted_at);


--
-- Name: idx_fin_sales_settlement_item_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_sales_settlement_item_deleted_at ON public.fin_sales_settlement_item USING btree (deleted_at);


--
-- Name: idx_fin_sales_settlement_item_settlement_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_sales_settlement_item_settlement_id ON public.fin_sales_settlement_item USING btree (settlement_id);


--
-- Name: idx_fin_sales_settlement_item_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_sales_settlement_item_tenant_id ON public.fin_sales_settlement_item USING btree (tenant_id);


--
-- Name: idx_fin_sales_settlement_settlement_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_fin_sales_settlement_settlement_no ON public.fin_sales_settlement USING btree (settlement_no);


--
-- Name: idx_fin_sales_settlement_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_sales_settlement_tenant_id ON public.fin_sales_settlement USING btree (tenant_id);


--
-- Name: idx_fin_statement_detail_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_statement_detail_deleted_at ON public.fin_statement_detail USING btree (deleted_at);


--
-- Name: idx_fin_statement_detail_statement_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_statement_detail_statement_id ON public.fin_statement_detail USING btree (statement_id);


--
-- Name: idx_fin_statement_detail_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_statement_detail_tenant_id ON public.fin_statement_detail USING btree (tenant_id);


--
-- Name: idx_fin_supplier_statement_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_supplier_statement_deleted_at ON public.fin_supplier_statement USING btree (deleted_at);


--
-- Name: idx_fin_supplier_statement_statement_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_fin_supplier_statement_statement_no ON public.fin_supplier_statement USING btree (statement_no);


--
-- Name: idx_fin_supplier_statement_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fin_supplier_statement_tenant_id ON public.fin_supplier_statement USING btree (tenant_id);


--
-- Name: idx_fqcs_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fqcs_deleted_at ON public.fqcs USING btree (deleted_at);


--
-- Name: idx_fqcs_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_fqcs_tenant_id ON public.fqcs USING btree (tenant_id);


--
-- Name: idx_inventories_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_inventories_deleted_at ON public.inventories USING btree (deleted_at);


--
-- Name: idx_inventories_location_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_inventories_location_id ON public.inventories USING btree (location_id);


--
-- Name: idx_inventories_material_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_inventories_material_id ON public.inventories USING btree (material_id);


--
-- Name: idx_inventories_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_inventories_tenant_id ON public.inventories USING btree (tenant_id);


--
-- Name: idx_inventories_warehouse_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_inventories_warehouse_id ON public.inventories USING btree (warehouse_id);


--
-- Name: idx_inventory_records_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_inventory_records_deleted_at ON public.inventory_records USING btree (deleted_at);


--
-- Name: idx_inventory_records_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_inventory_records_tenant_id ON public.inventory_records USING btree (tenant_id);


--
-- Name: idx_ip_qcs_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_ip_qcs_deleted_at ON public.ip_qcs USING btree (deleted_at);


--
-- Name: idx_ip_qcs_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_ip_qcs_tenant_id ON public.ip_qcs USING btree (tenant_id);


--
-- Name: idx_iqc_items_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_iqc_items_deleted_at ON public.iqc_items USING btree (deleted_at);


--
-- Name: idx_iqc_items_iqc_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_iqc_items_iqc_id ON public.iqc_items USING btree (iqc_id);


--
-- Name: idx_iqcs_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_iqcs_deleted_at ON public.iqcs USING btree (deleted_at);


--
-- Name: idx_iqcs_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_iqcs_tenant_id ON public.iqcs USING btree (tenant_id);


--
-- Name: idx_lab_calibration_instrument_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_lab_calibration_instrument_id ON public.lab_calibration USING btree (instrument_id);


--
-- Name: idx_lab_calibration_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_lab_calibration_tenant_id ON public.lab_calibration USING btree (tenant_id);


--
-- Name: idx_lab_instrument_calibration_status; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_lab_instrument_calibration_status ON public.lab_instrument USING btree (calibration_status);


--
-- Name: idx_lab_instrument_instrument_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_lab_instrument_instrument_code ON public.lab_instrument USING btree (instrument_code);


--
-- Name: idx_lab_instrument_instrument_type; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_lab_instrument_instrument_type ON public.lab_instrument USING btree (instrument_type);


--
-- Name: idx_lab_instrument_status; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_lab_instrument_status ON public.lab_instrument USING btree (status);


--
-- Name: idx_lab_instrument_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_lab_instrument_tenant_id ON public.lab_instrument USING btree (tenant_id);


--
-- Name: idx_lab_report_report_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_lab_report_report_no ON public.lab_report USING btree (report_no);


--
-- Name: idx_lab_report_sample_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_lab_report_sample_id ON public.lab_report USING btree (sample_id);


--
-- Name: idx_lab_report_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_lab_report_tenant_id ON public.lab_report USING btree (tenant_id);


--
-- Name: idx_lab_sample_inspection_type; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_lab_sample_inspection_type ON public.lab_sample USING btree (inspection_type);


--
-- Name: idx_lab_sample_sample_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_lab_sample_sample_code ON public.lab_sample USING btree (sample_code);


--
-- Name: idx_lab_sample_source_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_lab_sample_source_id ON public.lab_sample USING btree (source_id);


--
-- Name: idx_lab_sample_status; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_lab_sample_status ON public.lab_sample USING btree (status);


--
-- Name: idx_lab_sample_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_lab_sample_tenant_id ON public.lab_sample USING btree (tenant_id);


--
-- Name: idx_lab_test_item_sample_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_lab_test_item_sample_id ON public.lab_test_item USING btree (sample_id);


--
-- Name: idx_lab_test_item_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_lab_test_item_tenant_id ON public.lab_test_item USING btree (tenant_id);


--
-- Name: idx_locations_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_locations_deleted_at ON public.locations USING btree (deleted_at);


--
-- Name: idx_locations_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_locations_tenant_id ON public.locations USING btree (tenant_id);


--
-- Name: idx_login_logs_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_login_logs_deleted_at ON public.login_logs USING btree (deleted_at);


--
-- Name: idx_login_logs_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_login_logs_tenant_id ON public.login_logs USING btree (tenant_id);


--
-- Name: idx_material_categories_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_material_categories_deleted_at ON public.material_categories USING btree (deleted_at);


--
-- Name: idx_material_categories_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_material_categories_tenant_id ON public.material_categories USING btree (tenant_id);


--
-- Name: idx_materials_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_materials_deleted_at ON public.materials USING btree (deleted_at);


--
-- Name: idx_materials_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_materials_tenant_id ON public.materials USING btree (tenant_id);


--
-- Name: idx_mdm_bom_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_bom_code ON public.mdm_bom USING btree (bom_code);


--
-- Name: idx_mdm_bom_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_bom_deleted_at ON public.mdm_bom USING btree (deleted_at);


--
-- Name: idx_mdm_bom_item_bom; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_bom_item_bom ON public.mdm_bom_item USING btree (bom_id);


--
-- Name: idx_mdm_bom_item_bom_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_bom_item_bom_id ON public.mdm_bom_item USING btree (bom_id);


--
-- Name: idx_mdm_bom_item_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_bom_item_deleted_at ON public.mdm_bom_item USING btree (deleted_at);


--
-- Name: idx_mdm_bom_item_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_bom_item_tenant ON public.mdm_bom_item USING btree (tenant_id);


--
-- Name: idx_mdm_bom_item_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_bom_item_tenant_id ON public.mdm_bom_item USING btree (tenant_id);


--
-- Name: idx_mdm_bom_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_bom_tenant ON public.mdm_bom USING btree (tenant_id);


--
-- Name: idx_mdm_bom_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_bom_tenant_id ON public.mdm_bom USING btree (tenant_id);


--
-- Name: idx_mdm_customer_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_customer_code ON public.mdm_customer USING btree (code);


--
-- Name: idx_mdm_customer_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_customer_tenant ON public.mdm_customer USING btree (tenant_id);


--
-- Name: idx_mdm_material_category_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_material_category_deleted_at ON public.mdm_material_category USING btree (deleted_at);


--
-- Name: idx_mdm_material_category_parent_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_material_category_parent_id ON public.mdm_material_category USING btree (parent_id);


--
-- Name: idx_mdm_material_category_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_material_category_tenant ON public.mdm_material_category USING btree (tenant_id);


--
-- Name: idx_mdm_material_category_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_material_category_tenant_id ON public.mdm_material_category USING btree (tenant_id);


--
-- Name: idx_mdm_material_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_material_code ON public.mdm_material USING btree (material_code);


--
-- Name: idx_mdm_material_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_material_deleted_at ON public.mdm_material USING btree (deleted_at);


--
-- Name: idx_mdm_material_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_material_tenant ON public.mdm_material USING btree (tenant_id);


--
-- Name: idx_mdm_material_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_material_tenant_id ON public.mdm_material USING btree (tenant_id);


--
-- Name: idx_mdm_operation_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_operation_deleted_at ON public.mdm_operation USING btree (deleted_at);


--
-- Name: idx_mdm_operation_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_operation_tenant ON public.mdm_operation USING btree (tenant_id);


--
-- Name: idx_mdm_operation_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_operation_tenant_id ON public.mdm_operation USING btree (tenant_id);


--
-- Name: idx_mdm_production_line_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_production_line_deleted_at ON public.mdm_production_line USING btree (deleted_at);


--
-- Name: idx_mdm_production_line_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_production_line_tenant ON public.mdm_production_line USING btree (tenant_id);


--
-- Name: idx_mdm_production_line_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_production_line_tenant_id ON public.mdm_production_line USING btree (tenant_id);


--
-- Name: idx_mdm_production_line_workshop; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_production_line_workshop ON public.mdm_production_line USING btree (workshop_id);


--
-- Name: idx_mdm_shift_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_shift_deleted_at ON public.mdm_shift USING btree (deleted_at);


--
-- Name: idx_mdm_shift_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_shift_tenant ON public.mdm_shift USING btree (tenant_id);


--
-- Name: idx_mdm_shift_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_shift_tenant_id ON public.mdm_shift USING btree (tenant_id);


--
-- Name: idx_mdm_supplier_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_supplier_deleted_at ON public.mdm_supplier USING btree (deleted_at);


--
-- Name: idx_mdm_supplier_material_material_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_supplier_material_material_id ON public.mdm_supplier_material USING btree (material_id);


--
-- Name: idx_mdm_supplier_material_supplier_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_supplier_material_supplier_id ON public.mdm_supplier_material USING btree (supplier_id);


--
-- Name: idx_mdm_supplier_material_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_supplier_material_tenant_id ON public.mdm_supplier_material USING btree (tenant_id);


--
-- Name: idx_mdm_supplier_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_supplier_tenant_id ON public.mdm_supplier USING btree (tenant_id);


--
-- Name: idx_mdm_workshop_config_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_workshop_config_deleted_at ON public.mdm_workshop_config USING btree (deleted_at);


--
-- Name: idx_mdm_workshop_config_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_workshop_config_tenant_id ON public.mdm_workshop_config USING btree (tenant_id);


--
-- Name: idx_mdm_workshop_config_workshop_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_mdm_workshop_config_workshop_id ON public.mdm_workshop_config USING btree (workshop_id);


--
-- Name: idx_mdm_workshop_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_workshop_deleted_at ON public.mdm_workshop USING btree (deleted_at);


--
-- Name: idx_mdm_workshop_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_workshop_tenant ON public.mdm_workshop USING btree (tenant_id);


--
-- Name: idx_mdm_workshop_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_workshop_tenant_id ON public.mdm_workshop USING btree (tenant_id);


--
-- Name: idx_mdm_workstation_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_workstation_deleted_at ON public.mdm_workstation USING btree (deleted_at);


--
-- Name: idx_mdm_workstation_line; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_workstation_line ON public.mdm_workstation USING btree (line_id);


--
-- Name: idx_mdm_workstation_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_workstation_tenant ON public.mdm_workstation USING btree (tenant_id);


--
-- Name: idx_mdm_workstation_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mdm_workstation_tenant_id ON public.mdm_workstation USING btree (tenant_id);


--
-- Name: idx_mes_code_rule_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_code_rule_deleted_at ON public.mes_code_rule USING btree (deleted_at);


--
-- Name: idx_mes_code_rule_record_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_code_rule_record_deleted_at ON public.mes_code_rule_record USING btree (deleted_at);


--
-- Name: idx_mes_code_rule_record_gen_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_mes_code_rule_record_gen_code ON public.mes_code_rule_record USING btree (gen_code);


--
-- Name: idx_mes_code_rule_record_rule_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_code_rule_record_rule_id ON public.mes_code_rule_record USING btree (rule_id);


--
-- Name: idx_mes_code_rule_record_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_code_rule_record_tenant_id ON public.mes_code_rule_record USING btree (tenant_id);


--
-- Name: idx_mes_code_rule_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_code_rule_tenant_id ON public.mes_code_rule USING btree (tenant_id);


--
-- Name: idx_mes_container_lifecycle_container_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_container_lifecycle_container_id ON public.mes_container_lifecycle USING btree (container_id);


--
-- Name: idx_mes_container_lifecycle_location_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_container_lifecycle_location_id ON public.mes_container_lifecycle USING btree (location_id);


--
-- Name: idx_mes_container_lifecycle_operator_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_container_lifecycle_operator_id ON public.mes_container_lifecycle USING btree (operator_id);


--
-- Name: idx_mes_container_lifecycle_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_container_lifecycle_tenant_id ON public.mes_container_lifecycle USING btree (tenant_id);


--
-- Name: idx_mes_container_maintenance_container_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_container_maintenance_container_id ON public.mes_container_maintenance USING btree (container_id);


--
-- Name: idx_mes_container_maintenance_technician_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_container_maintenance_technician_id ON public.mes_container_maintenance USING btree (technician_id);


--
-- Name: idx_mes_container_maintenance_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_container_maintenance_tenant_id ON public.mes_container_maintenance USING btree (tenant_id);


--
-- Name: idx_mes_electronic_sop_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_electronic_sop_deleted_at ON public.mes_electronic_sop USING btree (deleted_at);


--
-- Name: idx_mes_electronic_sop_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_electronic_sop_tenant_id ON public.mes_electronic_sop USING btree (tenant_id);


--
-- Name: idx_mes_first_last_inspect_inspect_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_mes_first_last_inspect_inspect_no ON public.mes_first_last_inspect USING btree (inspect_no);


--
-- Name: idx_mes_first_last_inspect_inspector_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_first_last_inspect_inspector_id ON public.mes_first_last_inspect USING btree (inspector_id);


--
-- Name: idx_mes_first_last_inspect_production_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_first_last_inspect_production_order_id ON public.mes_first_last_inspect USING btree (production_order_id);


--
-- Name: idx_mes_first_last_inspect_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_first_last_inspect_tenant_id ON public.mes_first_last_inspect USING btree (tenant_id);


--
-- Name: idx_mes_first_last_inspect_workshop_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_first_last_inspect_workshop_id ON public.mes_first_last_inspect USING btree (workshop_id);


--
-- Name: idx_mes_flow_card_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_flow_card_deleted_at ON public.mes_flow_card USING btree (deleted_at);


--
-- Name: idx_mes_flow_card_detail_card_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_flow_card_detail_card_id ON public.mes_flow_card_detail USING btree (card_id);


--
-- Name: idx_mes_flow_card_detail_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_flow_card_detail_deleted_at ON public.mes_flow_card_detail USING btree (deleted_at);


--
-- Name: idx_mes_flow_card_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_flow_card_tenant_id ON public.mes_flow_card USING btree (tenant_id);


--
-- Name: idx_mes_job_report_log_report_time; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_job_report_log_report_time ON public.mes_job_report_log USING btree (report_time);


--
-- Name: idx_mes_job_report_log_reporter_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_job_report_log_reporter_id ON public.mes_job_report_log USING btree (reporter_id);


--
-- Name: idx_mes_job_report_log_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_job_report_log_tenant_id ON public.mes_job_report_log USING btree (tenant_id);


--
-- Name: idx_mes_job_report_log_work_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_job_report_log_work_order_id ON public.mes_job_report_log USING btree (work_order_id);


--
-- Name: idx_mes_order_day_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_order_day_deleted_at ON public.mes_order_day USING btree (deleted_at);


--
-- Name: idx_mes_order_day_item_day_plan_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_order_day_item_day_plan_id ON public.mes_order_day_item USING btree (day_plan_id);


--
-- Name: idx_mes_order_day_item_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_order_day_item_deleted_at ON public.mes_order_day_item USING btree (deleted_at);


--
-- Name: idx_mes_order_day_item_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_order_day_item_tenant_id ON public.mes_order_day_item USING btree (tenant_id);


--
-- Name: idx_mes_order_day_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_order_day_tenant_id ON public.mes_order_day USING btree (tenant_id);


--
-- Name: idx_mes_order_month_audit_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_order_month_audit_deleted_at ON public.mes_order_month_audit USING btree (deleted_at);


--
-- Name: idx_mes_order_month_audit_month_plan_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_order_month_audit_month_plan_id ON public.mes_order_month_audit USING btree (month_plan_id);


--
-- Name: idx_mes_order_month_audit_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_order_month_audit_tenant_id ON public.mes_order_month_audit USING btree (tenant_id);


--
-- Name: idx_mes_order_month_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_order_month_deleted_at ON public.mes_order_month USING btree (deleted_at);


--
-- Name: idx_mes_order_month_item_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_order_month_item_deleted_at ON public.mes_order_month_item USING btree (deleted_at);


--
-- Name: idx_mes_order_month_item_month_plan_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_order_month_item_month_plan_id ON public.mes_order_month_item USING btree (month_plan_id);


--
-- Name: idx_mes_order_month_item_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_order_month_item_tenant_id ON public.mes_order_month_item USING btree (tenant_id);


--
-- Name: idx_mes_order_month_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_order_month_tenant_id ON public.mes_order_month USING btree (tenant_id);


--
-- Name: idx_mes_package_container_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_package_container_id ON public.mes_package USING btree (container_id);


--
-- Name: idx_mes_package_customer_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_package_customer_id ON public.mes_package USING btree (customer_id);


--
-- Name: idx_mes_package_package_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_mes_package_package_no ON public.mes_package USING btree (package_no);


--
-- Name: idx_mes_package_production_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_package_production_order_id ON public.mes_package USING btree (production_order_id);


--
-- Name: idx_mes_package_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_package_tenant_id ON public.mes_package USING btree (tenant_id);


--
-- Name: idx_mes_package_workshop_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_package_workshop_id ON public.mes_package USING btree (workshop_id);


--
-- Name: idx_mes_person_skill_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_person_skill_deleted_at ON public.mes_person_skill USING btree (deleted_at);


--
-- Name: idx_mes_person_skill_person_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_person_skill_person_id ON public.mes_person_skill USING btree (person_id);


--
-- Name: idx_mes_person_skill_score_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_person_skill_score_deleted_at ON public.mes_person_skill_score USING btree (deleted_at);


--
-- Name: idx_mes_person_skill_score_person_skill_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_person_skill_score_person_skill_id ON public.mes_person_skill_score USING btree (person_skill_id);


--
-- Name: idx_mes_person_skill_score_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_person_skill_score_tenant_id ON public.mes_person_skill_score USING btree (tenant_id);


--
-- Name: idx_mes_person_skill_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_person_skill_tenant_id ON public.mes_person_skill USING btree (tenant_id);


--
-- Name: idx_mes_person_skill_workshop_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_person_skill_workshop_id ON public.mes_person_skill USING btree (workshop_id);


--
-- Name: idx_mes_person_skill_workstation_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_person_skill_workstation_id ON public.mes_person_skill USING btree (workstation_id);


--
-- Name: idx_mes_process_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_process_deleted_at ON public.mes_process USING btree (deleted_at);


--
-- Name: idx_mes_process_operation_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_process_operation_deleted_at ON public.mes_process_operation USING btree (deleted_at);


--
-- Name: idx_mes_process_operation_process_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_process_operation_process_id ON public.mes_process_operation USING btree (process_id);


--
-- Name: idx_mes_process_operation_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_process_operation_tenant_id ON public.mes_process_operation USING btree (tenant_id);


--
-- Name: idx_mes_process_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_process_tenant_id ON public.mes_process USING btree (tenant_id);


--
-- Name: idx_mes_team_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_team_deleted_at ON public.mes_team USING btree (deleted_at);


--
-- Name: idx_mes_team_member_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_team_member_deleted_at ON public.mes_team_member USING btree (deleted_at);


--
-- Name: idx_mes_team_member_team_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_team_member_team_id ON public.mes_team_member USING btree (team_id);


--
-- Name: idx_mes_team_member_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_team_member_tenant_id ON public.mes_team_member USING btree (tenant_id);


--
-- Name: idx_mes_team_shift_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_team_shift_deleted_at ON public.mes_team_shift USING btree (deleted_at);


--
-- Name: idx_mes_team_shift_shift_date; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_team_shift_shift_date ON public.mes_team_shift USING btree (shift_date);


--
-- Name: idx_mes_team_shift_team_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_team_shift_team_id ON public.mes_team_shift USING btree (team_id);


--
-- Name: idx_mes_team_shift_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_team_shift_tenant_id ON public.mes_team_shift USING btree (tenant_id);


--
-- Name: idx_mes_team_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_team_tenant_id ON public.mes_team USING btree (tenant_id);


--
-- Name: idx_mes_work_scheduling_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_work_scheduling_deleted_at ON public.mes_work_scheduling USING btree (deleted_at);


--
-- Name: idx_mes_work_scheduling_detail_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_work_scheduling_detail_deleted_at ON public.mes_work_scheduling_detail USING btree (deleted_at);


--
-- Name: idx_mes_work_scheduling_detail_scheduling_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_work_scheduling_detail_scheduling_id ON public.mes_work_scheduling_detail USING btree (scheduling_id);


--
-- Name: idx_mes_work_scheduling_detail_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_work_scheduling_detail_tenant_id ON public.mes_work_scheduling_detail USING btree (tenant_id);


--
-- Name: idx_mes_work_scheduling_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mes_work_scheduling_tenant_id ON public.mes_work_scheduling USING btree (tenant_id);


--
-- Name: idx_mps_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mps_deleted_at ON public.mps USING btree (deleted_at);


--
-- Name: idx_mps_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mps_tenant_id ON public.mps USING btree (tenant_id);


--
-- Name: idx_mrp_items_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mrp_items_deleted_at ON public.mrp_items USING btree (deleted_at);


--
-- Name: idx_mrp_items_mrp_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mrp_items_mrp_id ON public.mrp_items USING btree (mrp_id);


--
-- Name: idx_mrps_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mrps_deleted_at ON public.mrps USING btree (deleted_at);


--
-- Name: idx_mrps_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_mrps_tenant_id ON public.mrps USING btree (tenant_id);


--
-- Name: idx_ncrs_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_ncrs_deleted_at ON public.ncrs USING btree (deleted_at);


--
-- Name: idx_ncrs_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_ncrs_tenant_id ON public.ncrs USING btree (tenant_id);


--
-- Name: idx_notification_channel_channel_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_notification_channel_channel_code ON public.notification_channel USING btree (channel_code);


--
-- Name: idx_notification_channel_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_notification_channel_tenant_id ON public.notification_channel USING btree (tenant_id);


--
-- Name: idx_oee_report_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_oee_report_deleted_at ON public.oee_report USING btree (deleted_at);


--
-- Name: idx_oee_report_line_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_oee_report_line_id ON public.oee_report USING btree (line_id);


--
-- Name: idx_oee_report_report_date; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_oee_report_report_date ON public.oee_report USING btree (report_date);


--
-- Name: idx_oee_report_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_oee_report_tenant_id ON public.oee_report USING btree (tenant_id);


--
-- Name: idx_oee_report_workshop_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_oee_report_workshop_id ON public.oee_report USING btree (workshop_id);


--
-- Name: idx_oper_logs_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_oper_logs_deleted_at ON public.oper_logs USING btree (deleted_at);


--
-- Name: idx_oper_logs_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_oper_logs_tenant_id ON public.oper_logs USING btree (tenant_id);


--
-- Name: idx_oqcs_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_oqcs_deleted_at ON public.oqcs USING btree (deleted_at);


--
-- Name: idx_oqcs_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_oqcs_tenant_id ON public.oqcs USING btree (tenant_id);


--
-- Name: idx_pro_dispatch_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_dispatch_deleted_at ON public.pro_dispatch USING btree (deleted_at);


--
-- Name: idx_pro_dispatch_order; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_dispatch_order ON public.pro_dispatch USING btree (order_id);


--
-- Name: idx_pro_dispatch_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_dispatch_order_id ON public.pro_dispatch USING btree (order_id);


--
-- Name: idx_pro_dispatch_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_dispatch_tenant ON public.pro_dispatch USING btree (tenant_id);


--
-- Name: idx_pro_dispatch_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_dispatch_tenant_id ON public.pro_dispatch USING btree (tenant_id);


--
-- Name: idx_pro_production_order_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_production_order_deleted_at ON public.pro_production_order USING btree (deleted_at);


--
-- Name: idx_pro_production_order_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_production_order_no ON public.pro_production_order USING btree (order_no);


--
-- Name: idx_pro_production_order_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_production_order_tenant ON public.pro_production_order USING btree (tenant_id);


--
-- Name: idx_pro_production_order_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_production_order_tenant_id ON public.pro_production_order USING btree (tenant_id);


--
-- Name: idx_pro_production_report_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_production_report_deleted_at ON public.pro_production_report USING btree (deleted_at);


--
-- Name: idx_pro_production_report_order; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_production_report_order ON public.pro_production_report USING btree (order_id);


--
-- Name: idx_pro_production_report_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_production_report_order_id ON public.pro_production_report USING btree (order_id);


--
-- Name: idx_pro_production_report_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_production_report_tenant ON public.pro_production_report USING btree (tenant_id);


--
-- Name: idx_pro_production_report_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_production_report_tenant_id ON public.pro_production_report USING btree (tenant_id);


--
-- Name: idx_pro_sales_order_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_sales_order_deleted_at ON public.pro_sales_order USING btree (deleted_at);


--
-- Name: idx_pro_sales_order_item_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_sales_order_item_deleted_at ON public.pro_sales_order_item USING btree (deleted_at);


--
-- Name: idx_pro_sales_order_item_order; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_sales_order_item_order ON public.pro_sales_order_item USING btree (order_id);


--
-- Name: idx_pro_sales_order_item_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_sales_order_item_order_id ON public.pro_sales_order_item USING btree (order_id);


--
-- Name: idx_pro_sales_order_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_sales_order_no ON public.pro_sales_order USING btree (order_no);


--
-- Name: idx_pro_sales_order_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_sales_order_order_id ON public.pro_sales_order USING btree (order_id);


--
-- Name: idx_pro_sales_order_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_sales_order_tenant ON public.pro_sales_order USING btree (tenant_id);


--
-- Name: idx_pro_sales_order_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_pro_sales_order_tenant_id ON public.pro_sales_order USING btree (tenant_id);


--
-- Name: idx_processes_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_processes_deleted_at ON public.processes USING btree (deleted_at);


--
-- Name: idx_processes_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_processes_tenant_id ON public.processes USING btree (tenant_id);


--
-- Name: idx_production_daily_report_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_production_daily_report_deleted_at ON public.production_daily_report USING btree (deleted_at);


--
-- Name: idx_production_daily_report_report_date; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_production_daily_report_report_date ON public.production_daily_report USING btree (report_date);


--
-- Name: idx_production_daily_report_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_production_daily_report_tenant_id ON public.production_daily_report USING btree (tenant_id);


--
-- Name: idx_production_daily_report_workshop_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_production_daily_report_workshop_id ON public.production_daily_report USING btree (workshop_id);


--
-- Name: idx_production_lines_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_production_lines_deleted_at ON public.production_lines USING btree (deleted_at);


--
-- Name: idx_production_lines_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_production_lines_tenant_id ON public.production_lines USING btree (tenant_id);


--
-- Name: idx_production_offline_handle_method; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_production_offline_handle_method ON public.production_offline USING btree (handle_method);


--
-- Name: idx_production_offline_item_offline_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_production_offline_item_offline_id ON public.production_offline_item USING btree (offline_id);


--
-- Name: idx_production_offline_item_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_production_offline_item_tenant_id ON public.production_offline_item USING btree (tenant_id);


--
-- Name: idx_production_offline_offline_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_production_offline_offline_code ON public.production_offline USING btree (offline_code);


--
-- Name: idx_production_offline_offline_type; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_production_offline_offline_type ON public.production_offline USING btree (offline_type);


--
-- Name: idx_production_offline_product_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_production_offline_product_id ON public.production_offline USING btree (product_id);


--
-- Name: idx_production_offline_status; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_production_offline_status ON public.production_offline USING btree (status);


--
-- Name: idx_production_offline_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_production_offline_tenant_id ON public.production_offline USING btree (tenant_id);


--
-- Name: idx_production_offline_work_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_production_offline_work_order_id ON public.production_offline USING btree (work_order_id);


--
-- Name: idx_production_orders_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_production_orders_deleted_at ON public.production_orders USING btree (deleted_at);


--
-- Name: idx_production_orders_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_production_orders_tenant_id ON public.production_orders USING btree (tenant_id);


--
-- Name: idx_production_reports_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_production_reports_deleted_at ON public.production_reports USING btree (deleted_at);


--
-- Name: idx_production_reports_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_production_reports_order_id ON public.production_reports USING btree (order_id);


--
-- Name: idx_production_reports_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_production_reports_tenant_id ON public.production_reports USING btree (tenant_id);


--
-- Name: idx_qc_aql_level_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_aql_level_deleted_at ON public.qc_aql_level USING btree (deleted_at);


--
-- Name: idx_qc_aql_level_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_aql_level_tenant_id ON public.qc_aql_level USING btree (tenant_id);


--
-- Name: idx_qc_aql_table_row_aql_level_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_aql_table_row_aql_level_id ON public.qc_aql_table_row USING btree (aql_level_id);


--
-- Name: idx_qc_aql_table_row_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_aql_table_row_deleted_at ON public.qc_aql_table_row USING btree (deleted_at);


--
-- Name: idx_qc_aql_table_row_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_aql_table_row_tenant_id ON public.qc_aql_table_row USING btree (tenant_id);


--
-- Name: idx_qc_defect_code_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_defect_code_deleted_at ON public.qc_defect_code USING btree (deleted_at);


--
-- Name: idx_qc_defect_code_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_defect_code_tenant ON public.qc_defect_code USING btree (tenant_id);


--
-- Name: idx_qc_defect_code_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_defect_code_tenant_id ON public.qc_defect_code USING btree (tenant_id);


--
-- Name: idx_qc_defect_record_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_defect_record_deleted_at ON public.qc_defect_record USING btree (deleted_at);


--
-- Name: idx_qc_defect_record_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_defect_record_tenant ON public.qc_defect_record USING btree (tenant_id);


--
-- Name: idx_qc_defect_record_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_defect_record_tenant_id ON public.qc_defect_record USING btree (tenant_id);


--
-- Name: idx_qc_fqc_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_fqc_deleted_at ON public.qc_fqc USING btree (deleted_at);


--
-- Name: idx_qc_fqc_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_fqc_tenant ON public.qc_fqc USING btree (tenant_id);


--
-- Name: idx_qc_fqc_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_fqc_tenant_id ON public.qc_fqc USING btree (tenant_id);


--
-- Name: idx_qc_inspection_characteristic_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_qc_inspection_characteristic_code ON public.qc_inspection_characteristic USING btree (code);


--
-- Name: idx_qc_inspection_characteristic_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_inspection_characteristic_tenant_id ON public.qc_inspection_characteristic USING btree (tenant_id);


--
-- Name: idx_qc_inspection_feature_feature_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_qc_inspection_feature_feature_code ON public.qc_inspection_feature USING btree (feature_code);


--
-- Name: idx_qc_inspection_feature_inspection_type; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_inspection_feature_inspection_type ON public.qc_inspection_feature USING btree (inspection_type);


--
-- Name: idx_qc_inspection_feature_product_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_inspection_feature_product_code ON public.qc_inspection_feature USING btree (product_code);


--
-- Name: idx_qc_inspection_feature_product_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_inspection_feature_product_id ON public.qc_inspection_feature USING btree (product_id);


--
-- Name: idx_qc_inspection_feature_status; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_inspection_feature_status ON public.qc_inspection_feature USING btree (status);


--
-- Name: idx_qc_inspection_feature_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_inspection_feature_tenant_id ON public.qc_inspection_feature USING btree (tenant_id);


--
-- Name: idx_qc_ipqc_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_ipqc_deleted_at ON public.qc_ipqc USING btree (deleted_at);


--
-- Name: idx_qc_ipqc_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_ipqc_tenant ON public.qc_ipqc USING btree (tenant_id);


--
-- Name: idx_qc_ipqc_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_ipqc_tenant_id ON public.qc_ipqc USING btree (tenant_id);


--
-- Name: idx_qc_iqc_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_iqc_deleted_at ON public.qc_iqc USING btree (deleted_at);


--
-- Name: idx_qc_iqc_item_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_iqc_item_deleted_at ON public.qc_iqc_item USING btree (deleted_at);


--
-- Name: idx_qc_iqc_item_iqc; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_iqc_item_iqc ON public.qc_iqc_item USING btree (iqc_id);


--
-- Name: idx_qc_iqc_item_iqc_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_iqc_item_iqc_id ON public.qc_iqc_item USING btree (iqc_id);


--
-- Name: idx_qc_iqc_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_iqc_tenant ON public.qc_iqc USING btree (tenant_id);


--
-- Name: idx_qc_iqc_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_iqc_tenant_id ON public.qc_iqc USING btree (tenant_id);


--
-- Name: idx_qc_ncr_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_ncr_deleted_at ON public.qc_ncr USING btree (deleted_at);


--
-- Name: idx_qc_ncr_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_ncr_tenant ON public.qc_ncr USING btree (tenant_id);


--
-- Name: idx_qc_ncr_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_ncr_tenant_id ON public.qc_ncr USING btree (tenant_id);


--
-- Name: idx_qc_oqc_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_oqc_deleted_at ON public.qc_oqc USING btree (deleted_at);


--
-- Name: idx_qc_oqc_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_oqc_tenant ON public.qc_oqc USING btree (tenant_id);


--
-- Name: idx_qc_oqc_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_oqc_tenant_id ON public.qc_oqc USING btree (tenant_id);


--
-- Name: idx_qc_sampling_plan_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_sampling_plan_deleted_at ON public.qc_sampling_plan USING btree (deleted_at);


--
-- Name: idx_qc_sampling_plan_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_sampling_plan_tenant_id ON public.qc_sampling_plan USING btree (tenant_id);


--
-- Name: idx_qc_spc_data_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_spc_data_deleted_at ON public.qc_spc_data USING btree (deleted_at);


--
-- Name: idx_qc_spc_data_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_spc_data_tenant ON public.qc_spc_data USING btree (tenant_id);


--
-- Name: idx_qc_spc_data_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qc_spc_data_tenant_id ON public.qc_spc_data USING btree (tenant_id);


--
-- Name: idx_qms_sampling_plan_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qms_sampling_plan_deleted_at ON public.qms_sampling_plan USING btree (deleted_at);


--
-- Name: idx_qms_sampling_plan_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qms_sampling_plan_tenant_id ON public.qms_sampling_plan USING btree (tenant_id);


--
-- Name: idx_qms_sampling_record_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qms_sampling_record_deleted_at ON public.qms_sampling_record USING btree (deleted_at);


--
-- Name: idx_qms_sampling_record_inspection_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qms_sampling_record_inspection_id ON public.qms_sampling_record USING btree (inspection_id);


--
-- Name: idx_qms_sampling_record_plan_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qms_sampling_record_plan_id ON public.qms_sampling_record USING btree (plan_id);


--
-- Name: idx_qms_sampling_record_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qms_sampling_record_tenant_id ON public.qms_sampling_record USING btree (tenant_id);


--
-- Name: idx_qms_sampling_rule_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qms_sampling_rule_deleted_at ON public.qms_sampling_rule USING btree (deleted_at);


--
-- Name: idx_qms_sampling_rule_plan_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_qms_sampling_rule_plan_id ON public.qms_sampling_rule USING btree (plan_id);


--
-- Name: idx_quality_weekly_report_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_quality_weekly_report_deleted_at ON public.quality_weekly_report USING btree (deleted_at);


--
-- Name: idx_quality_weekly_report_report_week; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_quality_weekly_report_report_week ON public.quality_weekly_report USING btree (report_week);


--
-- Name: idx_quality_weekly_report_report_year; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_quality_weekly_report_report_year ON public.quality_weekly_report USING btree (report_year);


--
-- Name: idx_quality_weekly_report_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_quality_weekly_report_tenant_id ON public.quality_weekly_report USING btree (tenant_id);


--
-- Name: idx_quality_weekly_report_workshop_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_quality_weekly_report_workshop_id ON public.quality_weekly_report USING btree (workshop_id);


--
-- Name: idx_receive_order_items_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_receive_order_items_deleted_at ON public.receive_order_items USING btree (deleted_at);


--
-- Name: idx_receive_order_items_receive_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_receive_order_items_receive_id ON public.receive_order_items USING btree (receive_id);


--
-- Name: idx_receive_orders_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_receive_orders_deleted_at ON public.receive_orders USING btree (deleted_at);


--
-- Name: idx_receive_orders_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_receive_orders_tenant_id ON public.receive_orders USING btree (tenant_id);


--
-- Name: idx_resources_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_resources_deleted_at ON public.resources USING btree (deleted_at);


--
-- Name: idx_resources_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_resources_tenant_id ON public.resources USING btree (tenant_id);


--
-- Name: idx_route_operations_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_route_operations_deleted_at ON public.route_operations USING btree (deleted_at);


--
-- Name: idx_route_operations_route_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_route_operations_route_id ON public.route_operations USING btree (route_id);


--
-- Name: idx_routes_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_routes_deleted_at ON public.routes USING btree (deleted_at);


--
-- Name: idx_routes_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_routes_tenant_id ON public.routes USING btree (tenant_id);


--
-- Name: idx_sales_order_items_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sales_order_items_deleted_at ON public.sales_order_items USING btree (deleted_at);


--
-- Name: idx_sales_order_items_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sales_order_items_order_id ON public.sales_order_items USING btree (order_id);


--
-- Name: idx_sales_orders_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sales_orders_deleted_at ON public.sales_orders USING btree (deleted_at);


--
-- Name: idx_sales_orders_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sales_orders_tenant_id ON public.sales_orders USING btree (tenant_id);


--
-- Name: idx_schedule_plans_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_schedule_plans_deleted_at ON public.schedule_plans USING btree (deleted_at);


--
-- Name: idx_schedule_plans_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_schedule_plans_tenant_id ON public.schedule_plans USING btree (tenant_id);


--
-- Name: idx_schedule_results_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_schedule_results_deleted_at ON public.schedule_results USING btree (deleted_at);


--
-- Name: idx_schedule_results_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_schedule_results_order_id ON public.schedule_results USING btree (order_id);


--
-- Name: idx_schedule_results_plan_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_schedule_results_plan_id ON public.schedule_results USING btree (plan_id);


--
-- Name: idx_scp_customer_inquiry_inquiry_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_scp_customer_inquiry_inquiry_no ON public.scp_customer_inquiry USING btree (inquiry_no);


--
-- Name: idx_scp_customer_inquiry_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_customer_inquiry_tenant_id ON public.scp_customer_inquiry USING btree (tenant_id);


--
-- Name: idx_scp_inquiry_item_inquiry_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_inquiry_item_inquiry_id ON public.scp_inquiry_item USING btree (inquiry_id);


--
-- Name: idx_scp_mrs_item_mrs_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_mrs_item_mrs_id ON public.scp_mrs_item USING btree (mrs_id);


--
-- Name: idx_scp_mrs_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_mrs_tenant_id ON public.scp_mrs USING btree (tenant_id);


--
-- Name: idx_scp_po_change_log_po_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_po_change_log_po_id ON public.scp_po_change_log USING btree (po_id);


--
-- Name: idx_scp_po_change_log_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_po_change_log_tenant_id ON public.scp_po_change_log USING btree (tenant_id);


--
-- Name: idx_scp_purchase_order_item_po_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_purchase_order_item_po_id ON public.scp_purchase_order_item USING btree (po_id);


--
-- Name: idx_scp_purchase_order_po_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_scp_purchase_order_po_no ON public.scp_purchase_order USING btree (po_no);


--
-- Name: idx_scp_purchase_order_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_purchase_order_tenant_id ON public.scp_purchase_order USING btree (tenant_id);


--
-- Name: idx_scp_purchase_plan_item_plan_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_purchase_plan_item_plan_id ON public.scp_purchase_plan_item USING btree (plan_id);


--
-- Name: idx_scp_purchase_plan_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_purchase_plan_tenant_id ON public.scp_purchase_plan USING btree (tenant_id);


--
-- Name: idx_scp_qad_sync_log_mom_doc_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_qad_sync_log_mom_doc_no ON public.scp_qad_sync_log USING btree (mom_doc_no);


--
-- Name: idx_scp_qad_sync_log_qad_doc_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_qad_sync_log_qad_doc_no ON public.scp_qad_sync_log USING btree (qad_doc_no);


--
-- Name: idx_scp_qad_sync_log_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_qad_sync_log_tenant_id ON public.scp_qad_sync_log USING btree (tenant_id);


--
-- Name: idx_scp_quote_comparison_comparison_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_scp_quote_comparison_comparison_no ON public.scp_quote_comparison USING btree (comparison_no);


--
-- Name: idx_scp_quote_comparison_rfq_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_quote_comparison_rfq_id ON public.scp_quote_comparison USING btree (rfq_id);


--
-- Name: idx_scp_quote_comparison_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_quote_comparison_tenant_id ON public.scp_quote_comparison USING btree (tenant_id);


--
-- Name: idx_scp_quote_item_quote_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_quote_item_quote_id ON public.scp_quote_item USING btree (quote_id);


--
-- Name: idx_scp_rfq_invite_rfq_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_rfq_invite_rfq_id ON public.scp_rfq_invite USING btree (rfq_id);


--
-- Name: idx_scp_rfq_invite_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_rfq_invite_tenant_id ON public.scp_rfq_invite USING btree (tenant_id);


--
-- Name: idx_scp_rfq_item_rfq_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_rfq_item_rfq_id ON public.scp_rfq_item USING btree (rfq_id);


--
-- Name: idx_scp_rfq_rfq_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_scp_rfq_rfq_no ON public.scp_rfq USING btree (rfq_no);


--
-- Name: idx_scp_rfq_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_rfq_tenant_id ON public.scp_rfq USING btree (tenant_id);


--
-- Name: idx_scp_sales_order_item_so_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_sales_order_item_so_id ON public.scp_sales_order_item USING btree (so_id);


--
-- Name: idx_scp_sales_order_so_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_scp_sales_order_so_no ON public.scp_sales_order USING btree (so_no);


--
-- Name: idx_scp_sales_order_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_sales_order_tenant_id ON public.scp_sales_order USING btree (tenant_id);


--
-- Name: idx_scp_so_change_log_so_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_so_change_log_so_id ON public.scp_so_change_log USING btree (so_id);


--
-- Name: idx_scp_so_change_log_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_so_change_log_tenant_id ON public.scp_so_change_log USING btree (tenant_id);


--
-- Name: idx_scp_supplier_bank_supplier_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_supplier_bank_supplier_id ON public.scp_supplier_bank USING btree (supplier_id);


--
-- Name: idx_scp_supplier_bank_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_supplier_bank_tenant_id ON public.scp_supplier_bank USING btree (tenant_id);


--
-- Name: idx_scp_supplier_contact_supplier_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_supplier_contact_supplier_id ON public.scp_supplier_contact USING btree (supplier_id);


--
-- Name: idx_scp_supplier_contact_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_supplier_contact_tenant_id ON public.scp_supplier_contact USING btree (tenant_id);


--
-- Name: idx_scp_supplier_delivery_record_supplier_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_supplier_delivery_record_supplier_id ON public.scp_supplier_delivery_record USING btree (supplier_id);


--
-- Name: idx_scp_supplier_delivery_record_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_supplier_delivery_record_tenant_id ON public.scp_supplier_delivery_record USING btree (tenant_id);


--
-- Name: idx_scp_supplier_grade_standard_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_supplier_grade_standard_tenant_id ON public.scp_supplier_grade_standard USING btree (tenant_id);


--
-- Name: idx_scp_supplier_kpi_supplier_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_supplier_kpi_supplier_id ON public.scp_supplier_kpi USING btree (supplier_id);


--
-- Name: idx_scp_supplier_kpi_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_supplier_kpi_tenant_id ON public.scp_supplier_kpi USING btree (tenant_id);


--
-- Name: idx_scp_supplier_purchase_info_supplier_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_scp_supplier_purchase_info_supplier_id ON public.scp_supplier_purchase_info USING btree (supplier_id);


--
-- Name: idx_scp_supplier_purchase_info_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_supplier_purchase_info_tenant_id ON public.scp_supplier_purchase_info USING btree (tenant_id);


--
-- Name: idx_scp_supplier_quality_record_supplier_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_supplier_quality_record_supplier_id ON public.scp_supplier_quality_record USING btree (supplier_id);


--
-- Name: idx_scp_supplier_quality_record_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_supplier_quality_record_tenant_id ON public.scp_supplier_quality_record USING btree (tenant_id);


--
-- Name: idx_scp_supplier_quote_quote_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_scp_supplier_quote_quote_no ON public.scp_supplier_quote USING btree (quote_no);


--
-- Name: idx_scp_supplier_quote_rfq_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_supplier_quote_rfq_id ON public.scp_supplier_quote USING btree (rfq_id);


--
-- Name: idx_scp_supplier_quote_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_scp_supplier_quote_tenant_id ON public.scp_supplier_quote USING btree (tenant_id);


--
-- Name: idx_serial_numbers_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_serial_numbers_deleted_at ON public.serial_numbers USING btree (deleted_at);


--
-- Name: idx_serial_numbers_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_serial_numbers_tenant_id ON public.serial_numbers USING btree (tenant_id);


--
-- Name: idx_shifts_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_shifts_deleted_at ON public.shifts USING btree (deleted_at);


--
-- Name: idx_shifts_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_shifts_tenant_id ON public.shifts USING btree (tenant_id);


--
-- Name: idx_spc_data_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_spc_data_deleted_at ON public.spc_data USING btree (deleted_at);


--
-- Name: idx_spc_data_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_spc_data_tenant_id ON public.spc_data USING btree (tenant_id);


--
-- Name: idx_stock_checks_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_stock_checks_deleted_at ON public.stock_checks USING btree (deleted_at);


--
-- Name: idx_stock_checks_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_stock_checks_tenant_id ON public.stock_checks USING btree (tenant_id);


--
-- Name: idx_supplier_asns_asn_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_supplier_asns_asn_no ON public.supplier_asns USING btree (asn_no);


--
-- Name: idx_sys_dept_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_dept_deleted_at ON public.sys_dept USING btree (deleted_at);


--
-- Name: idx_sys_dept_parent; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_dept_parent ON public.sys_dept USING btree (parent_id);


--
-- Name: idx_sys_dept_parent_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_dept_parent_id ON public.sys_dept USING btree (parent_id);


--
-- Name: idx_sys_dept_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_dept_tenant ON public.sys_dept USING btree (tenant_id);


--
-- Name: idx_sys_dept_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_dept_tenant_id ON public.sys_dept USING btree (tenant_id);


--
-- Name: idx_sys_dict_data_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_dict_data_deleted_at ON public.sys_dict_data USING btree (deleted_at);


--
-- Name: idx_sys_dict_data_dict_type; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_dict_data_dict_type ON public.sys_dict_data USING btree (dict_type);


--
-- Name: idx_sys_dict_data_type; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_dict_data_type ON public.sys_dict_data USING btree (dict_type);


--
-- Name: idx_sys_dict_type_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_dict_type_deleted_at ON public.sys_dict_type USING btree (deleted_at);


--
-- Name: idx_sys_dict_type_dict_type; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_sys_dict_type_dict_type ON public.sys_dict_type USING btree (dict_type);


--
-- Name: idx_sys_interface_config_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_interface_config_deleted_at ON public.sys_interface_config USING btree (deleted_at);


--
-- Name: idx_sys_interface_config_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_interface_config_tenant_id ON public.sys_interface_config USING btree (tenant_id);


--
-- Name: idx_sys_interface_exec_log_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_interface_exec_log_deleted_at ON public.sys_interface_exec_log USING btree (deleted_at);


--
-- Name: idx_sys_interface_exec_log_interface_config_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_interface_exec_log_interface_config_id ON public.sys_interface_exec_log USING btree (interface_config_id);


--
-- Name: idx_sys_interface_field_map_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_interface_field_map_deleted_at ON public.sys_interface_field_map USING btree (deleted_at);


--
-- Name: idx_sys_interface_field_map_interface_config_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_interface_field_map_interface_config_id ON public.sys_interface_field_map USING btree (interface_config_id);


--
-- Name: idx_sys_interface_trigger_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_interface_trigger_deleted_at ON public.sys_interface_trigger USING btree (deleted_at);


--
-- Name: idx_sys_interface_trigger_interface_config_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_interface_trigger_interface_config_id ON public.sys_interface_trigger USING btree (interface_config_id);


--
-- Name: idx_sys_login_log_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_login_log_deleted_at ON public.sys_login_log USING btree (deleted_at);


--
-- Name: idx_sys_login_log_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_login_log_tenant ON public.sys_login_log USING btree (tenant_id);


--
-- Name: idx_sys_login_log_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_login_log_tenant_id ON public.sys_login_log USING btree (tenant_id);


--
-- Name: idx_sys_menu_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_menu_deleted_at ON public.sys_menu USING btree (deleted_at);


--
-- Name: idx_sys_menu_parent; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_menu_parent ON public.sys_menu USING btree (parent_id);


--
-- Name: idx_sys_menu_parent_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_menu_parent_id ON public.sys_menu USING btree (parent_id);


--
-- Name: idx_sys_menu_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_menu_tenant ON public.sys_menu USING btree (tenant_id);


--
-- Name: idx_sys_menu_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_menu_tenant_id ON public.sys_menu USING btree (tenant_id);


--
-- Name: idx_sys_notice_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_notice_deleted_at ON public.sys_notice USING btree (deleted_at);


--
-- Name: idx_sys_notice_read_record_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_notice_read_record_deleted_at ON public.sys_notice_read_record USING btree (deleted_at);


--
-- Name: idx_sys_notice_read_record_notice_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_notice_read_record_notice_id ON public.sys_notice_read_record USING btree (notice_id);


--
-- Name: idx_sys_notice_read_record_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_notice_read_record_tenant_id ON public.sys_notice_read_record USING btree (tenant_id);


--
-- Name: idx_sys_notice_read_record_user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_notice_read_record_user_id ON public.sys_notice_read_record USING btree (user_id);


--
-- Name: idx_sys_notice_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_notice_tenant_id ON public.sys_notice USING btree (tenant_id);


--
-- Name: idx_sys_oper_log_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_oper_log_deleted_at ON public.sys_oper_log USING btree (deleted_at);


--
-- Name: idx_sys_oper_log_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_oper_log_tenant ON public.sys_oper_log USING btree (tenant_id);


--
-- Name: idx_sys_oper_log_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_oper_log_tenant_id ON public.sys_oper_log USING btree (tenant_id);


--
-- Name: idx_sys_post_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_post_deleted_at ON public.sys_post USING btree (deleted_at);


--
-- Name: idx_sys_post_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_post_tenant ON public.sys_post USING btree (tenant_id);


--
-- Name: idx_sys_post_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_post_tenant_id ON public.sys_post USING btree (tenant_id);


--
-- Name: idx_sys_print_template_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_print_template_deleted_at ON public.sys_print_template USING btree (deleted_at);


--
-- Name: idx_sys_print_template_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_print_template_tenant_id ON public.sys_print_template USING btree (tenant_id);


--
-- Name: idx_sys_role_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_role_deleted_at ON public.sys_role USING btree (deleted_at);


--
-- Name: idx_sys_role_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_role_tenant ON public.sys_role USING btree (tenant_id);


--
-- Name: idx_sys_role_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_role_tenant_id ON public.sys_role USING btree (tenant_id);


--
-- Name: idx_sys_tenant_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_tenant_deleted_at ON public.sys_tenant USING btree (deleted_at);


--
-- Name: idx_sys_tenant_key; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_tenant_key ON public.sys_tenant USING btree (tenant_key);


--
-- Name: idx_sys_tenant_tenant_key; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_sys_tenant_tenant_key ON public.sys_tenant USING btree (tenant_key);


--
-- Name: idx_sys_user_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_user_deleted_at ON public.sys_user USING btree (deleted_at);


--
-- Name: idx_sys_user_dept; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_user_dept ON public.sys_user USING btree (dept_id);


--
-- Name: idx_sys_user_dept_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_user_dept_id ON public.sys_user USING btree (dept_id);


--
-- Name: idx_sys_user_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_user_tenant ON public.sys_user USING btree (tenant_id);


--
-- Name: idx_sys_user_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_user_tenant_id ON public.sys_user USING btree (tenant_id);


--
-- Name: idx_tenant_andon; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_andon ON public.andon_calls USING btree (call_no);


--
-- Name: idx_tenant_bom; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_bom ON public.boms USING btree (bom_code);


--
-- Name: idx_tenant_bom_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_bom_code ON public.mdm_bom USING btree (bom_code);


--
-- Name: idx_tenant_check; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_check ON public.stock_checks USING btree (check_no);


--
-- Name: idx_tenant_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_code ON public.sys_interface_config USING btree (code);


--
-- Name: idx_tenant_day_plan; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_day_plan ON public.mes_order_day USING btree (day_plan_no);


--
-- Name: idx_tenant_defect; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_defect ON public.defect_codes USING btree (defect_code);


--
-- Name: idx_tenant_defect_rec; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_defect_rec ON public.defect_records USING btree (record_no);


--
-- Name: idx_tenant_delivery; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_delivery ON public.delivery_orders USING btree (delivery_no);


--
-- Name: idx_tenant_form_field; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_form_field ON public.bpm_form_field USING btree (field_code);


--
-- Name: idx_tenant_fqc; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_fqc ON public.fqcs USING btree (fqc_no);


--
-- Name: idx_tenant_gauge; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_gauge ON public.equ_gauge USING btree (gauge_code);


--
-- Name: idx_tenant_inv_rec; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_inv_rec ON public.inventory_records USING btree (record_no);


--
-- Name: idx_tenant_ipqc; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_ipqc ON public.ip_qcs USING btree (ip_qc_no);


--
-- Name: idx_tenant_iqc; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_iqc ON public.iqcs USING btree (iqc_no);


--
-- Name: idx_tenant_kanban; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_kanban ON public.wms_kanban_pull USING btree (kanban_no);


--
-- Name: idx_tenant_line; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_line ON public.production_lines USING btree (line_code);


--
-- Name: idx_tenant_location; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_location ON public.locations USING btree (location_code);


--
-- Name: idx_tenant_material; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_material ON public.materials USING btree (material_code);


--
-- Name: idx_tenant_mdm_shift_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_mdm_shift_code ON public.mdm_shift USING btree (shift_code);


--
-- Name: idx_tenant_model_flow; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_model_flow ON public.bpm_sequence_flow USING btree (flow_id);


--
-- Name: idx_tenant_model_node; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_model_node ON public.bpm_node_definition USING btree (node_id);


--
-- Name: idx_tenant_model_version; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_model_version ON public.bpm_process_model USING btree (model_code);


--
-- Name: idx_tenant_mold; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_mold ON public.equ_mold USING btree (mold_code);


--
-- Name: idx_tenant_month_plan; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_month_plan ON public.mes_order_month USING btree (month_plan_no);


--
-- Name: idx_tenant_mps; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_mps ON public.mps USING btree (mps_no);


--
-- Name: idx_tenant_mrp; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_mrp ON public.mrps USING btree (mrp_no);


--
-- Name: idx_tenant_ncr; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_ncr ON public.ncrs USING btree (ncr_no);


--
-- Name: idx_tenant_op_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_op_code ON public.mdm_operation USING btree (operation_code);


--
-- Name: idx_tenant_oqc; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_oqc ON public.oqcs USING btree (oqc_no);


--
-- Name: idx_tenant_order; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_order ON public.sales_orders USING btree (order_no);


--
-- Name: idx_tenant_plan_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_plan_code ON public.qc_sampling_plan USING btree (code);


--
-- Name: idx_tenant_point; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_point ON public.dc_data_point USING btree (point_code);


--
-- Name: idx_tenant_porder; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_porder ON public.production_orders USING btree (order_no);


--
-- Name: idx_tenant_process; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_process ON public.processes USING btree (process_code);


--
-- Name: idx_tenant_process_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_process_code ON public.mes_process USING btree (process_code);


--
-- Name: idx_tenant_recv; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_recv ON public.receive_orders USING btree (receive_no);


--
-- Name: idx_tenant_resource; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_resource ON public.resources USING btree (resource_code);


--
-- Name: idx_tenant_route; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_route ON public.routes USING btree (route_code);


--
-- Name: idx_tenant_schedule; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_schedule ON public.schedule_plans USING btree (plan_no);


--
-- Name: idx_tenant_scheduling_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_scheduling_code ON public.mes_work_scheduling USING btree (tenant_id, scheduling_code);


--
-- Name: idx_tenant_serial; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_serial ON public.serial_numbers USING btree (serial_number);


--
-- Name: idx_tenant_session; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_session ON public.ai_chat_conversation USING btree (session_id);


--
-- Name: idx_tenant_shift; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_shift ON public.shifts USING btree (shift_code);


--
-- Name: idx_tenant_sideloc; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_sideloc ON public.wms_side_location USING btree (location_code);


--
-- Name: idx_tenant_station; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_station ON public.workstations USING btree (station_code);


--
-- Name: idx_tenant_supplier_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_supplier_code ON public.mdm_supplier USING btree (code);


--
-- Name: idx_tenant_team_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_team_code ON public.mes_team USING btree (team_code);


--
-- Name: idx_tenant_trace; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_trace ON public.trace_records USING btree (trace_no);


--
-- Name: idx_tenant_transfer; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_transfer ON public.transfer_orders USING btree (transfer_no);


--
-- Name: idx_tenant_warehouse; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_warehouse ON public.warehouses USING btree (warehouse_code);


--
-- Name: idx_tenant_wc; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_wc ON public.work_centers USING btree (work_center_code);


--
-- Name: idx_tenant_workshop; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_tenant_workshop ON public.workshops USING btree (workshop_code);


--
-- Name: idx_tra_andon_call_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_tra_andon_call_deleted_at ON public.tra_andon_call USING btree (deleted_at);


--
-- Name: idx_tra_andon_call_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_tra_andon_call_tenant ON public.tra_andon_call USING btree (tenant_id);


--
-- Name: idx_tra_andon_call_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_tra_andon_call_tenant_id ON public.tra_andon_call USING btree (tenant_id);


--
-- Name: idx_tra_data_collection_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_tra_data_collection_deleted_at ON public.tra_data_collection USING btree (deleted_at);


--
-- Name: idx_tra_data_collection_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_tra_data_collection_tenant ON public.tra_data_collection USING btree (tenant_id);


--
-- Name: idx_tra_data_collection_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_tra_data_collection_tenant_id ON public.tra_data_collection USING btree (tenant_id);


--
-- Name: idx_tra_serial_number_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_tra_serial_number_deleted_at ON public.tra_serial_number USING btree (deleted_at);


--
-- Name: idx_tra_serial_number_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_tra_serial_number_tenant ON public.tra_serial_number USING btree (tenant_id);


--
-- Name: idx_tra_serial_number_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_tra_serial_number_tenant_id ON public.tra_serial_number USING btree (tenant_id);


--
-- Name: idx_tra_trace_record_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_tra_trace_record_deleted_at ON public.tra_trace_record USING btree (deleted_at);


--
-- Name: idx_tra_trace_record_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_tra_trace_record_tenant ON public.tra_trace_record USING btree (tenant_id);


--
-- Name: idx_tra_trace_record_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_tra_trace_record_tenant_id ON public.tra_trace_record USING btree (tenant_id);


--
-- Name: idx_trace_records_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_trace_records_deleted_at ON public.trace_records USING btree (deleted_at);


--
-- Name: idx_trace_records_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_trace_records_tenant_id ON public.trace_records USING btree (tenant_id);


--
-- Name: idx_transfer_orders_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_transfer_orders_deleted_at ON public.transfer_orders USING btree (deleted_at);


--
-- Name: idx_transfer_orders_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_transfer_orders_tenant_id ON public.transfer_orders USING btree (tenant_id);


--
-- Name: idx_visual_inspection_result_task_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_visual_inspection_result_task_id ON public.visual_inspection_result USING btree (task_id);


--
-- Name: idx_visual_inspection_result_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_visual_inspection_result_tenant_id ON public.visual_inspection_result USING btree (tenant_id);


--
-- Name: idx_visual_inspection_task_task_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_visual_inspection_task_task_no ON public.visual_inspection_task USING btree (task_no);


--
-- Name: idx_visual_inspection_task_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_visual_inspection_task_tenant_id ON public.visual_inspection_task USING btree (tenant_id);


--
-- Name: idx_warehouses_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_warehouses_deleted_at ON public.warehouses USING btree (deleted_at);


--
-- Name: idx_warehouses_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_warehouses_tenant_id ON public.warehouses USING btree (tenant_id);


--
-- Name: idx_wms_agv_device_agv_code; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_wms_agv_device_agv_code ON public.wms_agv_device USING btree (agv_code);


--
-- Name: idx_wms_agv_task_task_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_wms_agv_task_task_no ON public.wms_agv_task USING btree (task_no);


--
-- Name: idx_wms_delivery_order_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_delivery_order_deleted_at ON public.wms_delivery_order USING btree (deleted_at);


--
-- Name: idx_wms_delivery_order_item_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_delivery_order_item_deleted_at ON public.wms_delivery_order_item USING btree (deleted_at);


--
-- Name: idx_wms_delivery_order_item_delivery; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_delivery_order_item_delivery ON public.wms_delivery_order_item USING btree (delivery_id);


--
-- Name: idx_wms_delivery_order_item_delivery_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_delivery_order_item_delivery_id ON public.wms_delivery_order_item USING btree (delivery_id);


--
-- Name: idx_wms_delivery_order_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_delivery_order_tenant ON public.wms_delivery_order USING btree (tenant_id);


--
-- Name: idx_wms_delivery_order_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_delivery_order_tenant_id ON public.wms_delivery_order USING btree (tenant_id);


--
-- Name: idx_wms_inventory_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_inventory_deleted_at ON public.wms_inventory USING btree (deleted_at);


--
-- Name: idx_wms_inventory_location; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_inventory_location ON public.wms_inventory USING btree (location_id);


--
-- Name: idx_wms_inventory_location_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_inventory_location_id ON public.wms_inventory USING btree (location_id);


--
-- Name: idx_wms_inventory_material; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_inventory_material ON public.wms_inventory USING btree (material_id);


--
-- Name: idx_wms_inventory_material_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_inventory_material_id ON public.wms_inventory USING btree (material_id);


--
-- Name: idx_wms_inventory_record_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_inventory_record_deleted_at ON public.wms_inventory_record USING btree (deleted_at);


--
-- Name: idx_wms_inventory_record_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_inventory_record_tenant ON public.wms_inventory_record USING btree (tenant_id);


--
-- Name: idx_wms_inventory_record_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_inventory_record_tenant_id ON public.wms_inventory_record USING btree (tenant_id);


--
-- Name: idx_wms_inventory_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_inventory_tenant ON public.wms_inventory USING btree (tenant_id);


--
-- Name: idx_wms_inventory_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_inventory_tenant_id ON public.wms_inventory USING btree (tenant_id);


--
-- Name: idx_wms_inventory_warehouse; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_inventory_warehouse ON public.wms_inventory USING btree (warehouse_id);


--
-- Name: idx_wms_inventory_warehouse_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_inventory_warehouse_id ON public.wms_inventory USING btree (warehouse_id);


--
-- Name: idx_wms_kanban_pull_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_kanban_pull_deleted_at ON public.wms_kanban_pull USING btree (deleted_at);


--
-- Name: idx_wms_kanban_pull_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_kanban_pull_tenant_id ON public.wms_kanban_pull USING btree (tenant_id);


--
-- Name: idx_wms_location_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_location_deleted_at ON public.wms_location USING btree (deleted_at);


--
-- Name: idx_wms_location_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_location_tenant ON public.wms_location USING btree (tenant_id);


--
-- Name: idx_wms_location_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_location_tenant_id ON public.wms_location USING btree (tenant_id);


--
-- Name: idx_wms_location_warehouse; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_location_warehouse ON public.wms_location USING btree (warehouse_id);


--
-- Name: idx_wms_production_complete_complete_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_wms_production_complete_complete_no ON public.wms_production_complete USING btree (complete_no);


--
-- Name: idx_wms_production_complete_item_complete_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_complete_item_complete_id ON public.wms_production_complete_item USING btree (complete_id);


--
-- Name: idx_wms_production_complete_item_material_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_complete_item_material_id ON public.wms_production_complete_item USING btree (material_id);


--
-- Name: idx_wms_production_complete_item_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_complete_item_tenant_id ON public.wms_production_complete_item USING btree (tenant_id);


--
-- Name: idx_wms_production_complete_production_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_complete_production_order_id ON public.wms_production_complete USING btree (production_order_id);


--
-- Name: idx_wms_production_complete_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_complete_tenant_id ON public.wms_production_complete USING btree (tenant_id);


--
-- Name: idx_wms_production_complete_workshop_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_complete_workshop_id ON public.wms_production_complete USING btree (workshop_id);


--
-- Name: idx_wms_production_complete_workstation_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_complete_workstation_id ON public.wms_production_complete USING btree (workstation_id);


--
-- Name: idx_wms_production_issue_issue_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_wms_production_issue_issue_no ON public.wms_production_issue USING btree (issue_no);


--
-- Name: idx_wms_production_issue_item_issue_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_issue_item_issue_id ON public.wms_production_issue_item USING btree (issue_id);


--
-- Name: idx_wms_production_issue_item_material_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_issue_item_material_id ON public.wms_production_issue_item USING btree (material_id);


--
-- Name: idx_wms_production_issue_item_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_issue_item_tenant_id ON public.wms_production_issue_item USING btree (tenant_id);


--
-- Name: idx_wms_production_issue_production_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_issue_production_order_id ON public.wms_production_issue USING btree (production_order_id);


--
-- Name: idx_wms_production_issue_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_issue_tenant_id ON public.wms_production_issue USING btree (tenant_id);


--
-- Name: idx_wms_production_issue_workshop_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_issue_workshop_id ON public.wms_production_issue USING btree (workshop_id);


--
-- Name: idx_wms_production_issue_workstation_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_issue_workstation_id ON public.wms_production_issue USING btree (workstation_id);


--
-- Name: idx_wms_production_return_item_material_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_return_item_material_id ON public.wms_production_return_item USING btree (material_id);


--
-- Name: idx_wms_production_return_item_return_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_return_item_return_id ON public.wms_production_return_item USING btree (return_id);


--
-- Name: idx_wms_production_return_item_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_return_item_tenant_id ON public.wms_production_return_item USING btree (tenant_id);


--
-- Name: idx_wms_production_return_production_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_return_production_order_id ON public.wms_production_return USING btree (production_order_id);


--
-- Name: idx_wms_production_return_return_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_wms_production_return_return_no ON public.wms_production_return USING btree (return_no);


--
-- Name: idx_wms_production_return_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_return_tenant_id ON public.wms_production_return USING btree (tenant_id);


--
-- Name: idx_wms_production_return_workshop_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_return_workshop_id ON public.wms_production_return USING btree (workshop_id);


--
-- Name: idx_wms_production_return_workstation_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_return_workstation_id ON public.wms_production_return USING btree (workstation_id);


--
-- Name: idx_wms_production_stock_in_complete_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_stock_in_complete_id ON public.wms_production_stock_in USING btree (complete_id);


--
-- Name: idx_wms_production_stock_in_item_material_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_stock_in_item_material_id ON public.wms_production_stock_in_item USING btree (material_id);


--
-- Name: idx_wms_production_stock_in_item_stock_in_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_stock_in_item_stock_in_id ON public.wms_production_stock_in_item USING btree (stock_in_id);


--
-- Name: idx_wms_production_stock_in_item_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_stock_in_item_tenant_id ON public.wms_production_stock_in_item USING btree (tenant_id);


--
-- Name: idx_wms_production_stock_in_location_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_stock_in_location_id ON public.wms_production_stock_in USING btree (location_id);


--
-- Name: idx_wms_production_stock_in_stock_in_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_wms_production_stock_in_stock_in_no ON public.wms_production_stock_in USING btree (stock_in_no);


--
-- Name: idx_wms_production_stock_in_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_stock_in_tenant_id ON public.wms_production_stock_in USING btree (tenant_id);


--
-- Name: idx_wms_production_stock_in_warehouse_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_production_stock_in_warehouse_id ON public.wms_production_stock_in USING btree (warehouse_id);


--
-- Name: idx_wms_purchase_return_item_material_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_purchase_return_item_material_id ON public.wms_purchase_return_item USING btree (material_id);


--
-- Name: idx_wms_purchase_return_item_return_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_purchase_return_item_return_id ON public.wms_purchase_return_item USING btree (return_id);


--
-- Name: idx_wms_purchase_return_item_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_purchase_return_item_tenant_id ON public.wms_purchase_return_item USING btree (tenant_id);


--
-- Name: idx_wms_purchase_return_purchase_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_purchase_return_purchase_order_id ON public.wms_purchase_return USING btree (purchase_order_id);


--
-- Name: idx_wms_purchase_return_return_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_wms_purchase_return_return_no ON public.wms_purchase_return USING btree (return_no);


--
-- Name: idx_wms_purchase_return_supplier_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_purchase_return_supplier_id ON public.wms_purchase_return USING btree (supplier_id);


--
-- Name: idx_wms_purchase_return_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_purchase_return_tenant_id ON public.wms_purchase_return USING btree (tenant_id);


--
-- Name: idx_wms_purchase_return_warehouse_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_purchase_return_warehouse_id ON public.wms_purchase_return USING btree (warehouse_id);


--
-- Name: idx_wms_receive_order_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_receive_order_deleted_at ON public.wms_receive_order USING btree (deleted_at);


--
-- Name: idx_wms_receive_order_item_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_receive_order_item_deleted_at ON public.wms_receive_order_item USING btree (deleted_at);


--
-- Name: idx_wms_receive_order_item_receive; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_receive_order_item_receive ON public.wms_receive_order_item USING btree (receive_id);


--
-- Name: idx_wms_receive_order_item_receive_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_receive_order_item_receive_id ON public.wms_receive_order_item USING btree (receive_id);


--
-- Name: idx_wms_receive_order_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_receive_order_tenant ON public.wms_receive_order USING btree (tenant_id);


--
-- Name: idx_wms_receive_order_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_receive_order_tenant_id ON public.wms_receive_order USING btree (tenant_id);


--
-- Name: idx_wms_sales_return_customer_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_sales_return_customer_id ON public.wms_sales_return USING btree (customer_id);


--
-- Name: idx_wms_sales_return_item_material_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_sales_return_item_material_id ON public.wms_sales_return_item USING btree (material_id);


--
-- Name: idx_wms_sales_return_item_return_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_sales_return_item_return_id ON public.wms_sales_return_item USING btree (return_id);


--
-- Name: idx_wms_sales_return_item_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_sales_return_item_tenant_id ON public.wms_sales_return_item USING btree (tenant_id);


--
-- Name: idx_wms_sales_return_return_no; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_wms_sales_return_return_no ON public.wms_sales_return USING btree (return_no);


--
-- Name: idx_wms_sales_return_sales_order_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_sales_return_sales_order_id ON public.wms_sales_return USING btree (sales_order_id);


--
-- Name: idx_wms_sales_return_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_sales_return_tenant_id ON public.wms_sales_return USING btree (tenant_id);


--
-- Name: idx_wms_sales_return_warehouse_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_sales_return_warehouse_id ON public.wms_sales_return USING btree (warehouse_id);


--
-- Name: idx_wms_side_location_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_side_location_deleted_at ON public.wms_side_location USING btree (deleted_at);


--
-- Name: idx_wms_side_location_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_side_location_tenant_id ON public.wms_side_location USING btree (tenant_id);


--
-- Name: idx_wms_stock_check_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_stock_check_deleted_at ON public.wms_stock_check USING btree (deleted_at);


--
-- Name: idx_wms_stock_check_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_stock_check_tenant ON public.wms_stock_check USING btree (tenant_id);


--
-- Name: idx_wms_stock_check_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_stock_check_tenant_id ON public.wms_stock_check USING btree (tenant_id);


--
-- Name: idx_wms_transfer_order_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_transfer_order_deleted_at ON public.wms_transfer_order USING btree (deleted_at);


--
-- Name: idx_wms_transfer_order_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_transfer_order_tenant ON public.wms_transfer_order USING btree (tenant_id);


--
-- Name: idx_wms_transfer_order_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_transfer_order_tenant_id ON public.wms_transfer_order USING btree (tenant_id);


--
-- Name: idx_wms_warehouse_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_warehouse_deleted_at ON public.wms_warehouse USING btree (deleted_at);


--
-- Name: idx_wms_warehouse_tenant; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_warehouse_tenant ON public.wms_warehouse USING btree (tenant_id);


--
-- Name: idx_wms_warehouse_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_wms_warehouse_tenant_id ON public.wms_warehouse USING btree (tenant_id);


--
-- Name: idx_work_centers_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_work_centers_deleted_at ON public.work_centers USING btree (deleted_at);


--
-- Name: idx_work_centers_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_work_centers_tenant_id ON public.work_centers USING btree (tenant_id);


--
-- Name: idx_workshops_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_workshops_deleted_at ON public.workshops USING btree (deleted_at);


--
-- Name: idx_workshops_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_workshops_tenant_id ON public.workshops USING btree (tenant_id);


--
-- Name: idx_workstations_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_workstations_deleted_at ON public.workstations USING btree (deleted_at);


--
-- Name: idx_workstations_tenant_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_workstations_tenant_id ON public.workstations USING btree (tenant_id);


--
-- Name: bpm_form_field fk_bpm_form_definition_fields; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_form_field
    ADD CONSTRAINT fk_bpm_form_definition_fields FOREIGN KEY (form_id) REFERENCES public.bpm_form_definition(id);


--
-- Name: bpm_task_instance fk_bpm_process_instance_tasks; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_task_instance
    ADD CONSTRAINT fk_bpm_process_instance_tasks FOREIGN KEY (instance_id) REFERENCES public.bpm_process_instance(id);


--
-- Name: bpm_sequence_flow fk_bpm_process_model_flows; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_sequence_flow
    ADD CONSTRAINT fk_bpm_process_model_flows FOREIGN KEY (model_id) REFERENCES public.bpm_process_model(id);


--
-- Name: bpm_node_definition fk_bpm_process_model_nodes; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.bpm_node_definition
    ADD CONSTRAINT fk_bpm_process_model_nodes FOREIGN KEY (model_id) REFERENCES public.bpm_process_model(id);


--
-- Name: mes_order_day_item fk_mes_order_day_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_order_day_item
    ADD CONSTRAINT fk_mes_order_day_items FOREIGN KEY (day_plan_id) REFERENCES public.mes_order_day(id);


--
-- Name: mes_order_month_item fk_mes_order_month_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_order_month_item
    ADD CONSTRAINT fk_mes_order_month_items FOREIGN KEY (month_plan_id) REFERENCES public.mes_order_month(id);


--
-- Name: mes_process_operation fk_mes_process_operations; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_process_operation
    ADD CONSTRAINT fk_mes_process_operations FOREIGN KEY (process_id) REFERENCES public.mes_process(id);


--
-- Name: mes_work_scheduling_detail fk_mes_work_scheduling_details; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mes_work_scheduling_detail
    ADD CONSTRAINT fk_mes_work_scheduling_details FOREIGN KEY (scheduling_id) REFERENCES public.mes_work_scheduling(id);


--
-- Name: scp_inquiry_item fk_scp_customer_inquiry_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_inquiry_item
    ADD CONSTRAINT fk_scp_customer_inquiry_items FOREIGN KEY (inquiry_id) REFERENCES public.scp_customer_inquiry(id);


--
-- Name: scp_mrs_item fk_scp_mrs_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_mrs_item
    ADD CONSTRAINT fk_scp_mrs_items FOREIGN KEY (mrs_id) REFERENCES public.scp_mrs(id);


--
-- Name: scp_purchase_order_item fk_scp_purchase_order_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_purchase_order_item
    ADD CONSTRAINT fk_scp_purchase_order_items FOREIGN KEY (po_id) REFERENCES public.scp_purchase_order(id);


--
-- Name: scp_purchase_plan_item fk_scp_purchase_plan_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_purchase_plan_item
    ADD CONSTRAINT fk_scp_purchase_plan_items FOREIGN KEY (plan_id) REFERENCES public.scp_purchase_plan(id);


--
-- Name: scp_rfq_invite fk_scp_rfq_invites; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_rfq_invite
    ADD CONSTRAINT fk_scp_rfq_invites FOREIGN KEY (rfq_id) REFERENCES public.scp_rfq(id);


--
-- Name: scp_rfq_item fk_scp_rfq_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_rfq_item
    ADD CONSTRAINT fk_scp_rfq_items FOREIGN KEY (rfq_id) REFERENCES public.scp_rfq(id);


--
-- Name: scp_sales_order_item fk_scp_sales_order_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_sales_order_item
    ADD CONSTRAINT fk_scp_sales_order_items FOREIGN KEY (so_id) REFERENCES public.scp_sales_order(id);


--
-- Name: scp_quote_item fk_scp_supplier_quote_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scp_quote_item
    ADD CONSTRAINT fk_scp_supplier_quote_items FOREIGN KEY (quote_id) REFERENCES public.scp_supplier_quote(id);


--
-- Name: supplier_asn_items fk_supplier_asns_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.supplier_asn_items
    ADD CONSTRAINT fk_supplier_asns_items FOREIGN KEY (asn_id) REFERENCES public.supplier_asns(id);


--
-- Name: wms_production_complete_item fk_wms_production_complete_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_complete_item
    ADD CONSTRAINT fk_wms_production_complete_items FOREIGN KEY (complete_id) REFERENCES public.wms_production_complete(id);


--
-- Name: wms_production_issue_item fk_wms_production_issue_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_issue_item
    ADD CONSTRAINT fk_wms_production_issue_items FOREIGN KEY (issue_id) REFERENCES public.wms_production_issue(id);


--
-- Name: wms_production_return_item fk_wms_production_return_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_return_item
    ADD CONSTRAINT fk_wms_production_return_items FOREIGN KEY (return_id) REFERENCES public.wms_production_return(id);


--
-- Name: wms_production_stock_in_item fk_wms_production_stock_in_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_production_stock_in_item
    ADD CONSTRAINT fk_wms_production_stock_in_items FOREIGN KEY (stock_in_id) REFERENCES public.wms_production_stock_in(id);


--
-- Name: wms_purchase_return_item fk_wms_purchase_return_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_purchase_return_item
    ADD CONSTRAINT fk_wms_purchase_return_items FOREIGN KEY (return_id) REFERENCES public.wms_purchase_return(id);


--
-- Name: wms_sales_return_item fk_wms_sales_return_items; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.wms_sales_return_item
    ADD CONSTRAINT fk_wms_sales_return_items FOREIGN KEY (return_id) REFERENCES public.wms_sales_return(id);


--
-- PostgreSQL database dump complete
--

\unrestrict NjEIdbnFSsNlsfQSLqiIFhLECW00clQSU4mt2jafdZCf6rRWgmLIH6QHDhBj390

